package handlers

import "github.com/gin-gonic/gin"

func Handlers(router *gin.Engine) {
	router.Use(serverErrorHandler())
	router.NoRoute(notFoundHandler)
	router.GET("/status", statusHandler)
}
