package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Delete song
// @Description Delete song from library using id
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {integer} integer songID
// @Failure 400 {object} errorResponse "Error message"
// @Failure 500 {object} errorResponse "Error message"
// @Router /{id} [delete]
func (h *handlerService) DeleteSongHandler(c *gin.Context) {
	logrus.Info("Received request to delete song.")
	// получаем id из URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("An error occured on getting song id: %d", err))
		return
	}

	// удаляем песню из БД методом Delete
	err = h.dbClient.DeleteSong(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("An error occured on deleting song: %v", err))
		return
	}

	// Ответ пользователю
	logrus.Infof("Song deleted successfully, song_id: %d", id)
	c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully", "id": id})
}
