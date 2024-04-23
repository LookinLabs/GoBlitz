package main

import (
	"log"
	"os"
	"web/middleware"
	db "web/repository/db"

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

	if err := db.NewDBConnection(); err != nil {
		log.Fatalf("failed to establish database connection: %v", err)
	}

	router := gin.New()
	httpRouter := middleware.NewRouter(router)

	if err := httpRouter.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Printf("failed to start the server: %v", err)
	}
}
