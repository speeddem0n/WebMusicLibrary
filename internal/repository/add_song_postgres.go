package repository

import (
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

func (r *SongPostgres) Add(song models.SongModel) (int, error) {
	var id int // переменная для записи id только что созданной песни
	query := `
    INSERT INTO song_lib (group_name, song_name, release_date, text, link)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id` // Запрос в БД
	row := r.db.QueryRow(query, song.GroupName, song.SongName, song.ReleaseDate, song.Text, song.Link) // Исполняем sql Запрос
	err := row.Scan(&id)                                                                               // Получаем id только что созданной песни
	if err != nil {
		return 0, err
	}

	return id, nil
}
