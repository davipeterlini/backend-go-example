// pkg/config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadDatabaseConfig() DBConfig {
	err := godotenv.Load() // Carrega as variáveis de ambiente do arquivo .env
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
