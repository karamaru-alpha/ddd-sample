package user

import (
	entity "github.com/karamaru-alpha/ddd-sample/domain/entity/user"
	ur "github.com/karamaru-alpha/ddd-sample/domain/repository/user"
)

// DomainService DomainService interface express domain activity.
type DomainService interface {
	Exists(user *entity.User) (bool, error)
}

type domainService struct {
	repository ur.IRepository
}

// NewDomainService create DomainService
func NewDomainService(userRepository ur.IRepository) DomainService {
	return &domainService{
		repository: userRepository,
	}
}

// Exists DomainService check user duplicate.
func (ds domainService) Exists(user *entity.User) (bool, error) {

	dupUser, err := ds.repository.Find(user.MailAdress)
	if err != nil {
		return false, err
	}

	return dupUser != nil, nil
}
