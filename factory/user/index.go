package user

import (
	"github.com/google/uuid"

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

	generatedUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	userID, err := valueObject.NewID(generatedUUID)
	if err != nil {
		return nil, err
	}

	user, err := entity.NewUser(userID, name, mailAdress)
	if err != nil {
		return nil, err
	}

	return user, nil
}
