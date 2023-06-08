package validation

import (
	"Song_API/pkg/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ValidateAccount(models.Account) validates the fields of the Account struct.
func ValidateAccount(account models.Account) error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.User, validation.Required, is.Alphanumeric, validation.Length(4, 20)),
		validation.Field(&account.Password, validation.Required, validation.Length(8, 20)))
}
