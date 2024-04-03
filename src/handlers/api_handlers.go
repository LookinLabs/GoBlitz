package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootAPIPath(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}

func GetUserIDHandler(c *gin.Context) {
	c.JSON(http.StatusOK, []gin.H{
		{
			"id":       "1",
			"username": "John Doe",
		},
		{
			"id":       "2",
			"username": "Jane Doe",
		},
	})
}
