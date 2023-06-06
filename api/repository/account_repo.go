package repository

import (
	"Song_API/api/models"
	"Song_API/api/utils"
	"Song_API/database"

	"golang.org/x/crypto/bcrypt"
)

// AccountInterface is an interface that defines all the helper methods required by controller functions.
type AccountInterface interface {
	CreateAccount(*models.Account) error
	GetAccount(*models.Account) (string, error)
}

// AccountRepo struct has the implementation of  all the methods of AccountInterface.
type AccountRepo struct{}

// CreateAccount(*models.Account) creates an account in database and returns error if any
func (ar AccountRepo) CreateAccount(account *models.Account) error {
	if err := database.DB.Where("user = ?", account.User).First(account).Error; err == nil {
		return &CustomError{message: "Account already exists"}
	}
	password := account.Password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	if err := database.DB.Create(account).Error; err != nil {
		return &CustomError{message: "failed to create account"}
	}
	account.Password = password
	return nil
}

// GetAccount(*models.Account) gets an account from database and returns authentication token and error if any
func (ar AccountRepo) GetAccount(account *models.Account) (string, error) {
	password := account.Password
	if err := database.DB.Where("user = ?", account.User).First(account).Error; err != nil {
		return "", &CustomError{message: "No account found"}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		return "", &CustomError{message: "Invalid password"}
	}
	account.Password = password
	token, _ := utils.GenerateToken(account)
	return token, nil
}
