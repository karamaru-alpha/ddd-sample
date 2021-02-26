package user

import (
	"errors"
)

// MailAddress Value object express user's mail address
type MailAddress string

// NewMailAddress Constructor generate user's MailAddress value object
func NewMailAddress(arg string) (*MailAddress, error) {

	if arg == "" {
		return nil, errors.New("UserMailAddress is null")
	}

	mailAddress := MailAddress(arg)
	return &mailAddress, nil
}

// ToString Convert UserMailAddress to String
func (mailAddress MailAddress) ToString() string {
	return string(mailAddress)
}
