package middleware

import (
	"net/http"

	helper "web/helpers"

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

func LoginPageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if helper.IsUserAuthenticated(c) {
			// If user is authenticated, serve the logout page
			c.HTML(http.StatusOK, "authenticated.html", gin.H{})
		} else {
			// If user is not authenticated, serve the authenticate page
			c.HTML(http.StatusOK, "login.html", gin.H{})
		}
	}
}
