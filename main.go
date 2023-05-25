package main

import (
	"Song_API/config"
	"Song_API/models"
	"Song_API/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.Db_url())
	if err != nil {
		panic(err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Song{})
	r := routes.Initialize()
	r.Run()
}
