package models

import "time"

type UpdateInput struct {
	GroupName   string    `json:"group"`
	SongName    string    `json:"song"`
	ReleaseDate time.Time `json:"release_date"`
	Text        string    `json:"text" db:"text"`
	Link        string    `json:"link" db:"link"`
}
