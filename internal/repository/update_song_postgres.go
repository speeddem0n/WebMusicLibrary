package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

// Метод БД для обнавления песни
func (r *SongPostgres) Update(id int, updateData models.SongModel) error {
	// Строим базовый SQL запрос
	query := "UPDATE song_lib SET "
	args := []interface{}{}
	paramIndex := 1

	// Динамически добавляем только те поля, которые были переданы
	if updateData.GroupName != "" {
		query += fmt.Sprintf("group_name = $%d, ", paramIndex)
		args = append(args, updateData.GroupName)
		paramIndex++
	}
	if updateData.SongName != "" {
		query += fmt.Sprintf("song_name = $%d, ", paramIndex)
		args = append(args, updateData.SongName)
		paramIndex++
	}
	if !updateData.ReleaseDate.IsZero() {
		query += fmt.Sprintf("release_date = $%d, ", paramIndex)
		args = append(args, updateData.ReleaseDate)
		paramIndex++
	}
	if updateData.Text != "" {
		query += fmt.Sprintf("text = $%d, ", paramIndex)
		args = append(args, updateData.Text)
		paramIndex++
	}
	if updateData.Link != "" {
		query += fmt.Sprintf("link = $%d, ", paramIndex)
		args = append(args, updateData.Link)
		paramIndex++
	}

	// Убираем последнюю запятую и пробел
	query = query[:len(query)-2]

	// Добавляем условие для идентификации записи
	query += fmt.Sprintf(" WHERE id = $%d", paramIndex)
	args = append(args, id)

	// Выполняем запрос
	logrus.Debugf("Executing query: %s with args: %+v", query, args)
	result, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update song with ID %d: %v", id, err)
	}
	// Из результата запроса получаем количество затронутых строк в БД
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Если количество строк = 0 то такой песни не существует
	if rowsAffected == 0 {
		return fmt.Errorf("song with id %d doesn't exists", id)
	}
	logrus.Infof("Query executed successfully, rows affected: %d", rowsAffected)

	return nil
}
