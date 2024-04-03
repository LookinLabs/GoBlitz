package handlers

import (
	"web/src/model"

	"github.com/gin-gonic/gin"
)

func Handlers(router *gin.Engine, env model.Config) {
	router.Use(serverErrorHandler())
	router.NoRoute(notFoundHandler)
	router.GET("/status", statusHandler(env))
}
