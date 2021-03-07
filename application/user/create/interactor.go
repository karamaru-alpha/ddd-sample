package user

import (
	"errors"

	domainModel "github.com/karamaru-alpha/ddd-sample/domain/model/user"
	domainService "github.com/karamaru-alpha/ddd-sample/domain/service/user"
)

type interactor struct {
	domainService domainService.IDomainService
	repository    domainModel.IRepository
	factory       domainModel.IFactory
}

// NewInteractor Create user application service about creation
func NewInteractor(
	userDomainService domainService.IDomainService,
	userRepository domainModel.IRepository,
	userFactory domainModel.IFactory,
) IInputPort {
	return &interactor{
		domainService: userDomainService,
		repository:    userRepository,
		factory:       userFactory,
	}
}

// Handle ApplicationService realize account registration
func (as interactor) Handle(inputData InputData) OutputData {

	userName, err := domainModel.NewName(inputData.Name)
	if err != nil {
		return OutputData{Token: "", Err: err}
	}

	userMailAddress, err := domainModel.NewMailAddress(inputData.MailAddress)
	if err != nil {
		return OutputData{Token: "", Err: err}
	}

	userPlainPassword, err := domainModel.NewPlainPassword(inputData.PlainPassword)
	if err != nil {
		return OutputData{Token: "", Err: err}
	}

	user, err := as.factory.Create(userName, userMailAddress, userPlainPassword)
	if err != nil {
		return OutputData{Token: "", Err: err}
	}

	isDup, err := as.domainService.Exists(user)
	if err != nil {
		return OutputData{Token: "", Err: err}
	}
	if isDup {
		return OutputData{Token: "", Err: errors.New("This user already exists")}
	}

	// TODO Implement transaction
	if err := as.repository.Save(user); err != nil {
		return OutputData{Token: "", Err: err}
	}

	token, err := JWTGenerator(&user.ID)
	return OutputData{
		Token: token,
		Err:   err,
	}
}
