package middleware

import (
	"net/http"
	"os"
	"web/controller/api"
	errorController "web/controller/error"
	httpTemplates "web/views/templates"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) *gin.Engine {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Security Headers
	router.Use(func(c *gin.Context) {
		hostHeader := c.Request.Host

		// Localhost Host Header is APP_HOST + APP_PORT (e.g. localhost:8000)
		localHostHeader := os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")
		if hostHeader != localHostHeader && hostHeader != os.Getenv("APP_HOST") {
			c.String(http.StatusBadRequest, "Invalid host")
			c.Abort()
			return
		}

		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
	})

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

	// API Handling
	apiGroup := router.Group(os.Getenv("API_PATH"))
	{
		apiGroup.GET("/ping", api.StatusOkPingResponse)
		apiGroup.GET("/users", api.GetUsers)
	}

	// HTML Templates (e.g Status page)
	router.LoadHTMLGlob("./views/**/*")
	router.GET("/status", httpTemplates.StatusPageResponse(), StatusPageMiddleware())

	// Error handling
	router.NoRoute(errorController.StatusNotFound)

	return router
}
