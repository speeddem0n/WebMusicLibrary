package repository

import "github.com/jmoiron/sqlx"

type Song interface{}

type SongPostgres struct {
	db *sqlx.DB
}

type Repository struct {
	Song
}

func NewSongPostgres(db *sqlx.DB) *SongPostgres {
	return &SongPostgres{db: db}
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Song: NewSongPostgres(db),
	}
}
