package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusPageController(c *gin.Context) {
	statuses := c.MustGet("statuses").([]map[string]string)
	c.HTML(http.StatusOK, "status.html", gin.H{
		"services": statuses,
	})
}
