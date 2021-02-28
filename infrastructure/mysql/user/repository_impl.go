package user

import (
	"errors"

	"github.com/jinzhu/gorm"

	domainModel "github.com/karamaru-alpha/ddd-sample/domain/model/user"
)

type repositoryImpl struct {
	db *gorm.DB
}

// NewRepositoryImpl Function create infrastructure to implement user repository
func NewRepositoryImpl(db *gorm.DB) domainModel.IRepository {
	return &repositoryImpl{
		db,
	}
}

// Save Create or Update user object
func (uri repositoryImpl) Save(user *domainModel.User) error {
	DTOUser := ConvertDTO(user)

	return uri.db.Save(DTOUser).Error
}

// TODO implement receive many argument
// Find Rebuild user object by mail address
func (uri repositoryImpl) Find(mailAddress *domainModel.MailAddress) (*domainModel.User, error) {
	var DTOUser User

	if err := uri.db.Where("mail_address = ?", mailAddress.ToString()).First(&DTOUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return ConvertEntity(&DTOUser)
}
