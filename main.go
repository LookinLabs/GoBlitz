package main

import (
	"log"
	"os"
	"web/middleware"
	repository "web/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading .env file")
	}

	if os.Getenv("DEBUG_MODE") != "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	database, err := repository.NewDBConnection()
	if err != nil {
		log.Fatalf("failed to establish database connection: %v", err)
	}

	if database != nil && database.DB != nil {
		defer database.DB.Close()
	}

	router := gin.New()
	httpRouter := middleware.NewRouter(router)

	if err := httpRouter.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Printf("failed to start the server: %v", err)
	}
}
