package config

import (
	"log"
	"os"
	"web/model"

	"github.com/joho/godotenv"
)

func ConfigureEnvironmentals() (model.AppConfig, model.PostgresConfig) {
	if err := godotenv.Load(".env"); os.Getenv("GO_ENV") == "development" && err != nil {
		log.Fatalf("Error loading .env file")
	}

	urlPrefix := "http://"
	if os.Getenv("ForceSSL") == "true" {
		urlPrefix = "https://"
	}

	appConfig := model.AppConfig{
		AppPort:     getEnv("APP_PORT", "8000"),
		AppHost:     getEnv("APP_HOST", "localhost"),
		APIPath:     getEnv("API_PATH", "/api/v1/"),
		ForceSSL:    getEnv("FORCE_SSL", "false"),
		PSQLEnabled: getEnv("PSQL_ENABLED", "false"),
		URLPrefix:   urlPrefix,
	}

	postgresConfig := model.PostgresConfig{
		DBHost:     getEnv("POSTGRES_HOST", "localhost"),
		DBPort:     getEnv("POSTGRES_PORT", "5432"),
		DBUser:     getEnv("POSTGRES_USER", "postgres"),
		DBPassword: getEnv("POSTGRES_PASSWORD", "password"),
		DBDatabase: getEnv("POSTGRES_DB", "postgres"),
	}

	return appConfig, postgresConfig
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
