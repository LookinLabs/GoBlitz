package main

import (
	"log"
	"os"
	"web/src/config"
	http "web/src/http"
	persistance "web/src/services"
)

func main() {
	appPort, _, _ := config.ConfigureEnvironmentals()

	if os.Getenv("PSQL_ENABLED") == "true" {
		db, err := persistance.NewDBConnection()
		if err != nil {
			log.Fatalf("Failed to establish database connection: %v", err)
		}
		defer db.Close()
	}

	router := http.NewRouter()
	router.Run(":" + appPort)
}
