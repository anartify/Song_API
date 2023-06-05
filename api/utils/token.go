package utils

import (
	"Song_API/api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func GenerateToken(account *models.Account) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = account.User
	claims["password"] = account.Password
	signingKey := []byte(viper.GetString("AUTH_KEY"))
	generatedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return generatedToken, nil
}
