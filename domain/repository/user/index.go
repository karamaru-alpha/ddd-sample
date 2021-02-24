package user

import (
	entity "github.com/karamaru-alpha/ddd-sample/domain/entity/user"
	valueObject "github.com/karamaru-alpha/ddd-sample/domain/value_object/user"
)

// IRepository Interface of Repository perpetuate/rebuild user entity
type IRepository interface {
	Save(*entity.User) error
	Find(*valueObject.MailAddress) (*entity.User, error)
}
