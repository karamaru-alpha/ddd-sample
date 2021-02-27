package user

import (
	"errors"
)

// User Entity express user
type User struct {
	ID          ID
	Name        Name
	MailAddress MailAddress
}

// NewUser Constructor generate user entity
func NewUser(id *ID, name *Name, mailAddress *MailAddress) (*User, error) {

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
