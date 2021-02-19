package user

import (
	"errors"

	"github.com/google/uuid"
)

// ID Value object express user's ID.
type ID uuid.UUID

// NewID Constructor generate user's ID value object.
func NewID(arg uuid.UUID) (*ID, error) {

	if arg == uuid.Nil {
		return nil, errors.New("UserID is null")
	}

	id := ID(arg)
	return &id, nil
}
