package user

import (
	"golang.org/x/crypto/bcrypt"

	domainModel "github.com/karamaru-alpha/ddd-sample/domain/model/user"
)

// PasswordEncoder Encode plain password to hash about user
func PasswordEncoder(plainPassword domainModel.PlainPassword) (*domainModel.HashedPassword, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(string(plainPassword)), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return domainModel.NewHashedPassword(hashedPassword)
}
