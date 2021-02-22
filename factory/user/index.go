package user

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"

	entity "github.com/karamaru-alpha/ddd-sample/domain/entity/user"
	valueObject "github.com/karamaru-alpha/ddd-sample/domain/value_object/user"
)

// IFactory Interface of Factory help to generate Object.
type IFactory interface {
	Create(name *valueObject.Name, mailAdress *valueObject.MailAdress) (*entity.User, error)
}

type factory struct{}

// NewFactory create UserFactory.
func NewFactory() IFactory {
	return &factory{}
}

// Create Factory help to generate UserEntity.
func (factory) Create(name *valueObject.Name, mailAdress *valueObject.MailAdress) (*entity.User, error) {

	generatedULID := generateULID()

	userID, err := valueObject.NewID(generatedULID)
	if err != nil {
		return nil, err
	}

	user, err := entity.NewUser(userID, name, mailAdress)
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
