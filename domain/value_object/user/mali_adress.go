package user

import (
	"errors"
)

// MailAdress Value object express user's mail adress
type MailAdress string

// NewMailAdress Constructor generate user's MailAdress value object
func NewMailAdress(arg string) (*MailAdress, error) {

	if arg == "" {
		return nil, errors.New("UserMailAdress is null")
	}

	mailAdress := MailAdress(arg)
	return &mailAdress, nil
}
