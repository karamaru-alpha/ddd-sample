package user

import (
	entity "github.com/karamaru-alpha/ddd-sample/domain/entity/user"
	repository "github.com/karamaru-alpha/ddd-sample/domain/repository/user"
	valueObject "github.com/karamaru-alpha/ddd-sample/domain/value_object/user"
)

type repositoryImpl struct{}

// NewRepositoryImpl Function create infrastructure to implement user repository
func NewRepositoryImpl() repository.IRepository {
	return &repositoryImpl{}
}

// TODO implment persist user object
// Save persist user object
func (uri repositoryImpl) Save(user *entity.User) error {
	return nil
}

// TODO implment rebuild user object by mail address
// Save rebuild user object by mail address
func (uri repositoryImpl) Find(mailAddress *valueObject.MailAddress) (*entity.User, error) {
	return nil, nil
}
