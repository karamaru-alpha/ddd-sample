package user

import (
	"errors"

	"github.com/google/uuid"
	valueObject "github.com/karamaru-alpha/ddd-sample/domain/user/valueobject"
)

// User Entity express user.
type User struct {
	id         valueObject.ID
	name       valueObject.Name
	mailAdress valueObject.MailAdress
}

// NewUser Constructor generate user entity.
func NewUser(id valueObject.ID, name valueObject.Name, mailAdress valueObject.MailAdress) (*User, error) {

	if id == valueObject.ID(uuid.Nil) {
		return nil, errors.New("UserID is null")
	}
	if name == "" {
		return nil, errors.New("UserName is null")
	}
	if mailAdress == "" {
		return nil, errors.New("UserMailAdress is null")
	}

	user := new(User)
	user.id = id
	user.name = name
	user.mailAdress = mailAdress

	return user, nil
}
