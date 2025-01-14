package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) DeleteSongHandler(c *gin.Context) {
	// получаем id из URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("An error occured on getting song id: %d", err))
		return
	}

	// удаляем песню из БД методом Delete
	err = h.songs.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("An error occured on deleting song: %v", err))
		return
	}

	// Ответ пользователю
	logrus.Infof("Song deleted successfully, song_id: %d", id)
	c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully", "id": id})
}
