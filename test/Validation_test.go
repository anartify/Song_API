package test

import (
	"Song_API/pkg/controllers/validation"
	"Song_API/pkg/models"
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
	err := validation.ValidateSong(song, true)
	assert.Error(err, "validation should fail for empty Song")

	song.Song = "My Song"
	song.Artist = ""
	err = validation.ValidateSong(song, true)
	assert.Error(err, "validation should fail for empty Artist")

	song.Artist = "My Artist"
	song.Plays = -1
	err = validation.ValidateSong(song, true)
	assert.Error(err, "validation should fail for negative Plays")

	song.Plays = 10
	song.ReleaseDate = ""
	err = validation.ValidateSong(song, true)
	assert.Error(err, "validation should fail for empty ReleaseDate")

	song.ReleaseDate = "2022-01-01 00:00:00"
	err = validation.ValidateSong(song, true)
	assert.Error(err, "validation should fail for invalid ReleaseDate")

	song.ReleaseDate = "2022-01-01"
	err = validation.ValidateSong(song, true)
	assert.NoError(err, "validation should pass")

	account := models.Account{
		User:     "anartify",
		Password: "kyuBatau",
		Role:     "general",
	}

	account.User = ""
	err = validation.ValidateCreateAccount(account)
	assert.Error(err, "validation should fail for user with length less than 4")
	err = validation.ValidateGetAccount(account)
	assert.Error(err, "validation should fail for user with length less than 4")
	err = validation.ValidateUpdateRole(account)
	assert.Error(err, "validation should fail for user with length less than 4")

	account.User = "4n_4rt1fy"
	err = validation.ValidateCreateAccount(account)
	assert.Error(err, "validation should fail if user contains non-alphanumeric characters")
	err = validation.ValidateGetAccount(account)
	assert.Error(err, "validation should fail if user contains non-alphanumeric characters")
	err = validation.ValidateUpdateRole(account)
	assert.Error(err, "validation should fail if user contains non-alphanumeric characters")

	account.Password = ""
	err = validation.ValidateCreateAccount(account)
	assert.Error(err, "validation should fail for password with length less than 8")
	err = validation.ValidateGetAccount(account)
	assert.Error(err, "validation should fail for password with length less than 8")

	account.Role = "guest"
	err = validation.ValidateUpdateRole(account)
	assert.Error(err, "Validation should fail for invalid role")
}
