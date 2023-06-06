package validation

import (
	"Song_API/pkg/models"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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
