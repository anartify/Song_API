package test

import (
	"Song_API/api/middleware"
	"Song_API/api/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSongValidation() tests the Validation() method of the Song struct.
func TestSongValidation(t *testing.T) {
	assert := assert.New(t)
	song := models.Song{
		Song:        "My Song",
		Artist:      "My Artist",
		Plays:       10,
		ReleaseDate: "2022-01-01",
	}

	song.Song = ""
	err := middleware.ValidateSong(song, true)
	assert.Error(err, "validation should fail for empty Song")

	song.Song = "My Song"
	song.Artist = ""
	err = middleware.ValidateSong(song, true)
	assert.Error(err, "validation should fail for empty Artist")

	song.Artist = "My Artist"
	song.Plays = -1
	err = middleware.ValidateSong(song, true)
	assert.Error(err, "validation should fail for negative Plays")

	song.Plays = 10
	song.ReleaseDate = ""
	err = middleware.ValidateSong(song, true)
	assert.Error(err, "validation should fail for empty ReleaseDate")

	song.ReleaseDate = "2022-01-01 00:00:00"
	err = middleware.ValidateSong(song, true)
	assert.Error(err, "validation should fail for invalid ReleaseDate")

	song.ReleaseDate = "2022-01-01"
	err = middleware.ValidateSong(song, true)
	assert.NoError(err, "validation should pass")

	account := models.Account{
		User:     "anartify",
		Password: "kyuBatau",
	}
	account.User = ""
	err = middleware.ValidateAccount(account)
	assert.Error(err, "validation should fail for user with length less than 4")

	account.User = "4n_4rt1fy"
	err = middleware.ValidateAccount(account)
	assert.Error(err, "validation should fail if user contains non-alphanumeric characters")

	account.User = "anartify"
	account.Password = ""
	err = middleware.ValidateAccount(account)
	assert.Error(err, "Validation should fail for password with length less than 8")

	account.Password = "kyuBt44u"
	err = middleware.ValidateAccount(account)
	assert.NoError(err, "Validation should pass")
}
