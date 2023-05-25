package test

import (
	"Song_API/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSongValidation(t *testing.T) {
	assert := assert.New(t)
	song := models.Song{
		Song:        "My Song",
		Artist:      "My Artist",
		Plays:       10,
		ReleaseDate: "2022-01-01",
	}

	song.Song = ""
	err := song.Validation()
	assert.Error(err, "validation should fail for empty Song")

	song.Song = "My Song"
	song.Artist = ""
	err = song.Validation()
	assert.Error(err, "validation should fail for empty Artist")

	song.Artist = "My Artist"
	song.Plays = 0
	err = song.Validation()
	assert.Error(err, "validation should fail for empty Plays")

	song.Plays = 10
	song.ReleaseDate = ""
	err = song.Validation()
	assert.Error(err, "validation should fail for empty ReleaseDate")

	song.ReleaseDate = "2022-01-01 00:00:00"
	err = song.Validation()
	assert.Error(err, "validation should fail for invalid ReleaseDate")

	song.ReleaseDate = "2022-01-01"
	err = song.Validation()
	assert.NoError(err, "validation should pass")
}
