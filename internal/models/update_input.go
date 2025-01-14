package models

// Модель для инпуты обновления песни
type UpdateInput struct {
	GroupName   string `json:"group" db:"group_name"`
	SongName    string `json:"song" db:"song_name"`
	ReleaseDate string `json:"release_date" db:"release_date"`
	Text        string `json:"text" db:"text"`
	Link        string `json:"link" db:"link"`
}
