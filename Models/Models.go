package Models

import (
	"Song_API/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Song struct {
	ID          uint   `json:"id"`
	Song        string `json:"song"`
	Artist      string `json:"artist"`
	Plays       uint   `json:"plays"`
	Playtime    string `json:"playtime"`
	ReleaseDate string `json:"release_date"`
}

func (b *Song) TableName() string {
	return "Songs"
}

func GetAllSong(b *[]Song) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewSong(b *Song) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetSong(b *Song, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSong(b *Song, id string) (err error) {
	if err := Config.DB.Save(b).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSong(b *Song, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
