package utils

import (
	"Song_API/pkg/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// GenerateToken generates a JWT token for the given account.
func GenerateToken(account *models.Account) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = account.User
	claims["password"] = account.Password
	claims["role"] = account.Role
	signingKey := []byte(viper.GetString("AUTH_KEY"))
	generatedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return generatedToken, nil
}
