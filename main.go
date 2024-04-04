package main

import (
	"log"
	"web/config"
	"web/middleware"
	"web/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	env := config.ConfigureEnvironmentals()

	if env.PSQLEnabled == "true" {
		db, err := repository.NewDBConnection()
		if err != nil {
			log.Fatalf("Failed to establish database connection: %v", err)
		}
		defer db.Close()
	}

	router := gin.Default()
	httpRouter := middleware.NewRouter(router, env)

	err := httpRouter.Run(":" + env.AppPort)
	if err != nil {
		log.Printf("Failed to start the server: %v", err)
	}
}
