package models

import "time"

type SongModel struct {
	Id          int       `json:"-" db:"id"`
	GroupName   string    `json:"group_name" binding:"required"`
	Song        string    `json:"song" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Text        string    `json:"text" binding:"required"`
	Link        string    `json:"link" binding:"required"`
}
