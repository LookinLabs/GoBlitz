package http_handlers

import (
	"github.com/gin-gonic/gin"
)

func notFoundHandler(c *gin.Context) {
	c.Status(404)
	c.File("./public/error/404.html")
}

func serverErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if request := recover(); request != nil {
				c.Status(502)
				c.File("./public/error/502.html")
			}
		}()
		c.Next()
	}
}

func Handlers(router *gin.Engine) {
	router.Use(serverErrorHandler())
	router.NoRoute(notFoundHandler)
}
