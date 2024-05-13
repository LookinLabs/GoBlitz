package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	helpers "web/helpers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CognitoAuth() gin.HandlerFunc {
	awsCognitoAPIUserEmail := os.Getenv("AWS_COGNITO_API_USER_EMAIL")
	if awsCognitoAPIUserEmail == "" {
		log.Fatal("environment AWS_COGNITO_API_USER_EMAIL must be set")
		os.Exit(1)
	}

	awsCognitoAPIPassword := os.Getenv("AWS_COGNITO_API_PASSWORD")
	if awsCognitoAPIPassword == "" {
		log.Fatal("environment AWS_COGNITO_API_PASSWORD must be set")
		os.Exit(1)
	}

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "401 Unauthorized Error. Authorization header must be provided and formatted correctly."})
			return
		}

		// Remove the "Bearer " prefix from the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, helpers.ParseRSAKeys)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "401 Unauthorized Error. Token Not Found."})
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok || time.Now().Unix() >= int64(exp) {
			refreshToken := c.GetHeader("Refresh-Token")
			var newToken string
			var err error

			if refreshToken == "" {
				newToken, err = helpers.FetchJWTToken(awsCognitoAPIUserEmail, awsCognitoAPIPassword)
			} else {
				newToken, err = helpers.FetchRefreshToken(refreshToken)
			}

			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.Header("Authorization", "Bearer "+newToken)
			return
		}

		c.Next()
	}
}
