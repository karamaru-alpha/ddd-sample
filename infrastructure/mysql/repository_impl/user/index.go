package user

import (
	"errors"

	"github.com/jinzhu/gorm"
	entity "github.com/karamaru-alpha/ddd-sample/domain/entity/user"
	repository "github.com/karamaru-alpha/ddd-sample/domain/repository/user"
	valueObject "github.com/karamaru-alpha/ddd-sample/domain/value_object/user"

	dto "github.com/karamaru-alpha/ddd-sample/infrastructure/mysql/dto/user"
)

type repositoryImpl struct {
	db *gorm.DB
}

// NewRepositoryImpl Function create infrastructure to implement user repository
func NewRepositoryImpl(db *gorm.DB) repository.IRepository {
	return &repositoryImpl{
		db,
	}
}

// Save Persist user object
func (uri repositoryImpl) Save(user *entity.User) error {
	DTOUser, err := dto.ConvertDTO(user)
	if err != nil {
		return err
	}

	if err = uri.db.Create(DTOUser).Error; err != nil {
		return err
	}

	return nil
}

// TODO implement receive many argument
// Find Rebuild user object by mail address
func (uri repositoryImpl) Find(mailAddress *valueObject.MailAddress) (*entity.User, error) {
	var DTOUser dto.User

	if err := uri.db.First(DTOUser, "mail_address = ?", mailAddress.ToString()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	user, err := dto.ConvertEntity(&DTOUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}
