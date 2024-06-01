package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusPageMiddleware handles the status page
func StatusPageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		statuses := c.MustGet("statuses").([]map[string]string)
		c.HTML(http.StatusOK, "status.html", gin.H{
			"services": statuses,
		})
	}
}

// WelcomePageMiddleware handles the welcome page
func WelcomePageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{})
	}
}

func LoginPageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}
}

func DocumentationPageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "docs.html", gin.H{})
	}
}
