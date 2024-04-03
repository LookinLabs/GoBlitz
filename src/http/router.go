package http

import (
	"net/http"
	"web/src/config"
	handlers "web/src/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	appPort, apiPath, appDomain := config.ConfigureEnvironmentals()
	router.LoadHTMLGlob("./views/*.html")

	// Setup Security Headers
	router.Use(func(c *gin.Context) {
		if c.Request.Host != "localhost:"+appPort && c.Request.Host != appDomain {
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
		c.Next()
	})

	// Setup route group for the API
	api := router.Group(apiPath)
	{
		apiHandler := func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to the API!",
			})
		}
		api.GET(apiPath, apiHandler)
	}

	// Serve frontend static files
	router.NoRoute(func(c *gin.Context) {
		c.File("./public" + c.Request.URL.Path)
	})

	handlers.Handlers(router)
	return router
}
