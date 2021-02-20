package user

import (
	"errors"

	repository "github.com/karamaru-alpha/ddd-sample/domain/repository/user"
	domainService "github.com/karamaru-alpha/ddd-sample/domain/service/user"
	valueObject "github.com/karamaru-alpha/ddd-sample/domain/value_object/user"
	factory "github.com/karamaru-alpha/ddd-sample/factory/user"
)

// IApplicationService Interface of ApplicationService realize usecase
type IApplicationService interface {
	Register(name string, mailAdress string) error
}

type applicationService struct {
	domainService domainService.IDomainService
	repository    repository.IRepository
	factory       factory.IFactory
}

// NewApplicationService Function create user ApplicationService
func NewApplicationService(
	userDomainService domainService.IDomainService,
	userRepository repository.IRepository,
	userFactory factory.IFactory,
) IApplicationService {
	return &applicationService{
		domainService: userDomainService,
		repository:    userRepository,
		factory:       userFactory,
	}
}

// Register ApplicationService realize account registration
func (as applicationService) Register(name string, mailAdress string) error {

	userName, err := valueObject.NewName(name)
	if err != nil {
		return err
	}

	userMailAdress, err := valueObject.NewMailAdress(mailAdress)
	if err != nil {
		return err
	}

	user, err := as.factory.Create(userName, userMailAdress)
	if err != nil {
		return err
	}

	isDup, err := as.domainService.Exists(user)
	if err != nil {
		return err
	}
	if isDup {
		return errors.New("This user already exists")
	}

	return as.repository.Save(user)
}
