package main

import (
	"log"
	"os"
	http "web/src/http"
	persistance "web/src/services"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("PSQL_ENABLED") == "true" {
		db, err := persistance.NewDBConnection()
		if err != nil {
			log.Fatalf("Failed to establish database connection: %v", err)
		}
		defer db.Close()
	}

	router := http.NewRouter()
	router.Run(":" + os.Getenv("APP_PORT"))
}
