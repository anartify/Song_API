package database

import (
	"fmt"

	"github.com/spf13/viper"
)

// DbUrl() reads the .env file and returns the database url.
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

// SongCache() reads the .env file and returns the song cache details (host, port, db, expire, password).
func SongCache() (string, int, int, int, string) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString("REDIS_HOST"), viper.GetInt("REDIS_PORT"), viper.GetInt("SONG_CACHE_DB"), viper.GetInt("SONG_CACHE_EXPIRE"), viper.GetString("REDIS_PASSWORD")
}

// AccountCache() reads the .env file and returns the account cache details (host, port, db, expire, password).
func AccountCache() (string, int, int, int, string) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString("REDIS_HOST"), viper.GetInt("REDIS_PORT"), viper.GetInt("ACCOUNT_CACHE_DB"), viper.GetInt("ACCOUNT_CACHE_EXPIRE"), viper.GetString("REDIS_PASSWORD")
}

// BucketCache() reads the .env file and returns the bucket cache details (host, port, db, expire, password).
func BucketCache() (string, int, int, int, string) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetString("REDIS_HOST"), viper.GetInt("REDIS_PORT"), viper.GetInt("BUCKET_CACHE_DB"), viper.GetInt("BUCKET_CACHE_EXPIRE"), viper.GetString("REDIS_PASSWORD")
}
