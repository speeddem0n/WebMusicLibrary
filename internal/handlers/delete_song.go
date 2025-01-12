package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) DeleteSong(c *gin.Context) { // Метод обработчика для удаления песни
	id, err := strconv.Atoi(c.Param("id")) // получаем id из URL param
	if err != nil {
		logrus.Errorf("Error on getting song id: %v", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.songs.Delete(id) // удаляем песню из БД методом Delete
	if err != nil {
		logrus.Errorf("can't delete song: %s", err)
	}

	logrus.Infof("Song deleted successfully, song_id: %d", id)
	c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully", "id": id}) // Ответ пользователю
}
