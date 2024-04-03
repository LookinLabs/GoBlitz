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

	appPort := os.Getenv("APP_PORT")
	appHost := os.Getenv("APP_HOST")
	apiPath := os.Getenv("API_PATH")
	forceSSL := os.Getenv("FORCE_SSL")
	psqlEnabled := os.Getenv("PSQL_ENABLED")

	switch {
	case forceSSL == "":
		forceSSL = "false"
	case psqlEnabled == "":
		psqlEnabled = "false"
	}

	return model.Config{
		AppPort:     appPort,
		AppHost:     appHost,
		APIPath:     apiPath,
		ForceSSL:    forceSSL,
		PSQLEnabled: psqlEnabled,
	}
}
