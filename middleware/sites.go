package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusPageMiddleware handles the status page
func StatusPageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.MustGet("statuses").([]map[string]string)
		c.HTML(http.StatusOK, "status.html", gin.H{
			"services": status,
		})
	}
}

// WelcomePageMiddleware handles the welcome page
func WelcomePageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{})
	}
}

func ConsolePageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "console.html", gin.H{})
	}
}
