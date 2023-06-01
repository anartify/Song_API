package models

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
