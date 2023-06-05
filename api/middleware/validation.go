package middleware

import (
	"Song_API/api/models"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ValidateSong(models.Song, bool) validates the fields of the Song struct.
func ValidateSong(song models.Song, valRequired bool) error {
	if !valRequired {
		return validation.ValidateStruct(&song,
			validation.Field(&song.Plays, validation.Min(0)),
			validation.Field(&song.ReleaseDate, validation.Date(time.DateOnly)))
	}
	return validation.ValidateStruct(&song,
		validation.Field(&song.Song, validation.Required),
		validation.Field(&song.Artist, validation.Required),
		validation.Field(&song.Plays, validation.Min(0)),
		validation.Field(&song.ReleaseDate, validation.Required, validation.Date(time.DateOnly)))
}

// ValidateAccount(models.Account) validates the fields of the Account struct.
func ValidateAccount(account models.Account) error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.User, validation.Required, is.Alphanumeric, validation.Length(4, 20)),
		validation.Field(&account.Password, validation.Required, validation.Length(8, 20)))
}

// Validation() is a middleware function that validates the request body.
func Validation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
			c.Abort()
			return
		}
		path := c.Request.URL.Path
		grp := strings.Split(path, "/")[2]
		if grp == "songs" {
			var song models.Song
			databytes, _ := json.Marshal(data)
			json.Unmarshal(databytes, &song)
			required := false
			if method := c.Request.Method; method == "POST" {
				required = true
			}
			if err := ValidateSong(song, required); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		} else if grp == "accounts" {
			var acc models.Account
			databytes, _ := json.Marshal(data)
			json.Unmarshal(databytes, &acc)
			if err := ValidateAccount(acc); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}
		c.Set("body", data)
		c.Next()
	}
}
