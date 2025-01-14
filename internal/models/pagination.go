package models

import "time"

// Модель для запроса на получение всех песен
type PaginationRequest struct {
	Page     int       `json:"page"`                // Номер страницы
	PageSize int       `json:"page_size"`           // Размер страницы
	Group    string    `json:"group,omitempty"`     // Фильтрация по названию группы
	Song     string    `json:"song,omitempty"`      // Фильтрация по названию песни
	FromDate time.Time `json:"from_date,omitempty"` // Фильтрация по дате (ПОСЛЕ)
	ToDate   time.Time `json:"to_date,omitempty"`   // Фильтрация по дате (ДО)
	Text     string    `json:"text,omitempty"`      // Фильтрация по тексту песни
	Link     string    `json:"link,omitempty"`      // // Фильтрация по ссылке
}
