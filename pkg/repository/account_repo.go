package repository

import (
	"Song_API/pkg/apperror"
	"Song_API/pkg/database"
	"Song_API/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

// AccountInterface is an interface that defines all the helper methods required by controller functions.
type AccountInterface interface {
	CreateAccount(*models.Account) error
	GetAccount(*models.Account) error
	GetAllAccount(*[]models.Account) error
	UpdateRole(*models.Account) error
	DeleteAccount(*models.Account) error
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

// GetAccount(*models.Account) gets an account from database and returns error if any
func (ar AccountRepo) GetAccount(account *models.Account) error {
	password := account.GetPassword()
	if err := database.GetDB().Where("user = ?", account.GetUser()).First(account).Error; err != nil {
		return &apperror.CustomError{Message: "No account found"}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.GetPassword()), []byte(password)); err != nil {
		return &apperror.CustomError{Message: "Invalid password"}
	}
	account.SetPassword(password)
	return nil
}

// GetAllAccount(*[]models.Account) populates user and role field of the models.Account struct from the database and returns error if any
func (ar AccountRepo) GetAllAccount(acc *[]models.Account) error {
	var account models.Account
	if err := database.GetDB().Table(account.TableName()).Select("user, role").Find(&acc).Error; err != nil {
		return &apperror.CustomError{Message: "No account found"}
	}
	return nil
}

// UpdateRole(*models.Account) updates the role of the user and returns error if any
func (ar AccountRepo) UpdateRole(acc *models.Account) error {
	if err := database.GetDB().Model(acc).Where("user = ?", acc.GetUser()).Update("role", acc.GetRole()).Error; err != nil {
		return &apperror.CustomError{Message: "User not found"}
	}
	return nil
}

func (ar AccountRepo) DeleteAccount(acc *models.Account) error {
	resp := database.GetDB().Where("user = ?", acc.GetUser()).Delete(acc)
	if resp.RowsAffected == 0 {
		return &apperror.CustomError{Message: "No User Found"}
	}
	return nil
}
