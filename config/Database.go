package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

type profile struct {
	host     string
	port     int
	user     string
	dbName   string
	password string
}

func dbConfig() *profile {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	dbConfig := profile{
		host:     viper.GetString("HOST"),
		port:     viper.GetInt("PORT"),
		user:     viper.GetString("USER"),
		password: viper.GetString("PASSWORD"),
		dbName:   viper.GetString("DB_NAME"),
	}
	return &dbConfig
}

func Db_url() string {
	conf := dbConfig()
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		conf.user,
		conf.password,
		conf.host,
		conf.port,
		conf.dbName,
	)
}
