package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func CognitoAuth() gin.HandlerFunc {
	keySet, err := jwk.Fetch(context.Background(), "https://cognito-idp."+os.Getenv("AWS_REGION")+".amazonaws.com/"+os.Getenv("AWS_COGNITO_USER_POOL_ID")+"/.well-known/jwks.json")
	if err != nil {
		log.Fatalf("Failed to fetch key set: %v", err)
	}

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseString(tokenString, jwt.WithKeySet(keySet))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if !contains(token.Audience(), os.Getenv("AWS_COGNITO_APP_CLIENT_ID")) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
			return
		}

		c.Next()
	}
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
