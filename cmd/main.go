package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"

	_ "github.com/speeddem0n/WebMusicLibrary/docs"
	client "github.com/speeddem0n/WebMusicLibrary/internal/clients"
	"github.com/speeddem0n/WebMusicLibrary/internal/connections"
	"github.com/speeddem0n/WebMusicLibrary/internal/handlers"
	"github.com/speeddem0n/WebMusicLibrary/internal/storage"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Функция для настройки логирования
func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

// Функция для запуска миграций
func runMigrations(db *sqlx.DB) {
	migrationsDir := "./scripts/migrations" // Путь к папке с миграциями

	logrus.Info("Running database migrations...")
	if err := goose.Up(db.DB, migrationsDir); err != nil {
		logrus.Fatalf("Failed to apply migrations: %v", err)
	}
	logrus.Info("Database migrations applied successfully.")
}

// Инициализация роутера
func setupRouter(h handlers.HandlerService) *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	songs := router.Group("/songs")
	{
		songs.GET("/list", h.GetAllSongsHandler)  // Получить все песни
		songs.POST("/", h.AddSongHandler)         // Добавить песню
		songs.GET("/:id", h.GetSongVerseHandler)  // Получить текст песни
		songs.DELETE("/:id", h.DeleteSongHandler) // Удалить песню
		songs.PUT("/:id", h.UpdateSongHandler)    // Обновить песню
	}

	return router
}

// @title Music Library API
// @version 1.0
// @description This is a service for managing a music library.

// @host localhost:8000
// @BasePath /
func main() {
	// Загружаем config конфиг приложения
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error on loading env file: %v", err)
	}

	// Инициализируем параметры логера
	initLogger()

	logrus.Info("Starting application")

	// Инициализируем новое подключение к базе данных
	db, err := connections.NewDbConnection()
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	// Запускаем миграции при старте приложения
	runMigrations(db)

	// Инициализация сервиса обрабочиков
	handlerService := handlers.NewHandlerService(
		storage.NewStorageFacade(db),
		client.NewRestClient(),
	)

	// Инициализируем структуру сервера
	srv := http.Server{
		Addr:           "localhost:" + os.Getenv("SERVER_PORT"), // Server address
		Handler:        setupRouter(handlerService),             // Handler
		MaxHeaderBytes: 1 << 20,                                 // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		logrus.Infof("Starting server on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Fatal("Server stopped")
		}
	}()
	logrus.Info("Server is running")

	// Реализация Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("Server shutting down...")

	ctx, ctxCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer ctxCancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		logrus.Errorf("Error occured on srv shutting down: %s", err.Error())
	}

	logrus.Info("Exiting application")
}
