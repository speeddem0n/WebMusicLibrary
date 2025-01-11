package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/speeddem0n/WebMusicLibrary/internal/repository"
)

type Handler struct {
	repository *repository.Repository
}

func NewHandler(repository *repository.Repository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	songs := router.Group("/songs")
	{
		songs.GET("/", h.GetAllSongs)
		songs.POST("/", h.AddSong)
		songs.GET("/:id", h.GetSongVerses)
		songs.DELETE("/:id", h.DeleteSong)
		songs.PUT("/:id", h.UpdateSong)

	}
}
