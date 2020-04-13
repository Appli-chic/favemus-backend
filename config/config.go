package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseDialect      string
	DatabaseHost         string
	DatabasePort         string
	DatabaseUser         string
	DatabaseName         string
	DatabasePassword     string
	DatabaseSSlActivated string
}

var Conf Config

func LoadConfiguration() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	Conf = Config{
		DatabaseDialect:      os.Getenv("DATABASE_DIALECT"),
		DatabaseHost:         os.Getenv("DATABASE_HOST"),
		DatabasePort:         os.Getenv("DATABASE_PORT"),
		DatabaseUser:         os.Getenv("DATABASE_USER"),
		DatabaseName:         os.Getenv("DATABASE_NAME"),
		DatabasePassword:     os.Getenv("DATABASE_PASSWORD"),
		DatabaseSSlActivated: os.Getenv("DATABASE_SSL_ACTIVATED"),
	}
}
