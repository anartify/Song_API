package models

// Song struct holds the fields required in the Table songs.
type Song struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Song        string `json:"song"`
	Artist      string `json:"artist"`
	Plays       int    `json:"plays"`
	ReleaseDate string `json:"release_date"`
	User        string `json:"user"`
}

// SetID sets the value of the ID property.
func (s *Song) SetID(id int) {
	s.ID = id
}

// GetID returns the value of the ID property.
func (s *Song) GetID() int {
	return s.ID
}

// SetSong sets the value of the Song property.
func (s *Song) SetSong(song string) {
	s.Song = song
}

// GetSong returns the value of the Song property.
func (s *Song) GetSong() string {
	return s.Song
}

// SetArtist sets the value of the Artist property.
func (s *Song) SetArtist(artist string) {
	s.Artist = artist
}

// GetArtist returns the value of the Artist property.
func (s *Song) GetArtist() string {
	return s.Artist
}

// SetPlays sets the value of the Plays property.
func (s *Song) SetPlays(plays int) {
	s.Plays = plays
}

// GetPlays returns the value of the Plays property.
func (s *Song) GetPlays() int {
	return s.Plays
}

// SetReleaseDate sets the value of the ReleaseDate property.
func (s *Song) SetReleaseDate(releaseDate string) {
	s.ReleaseDate = releaseDate
}

// GetReleaseDate returns the value of the ReleaseDate property.
func (s *Song) GetReleaseDate() string {
	return s.ReleaseDate
}

// SetUser sets the value of the User property.
func (s *Song) SetUser(user string) {
	s.User = user
}

// GetUser returns the value of the User property.
func (s *Song) GetUser() string {
	return s.User
}

// TableName returns the name of the table in the database.
func (s *Song) TableName() string {
	return "songs"
}
