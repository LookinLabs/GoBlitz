package main

import (
	"log"
	"os"
	"web/src/config"
	http "web/src/http"
	persistence "web/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	env := config.ConfigureEnvironmentals()

	if os.Getenv("PSQL_ENABLED") == "true" {
		db, err := persistence.NewDBConnection()
		if err != nil {
			log.Fatalf("Failed to establish database connection: %v", err)
		}
		defer db.Close()
	}

	router := gin.Default()
	httpRouter := http.NewRouter(router, env)

	err := httpRouter.Run(":" + env.AppPort)
	if err != nil {
		log.Printf("Failed to start the server: %v", err)
	}
}
