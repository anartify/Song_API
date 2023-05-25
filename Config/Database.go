// Config/Database.go
package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

type Profile struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func DBConfig() *Profile {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	dbConfig := Profile{
		Host:     "localhost",
		Port:     3306,
		User:     viper.Get("DB_USER").(string),
		Password: viper.Get("DB_PASSWORD").(string),
		DBName:   "SongDB",
	}
	return &dbConfig
}

func Db_url() string {
	profile := DBConfig()
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		profile.User,
		profile.Password,
		profile.Host,
		profile.Port,
		profile.DBName,
	)
}
