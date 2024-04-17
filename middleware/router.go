package middleware

import (
	"net/http"
	"os"
	apiHandler "web/handlers/api"
	errorHandler "web/handlers/error"
	templates "web/templates"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) *gin.Engine {
	router.Use(func(c *gin.Context) {
		host := c.Request.Host
		if host != os.Getenv("APP_HOST") {
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

	router.LoadHTMLGlob("./public/views/*.html")
	router.Use(static.Serve("/", static.LocalFile("./public/", true)))

	router.GET(os.Getenv("API_PATH")+"/ping", apiHandler.StatusOkPingResponse)
	router.GET("/status", templates.StatusPageResponse(), func(c *gin.Context) {
		statuses := c.MustGet("statuses").([]map[string]string)
		c.HTML(http.StatusOK, "status.html", gin.H{
			"services": statuses,
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "public/error/404.html", nil)
	})
	router.Use(errorHandler.StatusBadGateway())
	router.NoRoute(errorHandler.StatusNotFound)

	return router
}
