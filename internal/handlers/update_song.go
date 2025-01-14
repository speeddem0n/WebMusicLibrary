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

	var releaseDate time.Time

	if updateInput.ReleaseDate != "" { //Проверяем дату на корректность и форматируем ее к нужному формату
		releaseDate, err = time.Parse("02.06.2006", updateInput.ReleaseDate)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid release date format: %v", err))
			return
		}
	} else {
		newErrorResponse(c, http.StatusBadRequest, "Release date missing or invalid")
		return
	}

	response := models.SongModel{
		GroupName:   updateInput.GroupName,
		SongName:    updateInput.SongName,
		ReleaseDate: releaseDate,
		Text:        updateInput.Text,
		Link:        updateInput.Link,
	}

	if !ValidateInput(response) {
		newErrorResponse(c, http.StatusBadRequest, "request body is empty")
		return
	}

	// Обновляем песню в базе данных
	err = h.songs.Update(id, response)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to update song")
		return
	}

	logrus.Infof("Song updated successfully, song_id: %d", id)
	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Song updated successfully"})
}

// Функиця проверяем Не пустой ли инпут для обновления песни
func ValidateInput(input models.SongModel) bool {
	if input.GroupName == "" && input.SongName == "" && input.Text == "" && input.ReleaseDate.IsZero() && input.Link == "" {
		return false
	}
	return true
}
