package user

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// IFactory Interface of Factory help to generate User Entity
type IFactory interface {
	Create(*Name, *MailAddress, *PlainPassword) (*User, error)
}

type factory struct{}

// NewFactory create UserFactory.
func NewFactory() IFactory {
	return &factory{}
}

// Create Factory help to generate UserEntity.
func (factory) Create(name *Name, mailAddress *MailAddress, plainPassword *PlainPassword) (*User, error) {
	// ID
	generatedULID, err := generateULID()
	if err != nil {
		return nil, err
	}
	userID, err := NewID(generatedULID)
	if err != nil {
		return nil, err
	}

	// HashPassword
	hashedPassword, err := plainPassword.ToHash()
	if err != nil {
		return nil, err
	}

	userHashedPassword, err := NewHashedPassword(hashedPassword)
	if err != nil {
		return nil, err
	}

	return NewUser(userID, name, mailAddress, userHashedPassword)
}

func generateULID() (ulid.ULID, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.New(ulid.Timestamp(t), entropy)
}
