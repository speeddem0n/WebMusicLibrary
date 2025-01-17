package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Структура для кастомной ошибки в формате json
type errorResponse struct { // Структура для кастомной ошибки в формате json
	Message string `json:"mesage"`
}

// Функция для обработки http ошибок
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)                                              // Выводим сообщение об ошибке в консоль
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message}) // AbortWithStatusJSON принимает hhtp status code и тело ответа. AbortWithStatusJSON блокирует выполнение последующих обработчиков и записывает в ответ статус код и тело сообщения в формате JSON
}
