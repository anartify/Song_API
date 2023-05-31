package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func GenerateToken(user string) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the admin user ID and role as claims in the token's payload
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	signingKey := []byte(viper.GetString("AUTH_KEY"))

	generatedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return generatedToken, nil
}
