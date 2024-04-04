package main

import (
	"log"
	"web/config"
	"web/middleware"
	"web/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	appEnv, postgresEnv := config.ConfigureEnvironmentals()

	if appEnv.PSQLEnabled == "true" {
		db, err := repository.NewDBConnection(postgresEnv)
		if err != nil {
			log.Fatalf("Failed to establish database connection: %v", err)
		}
		defer db.Close()
	}

	router := gin.Default()
	httpRouter := middleware.NewRouter(router, appEnv)

	err := httpRouter.Run(":" + appEnv.AppPort)
	if err != nil {
		log.Printf("Failed to start the server: %v", err)
	}
}
