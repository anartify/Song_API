package database

import (
	"fmt"

	"github.com/spf13/viper"
)

func DbUrl() string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		viper.GetString("USER"),
		viper.GetString("PASSWORD"),
		viper.GetString("HOST"),
		viper.GetInt("PORT"),
		viper.GetString("DB_NAME"),
	)
}
