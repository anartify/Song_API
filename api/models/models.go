package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Song struct {
	ID          int    `json:"id"`
	Song        string `json:"song"`
	Artist      string `json:"artist"`
	Plays       int    `json:"plays"`
	ReleaseDate string `json:"release_date"`
}

func (s *Song) TableName() string {
	return "songs"
}

func (data Song) Validation() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Song, validation.Required),
		validation.Field(&data.Artist, validation.Required),
		validation.Field(&data.Plays, validation.Required),
		validation.Field(&data.ReleaseDate, validation.Required, validation.Date(time.DateOnly)),
	)
}
