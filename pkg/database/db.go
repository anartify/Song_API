package database

import (
	"Song_API/pkg/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// GetDB() returns the pointer to the database.
func GetDB() *gorm.DB {
	return db
}

// Connect() connects the mysql database and automigrates the tables.
func Connect() {
	var err error
	db, err = gorm.Open("mysql", DbUrl())
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.Song{})
	db.AutoMigrate(&models.Account{})
}
