package config

import (
	"os"
)

type Config struct {
	PORT  string
	DBHOST string
	DBPORT string
	DBUSER string
	DBPASS string
	DBNAME string
	DBSSL string
}

func Load() *Config {
	return &Config {
		PORT: os.Getenv("APP_PORT"),
		DBHOST: os.Getenv("DB_HOST"),
		DBPORT: os.Getenv("DB_PORT"),
		DBUSER: os.Getenv("DB_USER"),
		DBPASS: os.Getenv("DB_PASSWORD"),
		DBNAME: os.Getenv("DB_NAME"),
		DBSSL: os.Getenv("DB_SSLMODE"),
	}
}
