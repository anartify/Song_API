package Models

import (
	"Song_API/Config"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	_ "github.com/go-sql-driver/mysql"
)

type Song struct {
	ID          uint   `json:"id"`
	Song        string `json:"song"`
	Artist      string `json:"artist"`
	Plays       uint   `json:"plays"`
	ReleaseDate string `json:"release_date"`
}

func (b *Song) TableName() string {
	return "Songs"
}

func (data Song) Validation() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.Song, validation.Required),
		validation.Field(&data.Artist, validation.Required),
		validation.Field(&data.Plays, validation.Required),
		validation.Field(&data.ReleaseDate, validation.Required, validation.Date(time.DateOnly)),
	)
}

type SongInterface interface {
	GetAllSong(*[]Song) error
	AddNewSong(*Song) error
	GetSong(*Song, string) error
	UpdateSong(*Song, string) error
	DeleteSong(*Song, string) error
}

type SongModel struct{}

func (m SongModel) GetAllSong(b *[]Song) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func (m SongModel) AddNewSong(b *Song) (err error) {
	if err = b.Validation(); err != nil {
		return err
	}
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func (m SongModel) GetSong(b *Song, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func (m SongModel) UpdateSong(b *Song, id string) (err error) {
	if err = b.Validation(); err != nil {
		return err
	}
	Config.DB.Save(b)
	return nil
}

func (m SongModel) DeleteSong(b *Song, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}

// func () GetAllSong(b *[]Song) (err error) {
// 	if err = Config.DB.Find(b).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func AddNewSong(b *Song) (err error) {
// 	if err = b.Validation(); err != nil {
// 		return err
// 	}
// 	if err = Config.DB.Create(b).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetSong(b *Song, id string) (err error) {
// 	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func UpdateSong(b *Song, id string) (err error) {
// 	if err = b.Validation(); err != nil {
// 		return err
// 	}
// 	if err := Config.DB.Save(b).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func DeleteSong(b *Song, id string) (err error) {
// 	Config.DB.Where("id = ?", id).Delete(b)
// 	return nil
// }
