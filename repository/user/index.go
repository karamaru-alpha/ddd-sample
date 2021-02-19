package user

import (
	entity "github.com/karamaru-alpha/ddd-sample/domain/user/entity"
	valueObject "github.com/karamaru-alpha/ddd-sample/domain/user/valueobject"
)

// IRepository Repository perpetuate/rebuild user entity.
type IRepository interface {
	Find(valueObject.MailAdress) (*entity.User, error)
}
