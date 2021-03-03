package user

import (
	"errors"
)

// User Entity express user
type User struct {
	ID             ID
	Name           Name
	MailAddress    MailAddress
	HashedPassword HashedPassword
}

// NewUser Constructor generate user entity
func NewUser(id *ID, name *Name, mailAddress *MailAddress, hashedPassword *HashedPassword) (*User, error) {

	if id == nil {
		return nil, errors.New("UserID is null")
	}
	if name == nil {
		return nil, errors.New("UserName is null")
	}
	if mailAddress == nil {
		return nil, errors.New("UserMailAddress is null")
	}
	if hashedPassword == nil {
		return nil, errors.New("UserHashedPasseord is null")
	}

	user := new(User)
	user.ID = *id
	user.Name = *name
	user.MailAddress = *mailAddress
	user.HashedPassword = *hashedPassword

	return user, nil
}
