package http

import (
	"net/http"
	handlers "web/src/handlers"
	model "web/src/model"
	routes "web/src/routes"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, env model.Config) *gin.Engine {
	// Setup Security Headers and check for valid host
	router.Use(func(c *gin.Context) {
		if env.AppHost == "localhost" {
			c.Request.Host = "localhost"
		}

		if c.Request.Host != env.AppHost {
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

	routes.APIRoutes(router, env)
	routes.StaticRoutes(router)

	// Setup error handlers
	handlers.Handlers(router)

	return router
}
