package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
)

type AddSongRequest struct { // Структура для обязательного ввода пользователя
	Group string `json:"group" binding:"required"`
	Song  string `json:"song" binding:"required"`
}

func (h *Handler) AddSongHandler(c *gin.Context) {
	var req AddSongRequest        // Модель для считывания инпута пользователя
	err := c.ShouldBindJSON(&req) // Получаем инпут от пользователя в формате JSON
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Failed to fetch required params: %v", err))
		return
	}

	songDetails, err := h.client.GetSongDetails(req.Group, req.Song) // Методом GetSongDetails делаем запрос на внешний апи и получаем детали песни
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to fetch song details: %v", err))
		return
	}
	var releaseDate time.Time

	if songDetails.ReleaseDate != "" { //Проверяем дату на корректность и форматируем ее к нужному формату
		releaseDate, err = time.Parse("02.06.2006", songDetails.ReleaseDate)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid release date format: %v", err))
			return
		}
	} else {
		newErrorResponse(c, http.StatusBadRequest, "Release date missing or invalid")
		return
	}

	song := models.SongModel{ // Записываем весь полученый инпут в структуру SongModel
		GroupName:   req.Group,
		SongName:    req.Song,
		ReleaseDate: releaseDate,
		Text:        songDetails.Text,
		Link:        songDetails.Link,
	}

	id, err := h.songs.Add(song) // Методом Add из слоя репозитория добавляем новую песню в БД
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to insert song in to DB: %v", err))
		return
	}

	logrus.Infof("Song added successfully, song_id: %d", id)
	c.JSON(http.StatusCreated, gin.H{"message": "Song added successfully", "id": id}) // Возвращаем пользователю ответ
}
