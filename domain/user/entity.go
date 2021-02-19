package user

import (
	"errors"

	"github.com/google/uuid"
)

// User Entity express user.
type User struct {
	id         ID
	name       Name
	mailAdress MailAdress
}

// NewUser Constructor generate user entity.
func NewUser(id ID, name Name, mailAdress MailAdress) (*User, error) {

	if id == ID(uuid.Nil) {
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
