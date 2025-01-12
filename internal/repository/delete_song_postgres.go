package repository

import "fmt"

func (r *SongPostgres) Delete(id int) error { // Метод для удаления песни из БД
	query := "DELETE FROM song_lib WHERE id = $1" // sql запрос

	result, err := r.db.Exec(query, id) // Исполняем запрос
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected() // из результата запроса получаем количество затронутых строк в БД
	if err != nil {
		return err
	}

	if rowsAffected == 0 { // если количество строк = 0 то такой песни не существует
		return fmt.Errorf("no songs found with ID: %d", id)
	}
	return nil
}
