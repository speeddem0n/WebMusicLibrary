package repository

import (
	"fmt"

	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

// Метод БД для получения фсех песен с фильтрацией и пагинацией
func (r *SongPostgres) GetAll(pag models.PaginationRequest) ([]models.SongModel, error) {
	// Составной SQL запрос
	query := "SELECT group_name, song_name, release_date, text, link FROM song_lib WHERE 1=1 "
	args := []interface{}{} // Слайс для хранения аргументов подставляемых вместо плейсхолдеров
	paramIndex := 1         // Счетчик номера параметров для плейсхолдера

	// Реализация фильтрации
	if pag.Group != "" {
		query += fmt.Sprintf("AND group_name ILIKE $%d", paramIndex)
		args = append(args, "%"+pag.Group+"%")
		paramIndex++
	}

	if pag.Song != "" {
		query += fmt.Sprintf("AND song_name ILIKE $%d", paramIndex)
		args = append(args, "%"+pag.Song+"%")
		paramIndex++
	}

	if pag.Text != "" {
		query += fmt.Sprintf("AND text ILIKE $%d", paramIndex)
		args = append(args, "%"+pag.Text+"%")
		paramIndex++
	}

	if pag.Link != "" {
		query += fmt.Sprintf("AND link ILIKE $%d", paramIndex)
		args = append(args, "%"+pag.Link+"%")
		paramIndex++
	}

	if !pag.FromDate.IsZero() {
		query += fmt.Sprintf("AND release_date >= $%d", paramIndex)
		args = append(args, pag.FromDate)
		paramIndex++
	}

	if !pag.ToDate.IsZero() {
		query += fmt.Sprintf("AND release_date <= $%d", paramIndex)
		args = append(args, pag.ToDate)
		paramIndex++
	}

	// Реализация пагинации
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1)
	args = append(args, pag.PageSize, (pag.Page-1)*pag.PageSize)

	// Переменная для хранения всех полученных песен
	var Songs []models.SongModel

	// Выполняем sql запрос
	err := r.db.Select(&Songs, query, args...)
	if err != nil {
		return nil, err
	}

	// Возвращаем полученный слайс с песнями
	return Songs, nil
}
