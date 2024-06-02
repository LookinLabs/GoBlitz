package middleware

import (
	"net/http"
	"os"
	sql "web/repository/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		session := sessions.Default(c)
		sessionID := session.Get("userID")

		if sessionID == nil {
			// Check for API key in the request
			apiKey := c.GetHeader("X-API-Key")
			if apiKey != os.Getenv("STATUSPAGE_API_KEY") {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else {
			userID := sessionID.(uint)
			user, err := sql.GetUserByID(userID)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Set("userID", user.ID)
		}

		c.Next()
	}
}
