package middleware

import (
	"os"
	"web/controller/api"
	errorController "web/controller/error"
	httpTemplates "web/views/templates"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) *gin.Engine {

	if os.Getenv("PUBLIC_API_ACCESS") != "true" {
		CognitoJWTSession(router)
		router.Use(CognitoValidateJWT())
	}

	StaticPagesMiddleware(router)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(SecureConfigMiddleware())

	// API Handling
	// If PUBLIC_API_ACCESS is not set, then the API is without JWT
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
