package user

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// IFactory Interface of Factory help to generate User Entity
type IFactory interface {
	Create(*Name, *MailAddress, *HashedPassword) (*User, error)
}

type factory struct{}

// NewFactory create UserFactory.
func NewFactory() IFactory {
	return &factory{}
}

// Create Factory help to generate UserEntity.
func (factory) Create(name *Name, mailAddress *MailAddress, hashedPassword *HashedPassword) (*User, error) {
	// ID
	generatedULID, err := generateULID()
	if err != nil {
		return nil, err
	}
	userID, err := NewID(generatedULID)
	if err != nil {
		return nil, err
	}

	return NewUser(userID, name, mailAddress, hashedPassword)
}

func generateULID() (ulid.ULID, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.New(ulid.Timestamp(t), entropy)
}
