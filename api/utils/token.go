package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func GenerateToken(user string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	signingKey := []byte(viper.GetString("AUTH_KEY"))
	generatedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return generatedToken, nil
}
