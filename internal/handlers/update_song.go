package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

func (h *Handler) UpdateSongHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // получаем id из URL param
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error on getting song id: %v", err))
		return
	}

	var updateInput models.UpdateInput

	// Получаем данные для обновления песни из тела запроса
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if !ValidateInput(updateInput) {
		newErrorResponse(c, http.StatusBadRequest, "request body is empty")
		return
	}

	// Обновляем песню в базе данных
	err = h.songs.Update(id, updateInput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to update song")
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Song updated successfully"})
}

// Функиця проверяем Не пустой ли инпут для обновления песни
func ValidateInput(input models.UpdateInput) bool {
	if input.GroupName == "" && input.SongName == "" && input.Text == "" && input.ReleaseDate.IsZero() && input.Link == "" {
		return false
	}
	return true
}
