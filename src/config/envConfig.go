package config

import (
	"log"
	"os"
	"web/src/model"

	"github.com/joho/godotenv"
)

func ConfigureEnvironmentals() model.Config {
	err := godotenv.Load()
	if err != nil && os.Getenv("GO_ENV") == "development" {
		log.Fatalf("Error loading .env file")
	}

	return model.Config{
		AppPort:     getEnv("APP_PORT", "8000"),
		AppHost:     getEnv("APP_HOST", "localhost"),
		APIPath:     getEnv("API_PATH", "/api/v1/"),
		ForceSSL:    getEnv("FORCE_SSL", "false"),
		PSQLEnabled: getEnv("PSQL_ENABLED", "false"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
