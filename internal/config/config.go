package config

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var (
	Conf config
)

const (
	DbName = "song_lib"
)

type config struct {
	DbHost            string
	DbPort            string
	DbUsername        string
	DbPass            string
	DbName            string
	DbSsl             bool
	ExternalClientUrl string
	LogLvl            logrus.Level
}

func init() {
	Conf = config{}
}

// Инициализация конфига
func Init() {
	flag.StringVar(&Conf.DbHost, "db-host", "localhost", "database host")
	flag.StringVar(&Conf.DbPort, "db-port", "5437", "database port")
	flag.StringVar(&Conf.DbUsername, "db-username", "postgres", "database username")
	flag.StringVar(&Conf.DbPass, "db-pass", "postgres", "database password")
	flag.StringVar(&Conf.DbName, "db-name", "music_lib", "database name")
	flag.BoolVar(&Conf.DbSsl, "db-ssl", false, "enables ssl mode")
	flag.StringVar(&Conf.DbPass, "external-url", "http://localhost:8080", "external api client irl")
	logLevel := flag.String("log-lvl", "debug", "logger message level")

	Conf.LogLvl = parseLogLevel(*logLevel)

	flag.Parse()

}

// Парсинг loglvl параметра
func parseLogLevel(logLevel string) logrus.Level {
	switch logLevel {
	case "panic":
		return logrus.PanicLevel
	case "info":
		return logrus.InfoLevel
	case "trace":
		return logrus.TraceLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.DebugLevel
	}
}
