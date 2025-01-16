package connections

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/speeddem0n/WebMusicLibrary/internal/config"
)

func NewDbConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Conf.DbHost, config.Conf.DbPort, config.Conf.DbUsername, config.Conf.DbPass, "music_lib", configureSsl()))
	if err != nil {
		return nil, err
	}

	// Методом Ping() проверяем работоспособность подключения к БД
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil // Возвращаем *sqlx.DB
}

func configureSsl() string {
	if config.Conf.DbSsl {
		return "enable"
	}

	return "disable"
}
