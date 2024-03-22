package main

import (
	"os"
	"web/middleware"
)

func main() {
	router := middleware.NewRouter()
	router.Run(":" + os.Getenv("APP_PORT"))
}
