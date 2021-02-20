package user

import (
	"errors"

	valueObject "github.com/karamaru-alpha/ddd-sample/domain/value_object/user"
)

// User Entity express user
type User struct {
	ID         valueObject.ID
	Name       valueObject.Name
	MailAdress valueObject.MailAdress
}

// NewUser Constructor generate user entity
func NewUser(id *valueObject.ID, name *valueObject.Name, mailAdress *valueObject.MailAdress) (*User, error) {

	if id == nil {
		return nil, errors.New("UserID is null")
	}
	if name == nil {
		return nil, errors.New("UserName is null")
	}
	if mailAdress == nil {
		return nil, errors.New("UserMailAdress is null")
	}

	user := new(User)
	user.ID = *id
	user.Name = *name
	user.MailAdress = *mailAdress

	return user, nil
}
