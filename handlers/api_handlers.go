package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainAPIHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}
