package middleware

import (
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func StaticPagesMiddleware(router *gin.Engine) {
	if _, err := os.Stat("./public/index.html"); os.IsNotExist(err) {
		// Load welcome page from html template
		router.GET("/", WelcomePageMiddleware())
	} else {
		// Handle static files from the public folder
		router.Use(static.Serve("/", static.LocalFile("./public/", true)))
	}

	// Serve static assets
	router.Use(static.Serve("/assets", static.LocalFile("./public/assets", true)))
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})
}
