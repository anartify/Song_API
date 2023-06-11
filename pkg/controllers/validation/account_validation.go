package validation

import (
	"Song_API/pkg/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ValidateUser function validates user field of account
func ValidateUser(account models.Account) error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.User, validation.Required, is.Alphanumeric, validation.Length(4, 20)))
}

// ValidatePassword function validates password field of account
func ValidatePassword(account models.Account) error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.Password, validation.Required, is.Alphanumeric, validation.Length(8, 20)))
}

// ValidateRole function validates role field of account
func ValidateRole(account models.Account) error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.Role, validation.Required, validation.In("general", "admin")))
}
