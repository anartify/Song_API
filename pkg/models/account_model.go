package models

// Account struct defines the structure of an account.
type Account struct {
	ID       int    `json:"id" gorm:"primary_key"`
	User     string `json:"user"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// TableName returns the name of the table in the database.
func (a *Account) TableName() string {
	return "accounts"
}

// SetID sets the ID property of the Account model.
func (a *Account) SetID(id int) {
	a.ID = id
}

// GetID returns the ID property of the Account model.
func (a *Account) GetID() int {
	return a.ID
}

// SetUser sets the User property of the Account model.
func (a *Account) SetUser(user string) {
	a.User = user
}

// GetUser returns the User property of the Account model.
func (a *Account) GetUser() string {
	return a.User
}

// SetPassword sets the Password property of the Account model.
func (a *Account) SetPassword(password string) {
	a.Password = password
}

// GetPassword returns the Password property of the Account model.
func (a *Account) GetPassword() string {
	return a.Password
}

// SetRole sets the Role property of the Account model.
func (a *Account) SetRole(role string) {
	a.Role = role
}

// GetRole returns the Role property of the Account model.
func (a *Account) GetRole() string {
	return a.Role
}
