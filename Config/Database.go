// Config/Database.go
package Config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// DBConfig represents db configuration
type Profile struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func DBConfig() *Profile {
	err := godotenv.Load()
	if err != nil {
		// Handle error if .env file is not found or cannot be loaded
		panic("Failed to load .env file")
	}
	dbConfig := Profile{
		Host:     "localhost",
		Port:     3306,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   "SongDB",
	}
	return &dbConfig
}
func Db_url() string {
	profile := DBConfig()
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		profile.User,
		profile.Password,
		profile.Host,
		profile.Port,
		profile.DBName,
	)
}
