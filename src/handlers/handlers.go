package handlers

import (
	"web/src/model"

	"github.com/gin-gonic/gin"
)

func Handlers(router *gin.Engine, env model.Config) {
	router.Use(ServerErrorHandler())
	router.NoRoute(NotFoundHandler)
}
