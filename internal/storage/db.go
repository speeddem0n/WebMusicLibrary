package storage

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

// Метод БД для добавления новой песни
func (r *storage) AddSong(song models.SongModel) (int, error) {
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

// Метод для удаления песни из БД
func (r *storage) DeleteSong(id int) error {
	// SQL запрос
	query := "DELETE FROM song_lib WHERE id = $1"

	// Выполняем запрос
	logrus.Debugf("Executing query: %s with args: %+v", query, id)
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
	logrus.Infof("Query executed successfully, rows affected: %d", rowsAffected)
	return nil
}

// Метод БД для получения фсех песен с фильтрацией и пагинацией
func (r *storage) GetAllSongs(pag models.PaginationRequest) ([]models.SongModel, error) {
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
	debugMsg := fmt.Sprintf("Executing query: %s", query)
	logrus.Debugf(debugMsg+" with args: %+v", args...)
	err := r.db.Select(&Songs, query, args...)
	if err != nil {
		return nil, err
	}

	// Возвращаем полученный слайс с песнями
	logrus.Infof("Query executed successfully, %d songs selected on page %d", pag.PageSize, pag.Page)
	return Songs, nil
}

// Метод БД для получения текста песни с пагинацией для куплетов
func (r *storage) GetSongText(songId, page, pageSize int) ([]models.VerseModel, error) {
	// SQL запрос
	query := "SELECT text FROM song_lib WHERE id = $1"
	// Переменная для записи полного текста песни
	var fullText string

	// Методом Get делаем SQL запрос
	logrus.Debugf("Executing query: %s with args: %+v", query, songId)
	err := r.db.Get(&fullText, query, songId)
	if err != nil {
		return nil, fmt.Errorf("song with id %d doesn't exists", songId)
	}
	logrus.Infof("Query executed successfully, 1 row added with id: %d", songId)
	// Разбиваем текст на куплеты
	verses := splitTextToVerses(fullText)

	// Реализуем пагинацию
	start := (page - 1) * pageSize // Начальный индекс
	if start >= len(verses) {      // Если индекс за пределами массива куплетов то страница пустая
		return []models.VerseModel{}, nil
	}

	end := start + pageSize // Конечный индекс
	if end > len(verses) {  // Если конечный индекс больше длинны массива, обрезаем его
		end = len(verses)
	}

	return verses[start:end], nil
}

// Функция splitTextToVerses разбивает текст на куплеты
func splitTextToVerses(text string) []models.VerseModel {
	// Разбиваем текст на строки
	lines := strings.Split(text, "\n")

	// Массив для хранения куплетов
	var verses []models.VerseModel

	// Переменная для хранения текущего куплета
	var currentVerse string

	// Обходим все строки
	for _, line := range lines {
		// Удаляем лишние пробелы
		line = strings.TrimSpace(line)

		// Если строка пустая то это конец куплета
		if line == "" {
			// Если текущий куплет заполнен
			if currentVerse != "" {
				verses = append(verses, models.VerseModel{ // Добавляем куплет в массив
					Verse: currentVerse,
				})
				currentVerse = "" // Сбрасываем текущий куплет

			}
		} else { // В Если строка не пустая
			if currentVerse != "" { // Если текущий куплет не пустой добавляем в конце строки символ \n
				currentVerse += "\n"
			}
			currentVerse += line
		}
	}

	// Добавляем последний куплет если он остался
	if currentVerse != "" {
		verses = append(verses, models.VerseModel{
			Verse: currentVerse,
		})
	}

	return verses
}

// Метод БД для обнавления песни
func (r *storage) UpdateSongs(id int, updateData models.SongModel) error {
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
