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

func (h *Handler) AddSongHandler(c *gin.Context) {
	var req AddSongRequest

	// Получаем инпут от пользователя в формате JSON
	err := c.ShouldBindJSON(&req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Failed to fetch required params: %v", err))
		return
	}

	// Методом GetSongDetails делаем запрос на внешний API и получаем детали песни
	songDetails, err := h.client.GetSongDetails(req.Group, req.Song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to fetch song details: %v", err))
		return
	}
	var releaseDate time.Time

	// Валидируем дату и форматируем ее к нужному формату
	if songDetails.ReleaseDate != "" {
		releaseDate, err = time.Parse("02.06.2006", songDetails.ReleaseDate)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid release date format: %v", err))
			return
		}
	} else {
		newErrorResponse(c, http.StatusBadRequest, "Release date missing or invalid")
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
	id, err := h.songs.Add(song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to insert song in to DB: %v", err))
		return
	}

	// Возвращаем пользователю ответ
	logrus.Infof("Song added successfully, song_id: %d", id)
	c.JSON(http.StatusCreated, gin.H{"message": "Song added successfully", "id": id})
}
