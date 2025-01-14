package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Get song text
// @Description Get song text from library with pagination by verses using song id
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param page query int false "Page number (Default: 1)"
// @Param page_size query int false "Page size (Default: 10)"
// @Success 200 {array} models.VerseModel "Verses"
// @Failure 400 {object} errorResponse "Error message"
// @Failure 500 {object} errorResponse "Error message"
// @Router /{id} [get]
func (h *Handler) GetSongVerseHandler(c *gin.Context) {
	logrus.Info("Received request to fetch song text.")
	// Получаем id песни
	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid song id: %v", err))
		return
	}

	// Поулчаем параметры пагинации
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Получаем текст песни из бд
	verses, err := h.songs.GetText(songID, page, pageSize)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to fetch song text: %v", err))
		return
	}

	// Возвращаем результат пользователю
	logrus.Infof("Getting song text successfully, got %d verses,  page: %d", len(verses), page)
	c.JSON(http.StatusOK, gin.H{"text": verses})
}
