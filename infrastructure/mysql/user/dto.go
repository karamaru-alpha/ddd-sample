package user

import (
	"github.com/oklog/ulid"

	domainModel "github.com/karamaru-alpha/ddd-sample/domain/model/user"
)

// User DTO about user entity
type User struct {
	ID             string `gorm:"primary_key"`
	Name           string
	MailAddress    string
	HashedPassword string
}

// ConvertDTO Convert Entity to DTO about User
func ConvertDTO(entityUser *domainModel.User) *User {
	DTOUserID := ulid.ULID(entityUser.ID).String()
	DTOUserName := string(entityUser.Name)
	DTOUserMailAddress := string(entityUser.MailAddress)
	DTOUserHashedPassword := string(entityUser.HashedPassword)

	return &User{
		ID:             DTOUserID,
		Name:           DTOUserName,
		MailAddress:    DTOUserMailAddress,
		HashedPassword: DTOUserHashedPassword,
	}
}

// ConvertEntity Convert DTO to Entity about User
func ConvertEntity(DTOUser *User) (*domainModel.User, error) {
	parsedULID, err := ulid.Parse(DTOUser.ID)
	if err != nil {
		return nil, err
	}

	entityUserID, err := domainModel.NewID(parsedULID)
	if err != nil {
		return nil, err
	}

	entityUserName, err := domainModel.NewName(DTOUser.Name)
	if err != nil {
		return nil, err
	}

	entityUserMailAddress, err := domainModel.NewMailAddress(DTOUser.MailAddress)
	if err != nil {
		return nil, err
	}

	// TODO Prevent from memory copy
	entityUserHashedPassword, err := domainModel.NewHashedPassword([]byte(DTOUser.HashedPassword))
	if err != nil {
		return nil, err
	}

	return domainModel.NewUser(
		entityUserID,
		entityUserName,
		entityUserMailAddress,
		entityUserHashedPassword,
	)
}
