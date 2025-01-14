# Запуск
- go run cmd/main.go

# .env Параметры

<strong>Необходимо создать .env файл со следующими параметрами:</strong>

- SERVER_HOST = Адрес сервера (Пример - localhost)
- SERVER_PORT = Порт сервера (Пример - 8000)
- DB_HOST = Адрес postgres (Пример - localhost)
- DB_PORT = Порт postgres (Пример - 5432)
- DB_USERNAME = Имя пользователя postgres (Пример - postgres)
- DB_NAME = Имя базы данных postgres (Пример - music_lib)
- DB_PASSWORD = Пароль пользователя postgres (Пример - postgres)
- DB_SSLMODE = sslmode postgres (Пример - disable)
- API_HOST = Адрес внешнего API на который будет сделан запрос по адресу /info при добавлении песни (Пример - localhost)
- API_PORT = 8080 Порт внешнего API (Пример - 8000)

# Технология реализации приложения
- В качестве web феймворка используется <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- В качестве базы данных используется postgreSQL.
- Работа с БД осуществляется, используя библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Запрос на внешний API для получения деталей песни реализуется с помошью <a href="github.com/go-resty/resty">go-resty</a>.
- Миграции для БД реализуются с помошью <a href="github.com/pressly/goose">goose</a>.
- API описаны с помощью <a href="github.com/swaggo/swag">swagger</a>.