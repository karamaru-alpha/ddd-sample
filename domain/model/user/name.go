package user

import (
	"errors"
	"unicode/utf8"
)

// Name Value object express user's name
type Name string

// NewName Constructor generate user's Name value object
func NewName(arg string) (*Name, error) {

	if arg == "" {
		return nil, errors.New("UserName is null")
	}

	if utf8.RuneCountInString(arg) < 3 || 20 < utf8.RuneCountInString(arg) {
		return nil, errors.New("UserName should be Three to twenty characters")
	}

	name := Name(arg)
	return &name, nil
}

// ToString Convert UserName to String
func (name Name) ToString() string {
	return string(name)
}
