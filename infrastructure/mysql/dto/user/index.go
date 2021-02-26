package user

import (
	"errors"

	entity "github.com/karamaru-alpha/ddd-sample/domain/entity/user"
	valueObject "github.com/karamaru-alpha/ddd-sample/domain/value_object/user"
	"github.com/oklog/ulid"
)

// User DTO about user entity
type User struct {
	ID          string `gorm:"primary_key"`
	Name        string
	MailAddress string
}

// NewUser Constructor generate DTO of User
func NewUser(id string, name string, mailAddress string) (*User, error) {
	if id == "" {
		return nil, errors.New("UserID is null")
	}
	if name == "" {
		return nil, errors.New("UserName is null")
	}
	if mailAddress == "" {
		return nil, errors.New("UserMailAddress is null")
	}

	user := new(User)
	user.ID = id
	user.Name = name
	user.MailAddress = mailAddress

	return user, nil
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

// ConvertDTO Convert Entity to DTO about User
func ConvertDTO(entityUser *entity.User) (*User, error) {
	DTOUserID := entityUser.ID.ToString()
	DTOUserName := entityUser.Name.ToString()
	DTOUserMailAddress := entityUser.MailAddress.ToString()

	DTOUser, err := NewUser(DTOUserID, DTOUserName, DTOUserMailAddress)
	if err != nil {
		return nil, err
	}

	return DTOUser, nil
}
