package config

import (
	"log"
	"os"
	"web/model"

	"github.com/joho/godotenv"
)

func ConfigureEnvironmentals() model.Config {
	if err := godotenv.Load(); os.Getenv("GO_ENV") == "development" && err != nil {
		log.Fatalf("Error loading .env file")
	}

	urlPrefix := "http://"
	if os.Getenv("ForceSSL") == "true" {
		urlPrefix = "https://"
	}

	return model.Config{
		AppPort:     getEnv("APP_PORT", "8000"),
		AppHost:     getEnv("APP_HOST", "localhost"),
		APIPath:     getEnv("API_PATH", "/api/v1/"),
		ForceSSL:    getEnv("FORCE_SSL", "false"),
		PSQLEnabled: getEnv("PSQL_ENABLED", "false"),
		URLPrefix:   urlPrefix,
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
