package models

type VerseModel struct {
	Id     int    `json:"-" db:"id"`
	SongId int    `json:"-" db:"song_id"`
	Verse  string `json:"verse"`
}
