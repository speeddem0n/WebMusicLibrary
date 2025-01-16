package connections

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func NewDbConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE")))
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
