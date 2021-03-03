package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/karamaru-alpha/ddd-sample/constants"
	domainModel "github.com/karamaru-alpha/ddd-sample/domain/model/user"
)

// GenerateJWT Generate JWT token
func GenerateJWT(userID *domainModel.ID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte(constants.JWTSecretKey))
}
