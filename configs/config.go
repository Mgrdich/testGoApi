package configs

import (
	"log"
	"os"
	"path"
	"regexp"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port          string
	PostgresqlUrl string
	Environment   string
}

var appConfig *AppConfig

const projectDirName = "testGoApi"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, err := os.Getwd()

	if err != nil {
		log.Fatalf("Failed to get current working directory error: %v", err)
	}

	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err = godotenv.Load(path.Join(string(rootPath), "/.env"))

	if err != nil {
		log.Fatalf("Error loading .env file error: %v", err)
	}
}

func GetAppConfig() *AppConfig {
	if appConfig == nil {
		loadEnv()
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
func SetAppConfig(config *AppConfig) {
	appConfig = config
}
