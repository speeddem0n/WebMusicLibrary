package config

// Структура для параметров подключения к БД
type ConfigDB struct {
	Host     string
	Port     string
	Username string
	DBName   string
	SSLMode  string
	Password string
}

// Структура для параметров подключения к серверу
type ConfigServer struct {
	Host string
	Port string
}
