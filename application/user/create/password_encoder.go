package user

import (
	"golang.org/x/crypto/bcrypt"

	domainModel "github.com/karamaru-alpha/ddd-sample/domain/model/user"
)

// PasswordEncoder Encode plain password to hash
func PasswordEncoder(plainPassword domainModel.PlainPassword) (domainModel.HashedPassword, error) {
	return bcrypt.GenerateFromPassword([]byte(string(plainPassword)), bcrypt.DefaultCost)
}
