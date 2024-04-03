package handlers

import (
	"github.com/gin-gonic/gin"
)

func Handlers(router *gin.Engine) {
	router.Use(ServerErrorHandler())
	router.NoRoute(NotFoundHandler)
}
