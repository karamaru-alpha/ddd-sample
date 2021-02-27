package user

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// IFactory Interface of Factory help to generate User Entity
type IFactory interface {
	Create(*Name, *MailAddress) (*User, error)
}

type factory struct{}

// NewFactory create UserFactory.
func NewFactory() IFactory {
	return &factory{}
}

// Create Factory help to generate UserEntity.
func (factory) Create(name *Name, mailAddress *MailAddress) (*User, error) {

	generatedULID := generateULID()

	userID, err := NewID(generatedULID)
	if err != nil {
		return nil, err
	}

	user, err := NewUser(userID, name, mailAddress)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func generateULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}
