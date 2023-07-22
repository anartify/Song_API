package validation

import (
	"Song_API/pkg/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ValidateCreateAccount function validates user and password field for CreateAccount controller
func ValidateCreateAccount(account models.Account) error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.User, validation.Required, is.Alphanumeric, validation.Length(4, 20)),
		validation.Field(&account.Password, validation.Required, is.Alphanumeric, validation.Length(8, 20)))
}

// ValidateGetAccount function validates user and password field for GetAccount controller
func ValidateGetAccount(account models.Account) error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.User, validation.Required, is.Alphanumeric, validation.Length(4, 20)),
		validation.Field(&account.Password, validation.Required, is.Alphanumeric, validation.Length(8, 20)))
}

// ValidateUpdateRole function validates user and role field for UpdateRole controller
func ValidateUpdateRole(account models.Account) error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.User, validation.Required, is.Alphanumeric, validation.Length(4, 20)),
		validation.Field(&account.Role, validation.Required, validation.In("general", "admin")))
}
