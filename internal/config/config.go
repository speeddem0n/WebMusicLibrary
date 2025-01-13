package config

type ConfigDB struct {
	Host     string
	Port     string
	Username string
	DBName   string
	SSLMode  string
	Password string
}

type ConfigServer struct {
	Host string
	Port string
}

func initConfig() {

}
