package repository

import (
	"Song_API/api/models"
	"Song_API/database"
)

// SongInterface is an interface that defines all the helper methods required by controller functions.
type SongInterface interface {
	GetAllSong(b *[]models.Song) error
	AddSong(b *models.Song) error
	GetSong(b *models.Song, id string) error
	UpdateSong(b *models.Song) error
	DeleteSong(b *models.Song, id string) error
}

// SongRepo struct has the implementation of  all the methods of SongInterface.
type SongRepo struct{}

// This struct helps in writing custom error messages
type CustomError struct {
	message string
}

// Error() method returns the string containing the error message
func (e *CustomError) Error() string {
	return e.message
}

// GetAllSong(*[]models.Song) gets all songs from database and returns error if any
func (sr SongRepo) GetAllSong(b *[]models.Song) error {
	if err := database.DB.Find(b).Error; err != nil {
		return &CustomError{message: "No data found"}
	}
	return nil
}

// AddSong(*models.Song) adds a song in database and returns error if any
func (sr SongRepo) AddSong(b *models.Song) error {
	if err := b.Validation(); err != nil {
		return &CustomError{message: err.Error()}
	}
	if err := database.DB.Create(b).Error; err != nil {
		return &CustomError{message: "Failed to add data"}
	}
	return nil
}

// GetSong(*models.Song, id string) gets a song from database and returns error if any
func (sr SongRepo) GetSong(b *models.Song, id string) error {
	if err := database.DB.Where("id = ?", id).First(b).Error; err != nil {
		return &CustomError{message: "No data found"}
	}
	return nil
}

// UpdateSong(*models.Song) updates a song in database and returns error if any
func (sr SongRepo) UpdateSong(b *models.Song) error {
	if err := b.Validation(); err != nil {
		return &CustomError{message: err.Error()}
	}
	database.DB.Save(b)
	return nil
}

// DeleteSong(*models.Song, id string) deletes a song from database and returns error if any
func (sr SongRepo) DeleteSong(b *models.Song, id string) error {
	database.DB.Where("id = ?", id).Delete(b)
	return nil
}
