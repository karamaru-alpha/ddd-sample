package user

import (
	"github.com/oklog/ulid"

	entity "github.com/karamaru-alpha/ddd-sample/domain/entity/user"
	valueObject "github.com/karamaru-alpha/ddd-sample/domain/value_object/user"
)

// User DTO about user entity
type User struct {
	ID          string `gorm:"primary_key"`
	Name        string
	MailAddress string
}

// ConvertDTO Convert Entity to DTO about User
func ConvertDTO(entityUser *entity.User) *User {
	DTOUserID := entityUser.ID.ToString()
	DTOUserName := entityUser.Name.ToString()
	DTOUserMailAddress := entityUser.MailAddress.ToString()

	return &User{
		ID:          DTOUserID,
		Name:        DTOUserName,
		MailAddress: DTOUserMailAddress,
	}
}

// ConvertEntity Convert DTO to Entity about User
func ConvertEntity(DTOUser *User) (*entity.User, error) {
	parsedULID, err := ulid.Parse(DTOUser.ID)
	if err != nil {
		return nil, err
	}

	entityUserID, err := valueObject.NewID(parsedULID)
	if err != nil {
		return nil, err
	}

	entityUserName, err := valueObject.NewName(DTOUser.Name)
	if err != nil {
		return nil, err
	}

	entityUserMailAddress, err := valueObject.NewMailAddress(DTOUser.MailAddress)
	if err != nil {
		return nil, err
	}

	entityUser, err := entity.NewUser(entityUserID, entityUserName, entityUserMailAddress)
	if err != nil {
		return nil, err
	}

	return entityUser, nil
}
