package middleware

import (
	"Song_API/api/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ValidateSong(models.Song) validates the fields of the Song struct.
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

func Validation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
			c.Abort()
			return
		}
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
		c.Set("body", data)
		c.Next()
	}
}
