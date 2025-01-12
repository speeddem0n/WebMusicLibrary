package models

import "time"

type PaginationRequest struct {
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
	Group    string    `json:"group,omitempty"`
	Song     string    `json:"song,omitempty"`
	FromDate time.Time `json:"from_date,omitempty"`
	ToDate   time.Time `json:"to_date,omitempty"`
	Text     string    `json:"text,omitempty"`
	Link     string    `json:"link,omitempty"`
}
