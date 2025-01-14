package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

type songResonse struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// @Summary Get all songs
// @Description Retrieve all songs with pagination and optional filters
// @Tags Songs
// @Accept json
// @Produce json
// @Param group query string false "Filter by group"
// @Param song query string false "Filter by song name"
// @Param after query string false "Filtering after Date DD.MM.YYYY"
// @Param before query string false "Filtering before Date DD.MM.YYYY"
// @Param text query string false "Filter by text"
// @Param link query string false "Filter by link"
// @Param page query int false "Page number (Default: 1)"
// @Param page_size query int false "Page size (Default: 10)"
// @Success 200 {array} songResonse "All songs"
// @Failure 400 {object} errorResponse "Error message"
// @Failure 500 {object} errorResponse "Error message"
// @Router / [get]
func (h *Handler) GetAllSongsHandler(c *gin.Context) {
	group := c.DefaultQuery("group", "")   // Параметр group
	song := c.DefaultQuery("song", "")     // Параметр song
	after := c.DefaultQuery("after", "")   // Параметр after
	before := c.DefaultQuery("before", "") // Параметр before
	text := c.DefaultQuery("text", "")     // Параметр text
	link := c.DefaultQuery("link", "")     // Параметр link

	// Преобразуем page и pageSize в int
	page, err := strconv.Atoi(c.DefaultQuery("page", "1")) // Параметр page, по умолчанию 1
	if err != nil || page <= 0 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10")) // Параметр page_size, по умолчанию 10
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	// Парсим даты, если они предоставлены
	var afterParsed, beforeParsed time.Time
	if after != "" {
		afterParsed, err = time.Parse("01.02.2006", after)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Invalid FromDate format. Expected format is DD.MM.YYYY")
			return
		}
	}
	if before != "" {
		beforeParsed, err = time.Parse("01.02.2006", before)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Invalid ToDate format. Expected format is DD.MM.YYYY")
			return
		}
	}

	// Строим запрос для получения песен
	req := models.PaginationRequest{
		Page:     page,
		PageSize: pageSize,
		Group:    group,
		Song:     song,
		FromDate: afterParsed,
		ToDate:   beforeParsed,
		Text:     text,
		Link:     link,
	}

	// Получаем список песен с фильтрацией и пагинацией
	songs, err := h.songs.GetAll(req)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to  fetch song: %v", err))
		return
	}
	var response []songResonse

	for _, song := range songs {
		response = append(response, songResonse{
			Group:       song.GroupName,
			Song:        song.SongName,
			ReleaseDate: song.ReleaseDate.Format("02.06.2006"),
			Text:        song.Text,
			Link:        song.Link,
		})

	}

	// Отправляем успешный ответ с найденными песнями
	logrus.Infof("Getting songs,  page: %d", page)
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}
