package repository

import (
	"Song_API/api/models"
	"Song_API/database"
)

// SongInterface is an interface that defines all the helper methods required by controller functions.
type SongInterface interface {
	GetAllSong(song *[]models.Song, user string) error
	AddSong(song *models.Song, user string) error
	GetSong(song *models.Song, id string, user string) error
	UpdateSong(song *models.Song) error
	DeleteSong(song *models.Song, id string, user string) error
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

// GetAllSong(*[]models.Song, string) gets all songs from database and returns error if any
func (sr SongRepo) GetAllSong(song *[]models.Song, user string) error {
	if err := database.DB.Find(song, "user = ?", user).Error; err != nil {
		return &CustomError{message: "No data found"}
	}
	return nil
}

// AddSong(*models.Song, string) adds a song in database and returns error if any
func (sr SongRepo) AddSong(song *models.Song, user string) error {
	song.User = user
	if err := database.DB.Create(song).Error; err != nil {
		return &CustomError{message: "Failed to add data"}
	}
	return nil
}

// GetSong(*models.Song, string, string) gets a song from database and returns error if any
func (sr SongRepo) GetSong(song *models.Song, id string, user string) error {
	if err := database.DB.Where("id = ? AND user = ?", id, user).First(song).Error; err != nil {
		return &CustomError{message: "No data found"}
	}
	return nil
}

// UpdateSong(*models.Song) updates a song in database and returns error if any
func (sr SongRepo) UpdateSong(song *models.Song) error {
	database.DB.Save(song)
	return nil
}

// DeleteSong(*models.Song, string, string) deletes a song from database and returns error if any
func (sr SongRepo) DeleteSong(song *models.Song, id string, user string) error {
	resp := database.DB.Where("id = ? AND user = ?", id, user).Delete(song)
	if resp.RowsAffected == 0 {
		return &CustomError{message: "No record Found"}
	}
	return nil
}
