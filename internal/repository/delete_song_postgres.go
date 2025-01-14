package repository

import "fmt"

// Метод для удаления песни из БД
func (r *SongPostgres) Delete(id int) error {
	// SQL запрос
	query := "DELETE FROM song_lib WHERE id = $1"

	// Выполняем запрос
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
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
	return nil
}
