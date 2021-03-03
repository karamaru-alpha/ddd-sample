package user

import (
	"errors"
)

// HashedPassword Value object express hashed user's passsword
type HashedPassword []byte

// NewHashedPassword Constructor generate user's HashedPassword value object
func NewHashedPassword(arg []byte) (*HashedPassword, error) {

	if arg == nil {
		return nil, errors.New("UserHashedPassword is null")
	}

	hashedPassword := HashedPassword(arg)
	return &hashedPassword, nil
}

// ToString Convert UserHashedPassword to String
func (hashedPassword HashedPassword) ToString() string {
	return string(hashedPassword)
}
