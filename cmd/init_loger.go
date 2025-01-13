package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func initLogger() {
	// Настройка логирования
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}
