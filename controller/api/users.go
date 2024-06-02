package api

import (
	"net/http"

	sql "web/repository/db"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := sql.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
