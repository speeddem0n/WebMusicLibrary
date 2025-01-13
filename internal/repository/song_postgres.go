package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/speeddem0n/WebMusicLibrary/internal/config"
)

type SongPostgres struct {
	db *sqlx.DB
}

func NewSongPostgres(db *sqlx.DB) *SongPostgres {
	return &SongPostgres{db: db}
}

func NewPostgresDB(cfg config.ConfigDB) (*sqlx.DB, error) { // Функция для подключения к БД принимает Config struct
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)) // Считываем настройки для подключения к БД через fmt.Sprintf
	if err != nil {
		return nil, err // Обрабатываем ошибку
	}

	err = db.Ping() // Методом Ping() проверяем работоспособность подключения к БД
	if err != nil {
		return nil, err
	}

	return db, nil // Возвращаем *sqlx.DB
}
