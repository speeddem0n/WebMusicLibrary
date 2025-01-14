package repository

import (
	"github.com/sirupsen/logrus"
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
	logrus.Debugf("Executing query: %s with args: %+v, %+v, %+v, %+v, %+v", query, song.GroupName, song.SongName, song.ReleaseDate, song.Text, song.Link)
	row := r.db.QueryRow(query, song.GroupName, song.SongName, song.ReleaseDate, song.Text, song.Link)

	// Получаем id только что созданной песни
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	logrus.Infof("Query executed successfully, 1 row added with id: %d", id)
	// Вовращаем id успешно созданной песни
	return id, nil
}
