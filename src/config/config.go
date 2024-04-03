package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConfigureEnvironmentals() (string, string, string) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort := os.Getenv("APP_PORT")
	appDomain := os.Getenv("APP_DOMAIN")
	apiPath := os.Getenv("API_PATH")

	return appPort, apiPath, appDomain

}
