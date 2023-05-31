package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Song struct holds the fields required in the Table songs.
type Song struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Song        string `json:"song"`
	Artist      string `json:"artist"`
	Plays       int    `json:"plays"`
	ReleaseDate string `json:"release_date"`
}

// TableName returns the name of the table in the database.
func (s *Song) TableName() string {
	return "songs"
}

// Validation validates the fields of the Song struct.
func (data Song) Validation() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Song, validation.Required),
		validation.Field(&data.Artist, validation.Required),
		validation.Field(&data.Plays, validation.Required),
		validation.Field(&data.ReleaseDate, validation.Required, validation.Date(time.DateOnly)),
	)
}
