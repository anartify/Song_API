package repository

import (
	"Song_API/pkg/apperror"
	"Song_API/pkg/database"
	"Song_API/pkg/models"
	"Song_API/pkg/utils"

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
	if err := database.GetDB().Where("user = ?", account.GetUser()).First(account).Error; err == nil {
		return &apperror.CustomError{Message: "Account already exists"}
	}
	password := account.GetPassword()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	account.SetPassword(string(hashedPassword))
	if err := database.GetDB().Create(account).Error; err != nil {
		return &apperror.CustomError{Message: "failed to create account"}
	}
	account.SetPassword(password)
	return nil
}

// GetAccount(*models.Account) gets an account from database and returns authentication token and error if any
func (ar AccountRepo) GetAccount(account *models.Account) (string, error) {
	password := account.GetPassword()
	if err := database.GetDB().Where("user = ?", account.GetUser()).First(account).Error; err != nil {
		return "", &apperror.CustomError{Message: "No account found"}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.GetPassword()), []byte(password)); err != nil {
		return "", &apperror.CustomError{Message: "Invalid password"}
	}
	account.SetPassword(password)
	token, _ := utils.GenerateToken(account)
	return token, nil
}
