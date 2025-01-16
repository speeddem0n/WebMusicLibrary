# 🎶 Онлайн библиотека песен
Тестовое задание для effective mobile

## Описание проекта
Проект представляет собой REST API для управления онлайн библиотекой песен. Он предоставляет функционал для работы с песнями, включая получение списка песен с фильтрацией и пагинацией, просмотр текста песни с разделением на куплеты, а также операции добавления, изменения и удаления песен.
В качестве web феймворка используется <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.

## Установка и запуск

### Шаги установки
1. Склонируйте репозиторий:
```bash
git clone https://github.com/speeddem0n/WebMusicLibrary && cd WebMusicLibrary
```
2. Настройте .env файл:
   Создайте файл .env в корневой директории проекта и укажите следующие параметры:

Настройки приложения определяются в файле .env. Ниже перечислены все переменные окружения с их описанием:

| Переменная       | Описание                                          | Пример значения           |
|------------------|--------------------------------------------------|---------------------------|
| SERVER_PORT    | Порт сервера                                    | 8000                    |
| DB_HOST        | Адрес хоста базы данных                         | localhost               |
| DB_PORT        | Порт для подключения к базе данных              | 5432                    |
| DB_USERNAME    | Имя пользователя для подключения к базе данных  | postgres                |
| DB_NAME        | Имя базы данных                                 | music_library           |
| DB_PASSWORD    | Пароль пользователя базы данных                 | password                |
| DB_SSLMODE     | SSL mode                                        | disable                 |
| API_URL        | URL внешнего API для получения информации о песнях | http://localhost:8080 |

### Пример .env файла

```dotenv
SERVER_PORT = 8000
DB_HOST = localhost
DB_PORT = 5438
DB_USERNAME = postgres
DB_NAME = songlib
DB_PASSWORD = postgres
DB_SSLMODE = disable
API_URL = http://localhost:8080
```

3. Установите зависимости:

```bash
go mod tidy
```

4. Запустите приложение

```bash
go run cmd/main.go
```

### Основные функции
1. REST API Методы:
   - Получение данных библиотеки с фильтрацией и пагинацией:
     - Фильтрация по следующим полям: group, song, releaseDate, text, link.
     - Поддержка постраничного вывода.
   - Получение текста песни с пагинацией по куплетам:
     - Разделение текста песни на куплеты с использованием пустых строк в качестве разделителей.
   - Добавление новой песни
    Обогащение данных о песне через внешний API (описанный в Swagger) и сохранение в базе данных.
   - Изменение данных песни:
     - Обновление отдельных полей или всех полей сразу.
   - Удаление песни:
     - Удаление записи из библиотеки.

2. Интеграция с внешним API:
   - При добавлении новой песни отправляется запрос в API, описанный Swagger:
     - Путь: /info
     - Параметры запроса: group, song.
     - Ответ: Обогащенная информация о песне (дата выхода, текст, ссылка).

3. База данных:
   - Используется PostgreSQL.
   - Структура базы данных создается с помощью миграций, выполняемых при старте сервиса (с использованием <a href="https://github.com/pressly/goose">goose</a>).

4. Логирование:
   - Поддержка debug- и info-логов с использованием библиотеки <a href="https://github.com/sirupsen/logrus">logrus</a>.

5. Конфигурация:
   - Все настройки сервиса хранятся в .env файле.

6. Swagger документация:
   - Сгенерирован Swagger для всех методов API.
   - Доступна через /swagger/index.html после запуска сервера.

---

## Примеры API запросов

### Получение списка песен
GET songs/list
```http
GET http://localhost:8000/songs/list?page=1&pageSize=4&group=Muse
```

### Получение текста песни
GET songs/:id
```http
GET http://localhost:8080/songs/1?page=1&pageSize=5
```

### Добавление новой песни
POST songs/
```http
POST http://localhost:8000/songs
Content-Type: application/json

{
  "group": "Muse",
  "song": "Supermassive Black Hole"
}
```

### Изменение песни
PUT songs/:id
```http
PUT http://localhost:8000/songs/1
Content-Type: application/json

{
  "text": "Updated song text"
}
```

## Удаление песни
DELETE songs/:id
```http
DELETE http://localhost:8000/songs/1
```

---

## Swagger документация
Документация API доступна после запуска сервера по адресу:
```http
http://localhost:8080/swagger/index.html
```

