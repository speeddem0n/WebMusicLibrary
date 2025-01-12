package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

func (h *Handler) GetAllSongsHandler(c *gin.Context) {
	page := c.DefaultQuery("page", "1")           // Параметр page, по умолчанию 1
	pageSize := c.DefaultQuery("page_size", "10") // Параметр page_size, по умолчанию 10
	group := c.DefaultQuery("group", "")          // Параметр group
	song := c.DefaultQuery("song", "")            // Параметр song
	fromDate := c.DefaultQuery("from_date", "")   // Параметр from_date
	toDate := c.DefaultQuery("to_date", "")       // Параметр to_date
	text := c.DefaultQuery("text", "")            // Параметр text
	link := c.DefaultQuery("link", "")            // Параметр link

	// Преобразуем page и pageSize в int
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt <= 0 {
		pageSizeInt = 10
	}

	// Парсим даты, если они предоставлены
	var fromDateParsed, toDateParsed time.Time
	if fromDate != "" {
		fromDateParsed, err = time.Parse("2006-01-02", fromDate)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Invalid FromDate format. Expected format is YYYY-MM-DD")
			return
		}
	}
	if toDate != "" {
		toDateParsed, err = time.Parse("2006-01-02", toDate)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Invalid ToDate format. Expected format is YYYY-MM-DD")
			return
		}
	}

	// Строим запрос для получения песен
	req := models.PaginationRequest{
		Page:     pageInt,
		PageSize: pageSizeInt,
		Group:    group,
		Song:     song,
		FromDate: fromDateParsed,
		ToDate:   toDateParsed,
		Text:     text,
		Link:     link,
	}

	// Получаем список песен с фильтрацией и пагинацией
	songs, err := h.songs.GetAll(req)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to  fetch song: %v", err))
		return
	}

	// отправляем успешный ответ с найденными песнями
	c.JSON(http.StatusOK, gin.H{
		"data": songs,
	})
}
