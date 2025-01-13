package main

import "github.com/sirupsen/logrus"

func main() {
	initLogger() // Инициализируем параметры логера
	logrus.Info("Starting application")

}
