package repository

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

// Метод БД для получения текста песни с пагинацией для куплетов
func (r *SongPostgres) GetText(songId, page, pageSize int) ([]models.VerseModel, error) {
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
