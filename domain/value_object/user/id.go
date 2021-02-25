package user

import (
	"errors"
)

// ID Value object express user's ID
type ID string

// NewID Constructor generate user's ID value object
func NewID(arg string) (*ID, error) {

	if arg == "" {
		return nil, errors.New("UserID is null")
	}

	id := ID(arg)
	return &id, nil
}
