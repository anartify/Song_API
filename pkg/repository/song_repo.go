package repository

import (
	"Song_API/pkg/apperror"
	"Song_API/pkg/database"
	"Song_API/pkg/models"
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

// GetAllSong(*[]models.Song, string) gets all songs from database and returns error if any
func (sr SongRepo) GetAllSong(song *[]models.Song, user string) error {
	if err := database.GetDB().Find(song, "user = ?", user).Error; err != nil {
		return &apperror.CustomError{Message: "No data found"}
	}
	return nil
}

// AddSong(*models.Song, string) adds a song in database and returns error if any
func (sr SongRepo) AddSong(song *models.Song, user string) error {
	song.SetUser(user)
	if err := database.GetDB().Create(song).Error; err != nil {
		return &apperror.CustomError{Message: "Failed to add data"}
	}
	return nil
}

// GetSong(*models.Song, string, string) gets a song from database and returns error if any
func (sr SongRepo) GetSong(song *models.Song, id string, user string) error {
	if err := database.GetDB().Where("id = ? AND user = ?", id, user).First(song).Error; err != nil {
		return &apperror.CustomError{Message: "No data found"}
	}
	return nil
}

// UpdateSong(*models.Song) updates a song in database and returns error if any
func (sr SongRepo) UpdateSong(song *models.Song) error {
	database.GetDB().Save(song)
	return nil
}

// DeleteSong(*models.Song, string, string) deletes a song from database and returns error if any
func (sr SongRepo) DeleteSong(song *models.Song, id string, user string) error {
	resp := database.GetDB().Where("id = ? AND user = ?", id, user).Delete(song)
	if resp.RowsAffected == 0 {
		return &apperror.CustomError{Message: "No record Found"}
	}
	return nil
}
