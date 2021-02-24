package user

import (
	"errors"

	valueObject "github.com/karamaru-alpha/ddd-sample/domain/value_object/user"
)

// User Entity express user
type User struct {
	ID          valueObject.ID
	Name        valueObject.Name
	MailAddress valueObject.MailAddress
}

// NewUser Constructor generate user entity
func NewUser(id *valueObject.ID, name *valueObject.Name, mailAddress *valueObject.MailAddress) (*User, error) {

	if id == nil {
		return nil, errors.New("UserID is null")
	}
	if name == nil {
		return nil, errors.New("UserName is null")
	}
	if mailAddress == nil {
		return nil, errors.New("UserMailAddress is null")
	}

	user := new(User)
	user.ID = *id
	user.Name = *name
	user.MailAddress = *mailAddress

	return user, nil
}
