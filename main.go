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
	if err := godotenv.Load(".env"); os.Getenv("GO_ENV") == "development" && err != nil {
		log.Fatalf("Error loading .env file")
	}

	if os.Getenv("PSQLEnabled") == "true" {
		db, err := repository.NewDBConnection()
		if err != nil {
			log.Fatalf("Failed to establish database connection: %v", err)
		}
		defer db.Close()
	}

	router := gin.Default()
	httpRouter := middleware.NewRouter(router)

	err := httpRouter.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		log.Printf("Failed to start the server: %v", err)
	}
}
