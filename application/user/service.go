package user

import (
	"errors"

	domainModel "github.com/karamaru-alpha/ddd-sample/domain/model/user"
	domainService "github.com/karamaru-alpha/ddd-sample/domain/service/user"
)

// IApplicationService Interface of ApplicationService realize usecase
type IApplicationService interface {
	Register(CreateCommand) (string, error)
}

type applicationService struct {
	domainService domainService.IDomainService
	repository    domainModel.IRepository
	factory       domainModel.IFactory
}

// NewApplicationService Function create user ApplicationService
func NewApplicationService(
	userDomainService domainService.IDomainService,
	userRepository domainModel.IRepository,
	userFactory domainModel.IFactory,
) IApplicationService {
	return &applicationService{
		domainService: userDomainService,
		repository:    userRepository,
		factory:       userFactory,
	}
}

// Register ApplicationService realize account registration
func (as applicationService) Register(createCommand CreateCommand) (string, error) {

	userName, err := domainModel.NewName(createCommand.Name)
	if err != nil {
		return "", err
	}

	userMailAddress, err := domainModel.NewMailAddress(createCommand.MailAddress)
	if err != nil {
		return "", err
	}

	userPlainPassword, err := domainModel.NewPlainPassword(createCommand.PlainPassword)
	if err != nil {
		return "", err
	}

	user, err := as.factory.Create(userName, userMailAddress, userPlainPassword)
	if err != nil {
		return "", err
	}

	isDup, err := as.domainService.Exists(user)
	if err != nil {
		return "", err
	}
	if isDup {
		return "", errors.New("This user already exists")
	}

	// TODO Implement transaction
	if err := as.repository.Save(user); err != nil {
		return "", err
	}

	return GenerateJWT(&user.ID)
}
