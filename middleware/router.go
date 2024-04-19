package middleware

import (
	"net/http"
	"os"
	apiHandler "web/controller/api"
	errorHandler "web/controller/error"
	templates "web/views/view_templates"

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

	// Static File Handling
	router.Use(static.Serve("/", static.LocalFile("./public/", true)))
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	// API Handling
	router.GET(os.Getenv("API_PATH")+"/ping", apiHandler.StatusOkPingResponse)

	// HTML Templates (e.g Status page)
	router.LoadHTMLGlob("./views/*.html")
	router.Use(static.Serve("/templates/assets", static.LocalFile("./views/assets/", true)))
	router.GET("/status", templates.StatusPageResponse(), func(c *gin.Context) {
		statuses := c.MustGet("statuses").([]map[string]string)
		c.HTML(http.StatusOK, "status.html", gin.H{
			"services": statuses,
		})
	})

	// Error Handling
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "public/error/404.html", nil)
	})

	router.Use(errorHandler.StatusBadGateway())
	router.NoRoute(errorHandler.StatusNotFound)

	return router
}
