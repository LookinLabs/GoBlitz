package http_handlers

import (
	"github.com/gin-gonic/gin"
)

func notFoundHandler(c *gin.Context) {
	c.Status(404)
	c.File("./public/error/404.html")
}

func Handlers(router *gin.Engine) {
	router.NoRoute(notFoundHandler)
}
