package main

import (
	"Song_API/Config"
	"Song_API/Models"
	"Song_API/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.Db_url())
	if err != nil {
		panic(err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Song{})
	r := Routes.Initialize()
	r.Run()
}
