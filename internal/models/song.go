package models

import "time"

// Модель для песни
type SongModel struct {
	Id          int       `json:"-" db:"id"`
	GroupName   string    `json:"group" binding:"required" db:"group_name"`
	SongName    string    `json:"song" binding:"required" db:"song_name"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	Text        string    `json:"text" db:"text"`
	Link        string    `json:"link" db:"link"`
}
