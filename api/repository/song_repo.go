package repository

import (
	"Song_API/api/models"
	"Song_API/database"
)

type SongInterface interface {
	GetAllSong(b *[]models.Song) error
	AddSong(b *models.Song) error
	GetSong(b *models.Song, id string) error
	UpdateSong(b *models.Song) error
	DeleteSong(b *models.Song, id string) error
}
type SongRepo struct{}
type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func (sr SongRepo) GetAllSong(b *[]models.Song) error {
	if err := database.DB.Find(b).Error; err != nil {
		return &CustomError{message: "No data found"}
	}
	return nil
}
func (sr SongRepo) AddSong(b *models.Song) error {
	if err := b.Validation(); err != nil {
		return &CustomError{message: err.Error()}
	}
	if err := database.DB.Create(b).Error; err != nil {
		return &CustomError{message: "Failed to add data"}
	}
	return nil
}

func (sr SongRepo) GetSong(b *models.Song, id string) error {
	if err := database.DB.Where("id = ?", id).First(b).Error; err != nil {
		return &CustomError{message: "No data found"}
	}
	return nil
}

func (sr SongRepo) UpdateSong(b *models.Song) error {
	if err := b.Validation(); err != nil {
		return &CustomError{message: err.Error()}
	}
	database.DB.Save(b)
	return nil
}

func (sr SongRepo) DeleteSong(b *models.Song, id string) error {
	database.DB.Where("id = ?", id).Delete(b)
	return nil
}
