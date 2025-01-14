package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"github.com/speeddem0n/WebMusicLibrary/internal/config"
	"github.com/speeddem0n/WebMusicLibrary/internal/handlers"
	"github.com/speeddem0n/WebMusicLibrary/internal/repository"
	client "github.com/speeddem0n/WebMusicLibrary/internal/rest_client"
	"github.com/speeddem0n/WebMusicLibrary/internal/server"
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
	migrationsDir := "./internal/repository/migrations" // Путь к папке с миграциями

	logrus.Info("Running database migrations...")
	if err := goose.Up(db.DB, migrationsDir); err != nil {
		logrus.Fatalf("Failed to apply migrations: %v", err)
	}
	logrus.Info("Database migrations applied successfully.")
}

// @title Music Library API
// @version 1.0
// @description This is a service for managing a music library.

// @host localhost:8000
// @BasePath /
func main() {
	// Инициализируем параметры логера
	initLogger()
	logrus.Info("Starting application")

	// Загружаем .env файл с параметрами приложения
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loadong env variables: %s", err.Error())
	}
	logrus.Info("Env variables loaded successfully")

	// Инициализируем новое подключение к базе данных и передаем в него параметры из .env
	db, err := repository.NewPostgresDB(config.ConfigDB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}

	// Запускаем миграции при старте приложения
	runMigrations(db)

	// Инициализация зависимостей
	repo := repository.NewSongPostgres(db)                                          // Слой репозитория
	restClient := client.NewRestClient(os.Getenv("API_URL"), os.Getenv("API_PORT")) // Создаем новый REST клиент для запроса на внешний API
	handler := handlers.NewHandler(repo, restClient)                                // Обработчики

	// Инициализируем структуру сервера
	srv := new(server.Server)

	// Запускаем сервер в отдельной горутине
	logrus.Info("Starting server...")
	go func() {
		err := srv.Run(os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"), handler.InitRoutes())
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Failed to run server: %s", err.Error())
		}
	}()
	logrus.Info("Server is running")

	// Реализация Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("Server shutting down...")

	err = srv.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("Error occured on server shutting down: %s", err.Error())
	}

	logrus.Info("Server gracefully shuted down")

	// Закрываем подключение к бд
	logrus.Info("Closing connection to Data Base...")
	err = db.Close()
	if err != nil {
		logrus.Errorf("Error occured on db connection close: %s", err.Error())
	}

	logrus.Info("Connection to DB is closed")
	logrus.Info("Exiting application")
}
