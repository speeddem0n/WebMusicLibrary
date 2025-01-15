package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

// Структура для обязательного ввода пользователя
type AddSongRequest struct {
	Group string `json:"group" binding:"required"`
	Song  string `json:"song" binding:"required"`
}

// @Summary Add a new song
// @Description Add a new song to the music library
// @Tags Songs
// @Accept json
// @Produce json
// @Param input body AddSongRequest true "Song details"
// @Success 201 {integer} integer songID
// @Failure 400 {object} errorResponse "Error message"
// @Failure 500 {object} errorResponse "Error message"
// @Router /songs [post]
func (h *handlerService) AddSongHandler(c *gin.Context) {
	logrus.Infof("Recived request to add song")
	var req AddSongRequest

	// Получаем инпут от пользователя в формате JSON
	err := c.ShouldBindJSON(&req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Failed to fetch required params: %v", err))
		return
	}

	// Методом GetSongDetails делаем запрос на внешний API и получаем детали песни
	logrus.Infof("Fetching song details for group: %s, song: %s", req.Group, req.Song)
	songDetails, err := h.externalHttpClient.GetSongDetails(req.Group, req.Song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Invalid release date format: %v. Expected format is DD.MM.YYYY", err))
		return
	}
	var releaseDate time.Time

	// Валидируем дату и форматируем ее к нужному формату
	if songDetails.ReleaseDate != "" {
		releaseDate, err = time.Parse("02.01.2006", songDetails.ReleaseDate)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid release date format: %v. Expected format is DD.MM.YYYY", err))
			return
		}
	} else {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Release date missing or invalid: %v", err))
		return
	}

	// Создаем модель песни
	song := models.SongModel{
		GroupName:   req.Group,
		SongName:    req.Song,
		ReleaseDate: releaseDate,
		Text:        songDetails.Text,
		Link:        songDetails.Link,
	}

	// Методом Add из слоя репозитория добавляем новую песню в БД
	id, err := h.dbClient.AddSong(song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to insert song in to DB: %v", err))
		return
	}

	// Возвращаем пользователю ответ
	logrus.Infof("Song added successfully, song_id: %d", id)
	c.JSON(http.StatusCreated, gin.H{"message": "Song added successfully", "id": id})
}
