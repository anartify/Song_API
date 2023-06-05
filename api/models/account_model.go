package models

// Account struct defines the structure of an account.
type Account struct {
	ID       int    `json:"id" gorm:"primary_key"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// TableName returns the name of the table in the database.
func (a *Account) TableName() string {
	return "accounts"
}
