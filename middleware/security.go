package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SecureConfigMiddleware() gin.HandlerFunc {
	return func(router *gin.Context) {
		hostHeader := router.Request.Host

		// Localhost Host Header is APP_HOST + APP_PORT (e.g. localhost:8000)
		localHostHeader := os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")
		if hostHeader != localHostHeader && hostHeader != os.Getenv("APP_HOST") {
			router.String(http.StatusBadRequest, "Invalid host")
			router.Abort()
			return
		}

		router.Header("X-Frame-Options", "DENY")
		router.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		router.Header("X-XSS-Protection", "1; mode=block")
		router.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		router.Header("Referrer-Policy", "strict-origin")
		router.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
	}
}
