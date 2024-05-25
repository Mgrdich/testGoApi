package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port          string
	PostgresqlUrl string
	Environment   string
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
			PostgresqlUrl: os.Getenv("POSTGRESQL"),
			Environment:   os.Getenv("ENVIRONMENT"),
		}
	}

	return appConfig
}
