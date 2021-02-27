package user

import (
	domainModel "github.com/karamaru-alpha/ddd-sample/domain/model/user"
)

// IDomainService Inteface of DomainService express domain activity
type IDomainService interface {
	Exists(user *domainModel.User) (bool, error)
}

type domainService struct {
	repository domainModel.IRepository
}

// NewDomainService Function create DomainService
func NewDomainService(userRepository domainModel.IRepository) IDomainService {
	return &domainService{
		repository: userRepository,
	}
}

// Exists DomainService check user duplicate
func (ds domainService) Exists(user *domainModel.User) (bool, error) {

	dupUser, err := ds.repository.Find(&user.MailAddress)
	if err != nil {
		return false, err
	}

	return dupUser != nil, nil
}
