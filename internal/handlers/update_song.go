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

// @Summary Update song
// @Description Update song from library using id
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param input body models.UpdateInput true "Input to update song (At least one field must be filled in)"
// @Success 200 {integer} integer songID
// @Failure 400 {object} errorResponse "Error message"
// @Failure 500 {object} errorResponse "Error message"
// @Router /songs/{id} [put]
func (h *handlerService) UpdateSongHandler(c *gin.Context) {
	logrus.Info("Received request to update song.")
	// получаем id песни
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("An error occured on getting song id: %v", err))
		return
	}

	var updateInput models.UpdateInput

	// Получаем данные для обновления песни из тела запроса
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Проверяем дату на корректность и форматируем ее к нужному формату
	var releaseDate time.Time

	if updateInput.ReleaseDate != "" {
		releaseDate, err = time.Parse("02.01.2006", updateInput.ReleaseDate)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid release date format: %v. Expected format is DD.MM.YYYY", err))
			return
		}
	}

	// Строим запрос к бд для обновления песни
	response := models.SongModel{
		GroupName:   updateInput.GroupName,
		SongName:    updateInput.SongName,
		ReleaseDate: releaseDate,
		Text:        updateInput.Text,
		Link:        updateInput.Link,
	}

	// Валидируем инпут
	if !ValidateInput(response) {
		newErrorResponse(c, http.StatusBadRequest, "Request body is empty")
		return
	}

	// Обновляем песню в базе данных
	err = h.dbClient.UpdateSong(id, response)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to update song: %v", err))
		return
	}

	// Возвращаем успешный ответ
	logrus.Infof("Song updated successfully, song_id: %d", id)
	c.JSON(http.StatusOK, gin.H{"message": "Song updated successfully", "id": id})
}

// Функиця проверяем Не пустой ли инпут для обновления песни
func ValidateInput(input models.SongModel) bool {
	if input.GroupName == "" && input.SongName == "" && input.Text == "" && input.ReleaseDate.IsZero() && input.Link == "" {
		return false
	}
	return true
}
