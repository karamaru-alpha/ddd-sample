package user

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// PlainPassword Value object express user's PlainPassword
type PlainPassword string

// NOTE: define global because regexp's compile is slow.
var passwordRegexp = regexp.MustCompile(`^[a-zA-Z_\d]{8,30}$`)

// NewPlainPassword Constructor generate user's PlainPassword value object
func NewPlainPassword(arg string) (*PlainPassword, error) {

	if arg == "" {
		return nil, errors.New("UserPlainPassword is null")
	}

	if !passwordRegexp.MatchString(arg) {
		return nil, errors.New("UserPassword should consist of alphabet or half-size number or underscore, using more than 8 characters but less than 30")
	}

	plainPassword := PlainPassword(arg)
	return &plainPassword, nil
}

// ToHash Convert plainPassword string to []byte
func (plainPassword PlainPassword) ToHash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
}
