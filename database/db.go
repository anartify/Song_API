package database

import (
	"Song_API/api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// Connect() connects the mysql database and automigrates the tables.
func Connect() {
	var err error
	DB, err = gorm.Open("mysql", DbUrl())
	if err != nil {
		panic("failed to connect to database")
	}
	DB.AutoMigrate(&models.Song{})
}
