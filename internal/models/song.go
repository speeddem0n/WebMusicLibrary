package models

import "time"

type SongModel struct {
	Id          int       `json:"-" db:"id"`
	GroupName   string    `json:"group_name" binding:"required"`
	Song        string    `json:"song_name" binding:"required"`
	ReleaseDate time.Time `json:"release_date"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}
