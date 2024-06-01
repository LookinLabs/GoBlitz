package middleware

import (
	"net/http"
	"os"
	"web/controller/api"
	errorController "web/controller/error"
	helper "web/helpers"
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

	// Frontend & Static Files Handling
	router.LoadHTMLGlob("./views/**/*.html")

	router.Use(static.Serve("/assets", static.LocalFile("./public/assets", true)))
	if helper.CheckFileNotExists("./public/index.html") {
		router.GET("/", WelcomePageMiddleware())
		router.GET("/favicon.ico", func(c *gin.Context) {
			c.String(http.StatusNoContent, "")
		})
	} else {
		router.Use(static.Serve("/", static.LocalFile("./public/", true)))
	}

	helper.ServePageAssets(router)

	// HTML Template generated pages
	router.GET("/status", httpTemplates.StatusPageResponse(), StatusPageMiddleware())
	router.GET("/login", LoginPageMiddleware())

	// Authentication
	router.POST("/signup", api.Signup)
	router.POST("/signin", api.Signin)
	router.GET("/signout", api.Signout)

	// API Handling
	apiGroup := router.Group(os.Getenv("API_PATH"), Authentication())
	{
		apiGroup.GET("/ping", api.StatusOkPingResponse)
		apiGroup.GET("/users", api.GetUsers)
	}

	// Error handling
	router.NoRoute(errorController.StatusNotFound)

	return router
}
