package user

import (
	"errors"

	"github.com/oklog/ulid"
)

// ID Value object express user's ID
type ID ulid.ULID

// NewID Constructor generate user's ID value object
func NewID(arg ulid.ULID) (*ID, error) {
	var zeroValueULID ulid.ULID

	if arg.Compare(zeroValueULID) == 0 {
		return nil, errors.New("UserID is null")
	}

	id := ID(arg)
	return &id, nil
}
