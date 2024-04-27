package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomePageController(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{})
}
