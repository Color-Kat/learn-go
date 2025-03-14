package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		Db:   DbConfig{Dsn: os.Getenv("DSN")},
		Auth: AuthConfig{SecretToken: os.Getenv("SECRET_TOKEN")},
	}
}

type AuthConfig struct {
	SecretToken string
}
