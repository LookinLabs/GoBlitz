package main

import (
	"log"
	"os"
	"web/middleware"
	"web/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "production" && os.Getenv("DEBUG_MODE") != "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	psql := os.Getenv("PSQLEnabled")
	if psql == "true" {
		db, err := repository.NewDBConnection()
		if err != nil {
			log.Fatalf("Failed to establish database connection: %v", err)
		}
		defer db.Close()
	}

	router := gin.New()
	httpRouter := middleware.NewRouter(router)

	err := httpRouter.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		log.Printf("Failed to start the server: %v", err)
	}
}
