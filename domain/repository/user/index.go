package user

import (
	entity "github.com/karamaru-alpha/ddd-sample/domain/entity/user"
	valueObject "github.com/karamaru-alpha/ddd-sample/domain/value_object/user"
)

// IRepository Repository perpetuate/rebuild user entity.
type IRepository interface {
	Find(valueObject.MailAdress) (*entity.User, error)
}
