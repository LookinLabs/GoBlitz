package helper

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, userID uint) error {
	session := sessions.Default(c)
	session.Set("userID", userID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return err
	}
	return nil
}

func IsUserAuthenticated(c *gin.Context) bool {
	session := sessions.Default(c)
	userID := session.Get("userID")
	return userID != nil
}
