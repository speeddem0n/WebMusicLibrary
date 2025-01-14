package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/speeddem0n/WebMusicLibrary/internal/models"
	client "github.com/speeddem0n/WebMusicLibrary/internal/rest_client"
)

// Интерфейс для связи обработчиков со слоем репозитория
type SongRepository interface {
	Add(input models.SongModel) (int, error)
	Delete(id int) error
	GetAll(pag models.PaginationRequest) ([]models.SongModel, error)
	Update(id int, updateData models.SongModel) error
	GetText(songId, page, pageSize int) ([]models.VerseModel, error)
}

// Интерфейс для связи обработчиков с REST клиетом
type RestClient interface {
	GetSongDetails(group, song string) (*client.SongDetail, error)
}

// Структура обработчика
type Handler struct {
	songs  SongRepository // Интерфес для работы со слоем репозитория
	client RestClient     // Интерфейс для работы с Rest клиентом
}

// Конструктор для обработчика
func NewHandler(songs SongRepository, client RestClient) *Handler {
	return &Handler{
		songs:  songs,
		client: client,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	// Инициализация роутера
	router := gin.New()

	songs := router.Group("/songs")
	{
		songs.GET("/", h.GetAllSongsHandler)      // Получить все песни
		songs.POST("/", h.AddSongHandler)         // Добавить песню
		songs.GET("/:id", h.GetSongVerseHandler)  // Получить текст песни
		songs.DELETE("/:id", h.DeleteSongHandler) // Удалить песню
		songs.PUT("/:id", h.UpdateSongHandler)    // Обновить песню

	}

	return router
}
