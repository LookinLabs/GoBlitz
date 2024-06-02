package helper

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession(ctx *gin.Context, userID uint) error {
	session := sessions.Default(ctx)
	session.Set("userID", userID)

	if err := session.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return err
	}

	return nil
}

func IsUserAuthenticated(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	userID := session.Get("userID")

	return userID != nil
}
