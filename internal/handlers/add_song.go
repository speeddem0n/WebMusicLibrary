package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
	client "github.com/speeddem0n/WebMusicLibrary/internal/rest_client"
)

type AddSongRequest struct { // Структура для обязательного ввода пользователя
	Group string `json:"group" binding:"required"`
	Song  string `json:"song" binding:"required"`
}

func (h *Handler) AddSong(c *gin.Context) {
	var req AddSongRequest        // Модель для считывания инпута пользователя
	err := c.ShouldBindJSON(&req) // Получаем инпут от пользователя в формате JSON
	if err != nil {
		logrus.Errorf("Failed to fetch required params: %v", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	restClient := client.NewRestClient("http://localhost:8080") // Создаем новый REST клиент

	songDetails, err := restClient.GetSongDetails(req.Group, req.Song) // Методом GetSongDetails делаем запрос на внешний апи и получаем детали песни
	if err != nil {
		logrus.Errorf("Failed to fetch song details: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var releaseDate time.Time

	if songDetails.ReleaseDate != "" { //Проверяем дату на корректность и форматируем ее к нужному формату
		releaseDate, err = time.Parse("2006-01-02", songDetails.ReleaseDate)
		if err != nil {
			logrus.Errorf("Invalid release date format: %v", err)
			newErrorResponse(c, http.StatusBadRequest, "Invalid release date format")
			return
		}
	} else {
		logrus.Warn("Release date missing or invalid")
		newErrorResponse(c, http.StatusBadRequest, "Invalid release date")
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
		logrus.Errorf("Failed to insert song in to DB: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.Infof("Song added successfully, song_id: %d", id)
	c.JSON(http.StatusCreated, gin.H{"message": "Song added successfully", "id": id, "song": song}) // Возвращаем пользователю ответ
}
