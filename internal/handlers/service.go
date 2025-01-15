package handlers

import (
	"github.com/gin-gonic/gin"
	clients "github.com/speeddem0n/WebMusicLibrary/internal/clients"
	"github.com/speeddem0n/WebMusicLibrary/internal/storage"
)

type HandlerService interface {
	UpdateSongHandler(c *gin.Context)
	GetSongVerseHandler(c *gin.Context)
	GetAllSongsHandler(c *gin.Context)
	DeleteSongHandler(c *gin.Context)
	AddSongHandler(c *gin.Context)
}

// Структура клиента обработчика
type handlerService struct {
	dbClient           storage.StorageFacade // Интерфес для работы со слоем репозитория
	externalHttpClient clients.RestClient    // Интерфейс для работы с Rest клиентом
}

// Конструктор для обработчика
func NewHandlerService(dbClient storage.StorageFacade, httpClient clients.RestClient) HandlerService {
	return &handlerService{
		dbClient:           dbClient,
		externalHttpClient: httpClient,
	}
}
