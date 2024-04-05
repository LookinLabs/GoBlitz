package middleware

import (
	"net/http"
	"web/config"
	apiHandler "web/handlers/api"
	errorHandler "web/handlers/error"
	envModel "web/model/config"
	templates "web/templates"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, env envModel.AppEnv) *gin.Engine {
	// Setup Security Headers and check for valid host
	router.Use(func(c *gin.Context) {
		config.MiddlewareHTTPConfig(c, env)
	})

	// Load HTML Templates
	router.LoadHTMLGlob("./public/views/*.html")

	// Serve static files
	router.Use(static.Serve("/", static.LocalFile("./public/", true)))

	// API Routes
	router.GET(env.APIPath+"/ping", apiHandler.StatusOkPingResponse)
	router.GET("/status", templates.StatusPageResponse(env), func(c *gin.Context) {
		statuses := c.MustGet("statuses").([]map[string]string)
		c.HTML(http.StatusOK, "status.html", gin.H{
			"services": statuses,
		})
	})

	// Error Handlers
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "public/error/404.html", nil)
	})
	router.Use(errorHandler.StatusBadGateway())
	router.NoRoute(errorHandler.StatusNotFound)

	return router
}
