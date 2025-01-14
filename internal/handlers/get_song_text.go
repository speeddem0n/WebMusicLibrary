package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetSongVerseHandler(c *gin.Context) {
	// Получаем id песни
	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid song id")
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
		newErrorResponse(c, http.StatusInternalServerError, "Failed to fetch song text")
		return
	}

	// Возвращаем результат пользователю
	logrus.Infof("Getting song text,  page: %d", page)
	c.JSON(http.StatusOK, gin.H{"text": verses})
}
