package middleware

import (
	"net/http"
	"os"

	sql "web/repository/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionID := session.Get("userID")

		if sessionID == nil {
			// Check for API key in the request
			apiKey := ctx.GetHeader("X-API-Key")
			if apiKey != os.Getenv("STATUSPAGE_API_KEY") {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

		} else {
			userID := sessionID.(uint)

			user, err := sql.GetUserByID(userID)
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			ctx.Set("userID", user.ID)
		}
		ctx.Next()
	}
}
