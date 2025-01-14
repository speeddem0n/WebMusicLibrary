package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/speeddem0n/WebMusicLibrary/internal/config"
)

type SongPostgres struct {
	db *sqlx.DB
}

// Конуструкто для структуры SongPostgres
func NewSongPostgres(db *sqlx.DB) *SongPostgres {
	return &SongPostgres{db: db}
}

// Функция для подключения к БД принимает
func NewPostgresDB(cfg config.ConfigDB) (*sqlx.DB, error) {
	logrus.Info("Connecting to Data Base...")
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	// Методом Ping() проверяем работоспособность подключения к БД
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	logrus.Info("Connected to Data Base successfully")
	return db, nil // Возвращаем *sqlx.DB
}
