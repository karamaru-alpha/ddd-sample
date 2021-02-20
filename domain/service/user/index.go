package user

import (
	entity "github.com/karamaru-alpha/ddd-sample/domain/entity/user"
	repository "github.com/karamaru-alpha/ddd-sample/domain/repository/user"
)

// IDomainService Inteface of DomainService express domain activity
type IDomainService interface {
	Exists(user *entity.User) (bool, error)
}

type domainService struct {
	repository repository.IRepository
}

// NewDomainService Function create DomainService
func NewDomainService(userRepository repository.IRepository) IDomainService {
	return &domainService{
		repository: userRepository,
	}
}

// Exists DomainService check user duplicate
func (ds domainService) Exists(user *entity.User) (bool, error) {

	dupUser, err := ds.repository.Find(&user.MailAdress)
	if err != nil {
		return false, err
	}

	return dupUser != nil, nil
}
