package utils

import (
	"Song_API/pkg/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// GenerateToken generates a JWT token for the given account.
func GenerateToken(account *models.Account) (string, time.Duration, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = account.User
	claims["password"] = account.Password
	claims["role"] = account.Role
	expTime := time.Now().Add(time.Hour * 24).Unix()
	claims["exp"] = expTime // standard claim
	remTime := time.Until(time.Unix(int64(expTime), 0))
	signingKey := []byte(viper.GetString("AUTH_KEY"))
	generatedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", remTime, err
	}
	return generatedToken, remTime, nil
}
