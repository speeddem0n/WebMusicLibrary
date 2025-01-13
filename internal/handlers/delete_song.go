package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) DeleteSongHandler(c *gin.Context) { // Метод обработчика для удаления песни
	id, err := strconv.Atoi(c.Param("id")) // получаем id из URL param
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error on getting song id: %v", err))
		return
	}

	err = h.songs.Delete(id) // удаляем песню из БД методом Delete
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("can't delete song: %s", err))
		return
	}

	logrus.Infof("Song deleted successfully, song_id: %d", id)
	c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully", "id": id}) // Ответ пользователю
}
