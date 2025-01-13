package repository

import (
	"strings"

	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

func (r *SongPostgres) GetText(songId, page, pageSize int) ([]models.VerseModel, error) {
	query := "SELECT text FROM song_lib WHERE id = $1" // SQL запрос
	var fullText string                                // Переменная для записи полного текста песни

	err := r.db.Get(&fullText, query, songId) // Методом Get делаем SQL запрос
	if err != nil {
		return nil, err
	}

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
	lines := strings.Split(text, "\n") // Разбиваем текст на строки

	var verses []models.VerseModel // Массив для хранения куплетов
	var currentVerse string        // Переменная для хранения текущего куплета
	verseNumber := 1               // Счетчик для номера куплета

	for _, line := range lines { // Обходим все строки
		line = strings.TrimSpace(line) // Удаляем лишние пробелы
		if line == "" {                // Если строка пустая то это конец куплета
			if currentVerse != "" { // Если текущий куплет заполнен
				verses = append(verses, models.VerseModel{ // Добавляем куплет в массив
					Number: verseNumber,
					Verse:  currentVerse,
				})
				verseNumber++     // Увеличиваем номер куплета
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
			Number: verseNumber,
			Verse:  currentVerse,
		})
	}

	return verses
}
