package main

import (
	"log"
	"os"
	"web/middleware"
	sql "web/repository/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
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

	var err error
	sql.DB, err = sql.NewDBConnection()
	if err != nil {
		log.Fatalf("failed to establish database connection: %v", err)
	}

	router := gin.New()
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))
	httpRouter := middleware.NewRouter(router)

	if err := httpRouter.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Printf("failed to start the server: %v", err)
	}
}
