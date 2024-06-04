package configs

import (
	"log"
	"os"
	"path"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
	"testGoApi/internal/util"
)

type AppConfig struct {
	Port                   string
	PostgresqlUrl          string
	Environment            util.EnvironmentsType
	JwtSecretKey           []byte
	TokenExpirationMinutes int
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

	err = godotenv.Load(path.Join(string(rootPath), ".env"))

	if err != nil {
		log.Fatalf("Error loading .env file error: %v", err)
	}
}

// GetAppConfig is the Singleton getter for the AppConfig
// This should be used in the application wherever we need to get the application config
func GetAppConfig() *AppConfig {
	if appConfig == nil {
		loadEnv()

		TokenExpirationMinutes, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))

		if err != nil {
			log.Fatalf("Token Expire should be of type integer error: %v", err)
		}

		osEnv := os.Getenv("ENVIRONMENT")
		env, ok := util.LookUpEnv(osEnv)

		if !ok {
			log.Fatalf("can not parse this environment value %v", osEnv)
		}

		appConfig = &AppConfig{
			Port:                   os.Getenv("PORT"),
			PostgresqlUrl:          os.Getenv("POSTGRESQL"),
			Environment:            env,
			JwtSecretKey:           []byte(os.Getenv("JWT_SECRET_KEY")),
			TokenExpirationMinutes: TokenExpirationMinutes}
	}

	return appConfig
}

// SetAppConfig should only be in test
// AppConfig config should be read from the .env files
func SetAppConfig(config *AppConfig) {
	appConfig = config
}
