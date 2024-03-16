package configs

import (
	"log"
	"os"
)
import "github.com/joho/godotenv"

type AppConfig struct {
	Port          string
	PostgresqlUrl string
}

var appConfig *AppConfig

func GetAppConfig() *AppConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if appConfig == nil {
		appConfig = &AppConfig{
			Port:          os.Getenv("PORT"),
			PostgresqlUrl: os.Getenv("postgresql"),
		}
	}

	return appConfig
}
