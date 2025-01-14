package repository

import (
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

// Метод БД для добавления новой песни
func (r *SongPostgres) Add(song models.SongModel) (int, error) {
	// Переменная для записи id только что созданной песни
	var id int

	// SQL запрос
	query := `
    INSERT INTO song_lib (group_name, song_name, release_date, text, link)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id`

	// Исполняем sql Запрос
	row := r.db.QueryRow(query, song.GroupName, song.SongName, song.ReleaseDate, song.Text, song.Link)

	// Получаем id только что созданной песни
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	// Вовращаем id успешно созданной песни
	return id, nil
}
