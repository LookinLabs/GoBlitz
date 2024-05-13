package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	helpers "web/helpers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CognitoJWTSession(router *gin.Engine) {
	// Secret Key to encrypt the session
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("environment JWT_SECRET_KEY must be set")
		os.Exit(1)
	}

	store := cookie.NewStore([]byte(secretKey))
	router.Use(sessions.Sessions("authenticatedSession", store))

	router.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("AWS_COGNITO_LOGIN_URL"))
	})

	router.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No code provided"})
			return
		}

		// Fetch the JWT token
		token, err := helpers.JWTInjector()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch JWT token"})
			return
		}

		// Store the JWT token in the session
		session := sessions.Default(c)
		session.Set("jwt", token)
		session.Save()

		// Redirect to home page after successful authentication
		c.Redirect(http.StatusTemporaryRedirect, "/")
	})

	router.Use(func(c *gin.Context) {
		session := sessions.Default(c)
		jwt := session.Get("jwt")
		if jwt != nil {
			c.Request.Header.Add("Authorization", "Bearer "+jwt.(string))
		}
	})
}

func CognitoValidateJWT() gin.HandlerFunc {
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

			// Send the new token in the response body
			c.JSON(http.StatusOK, gin.H{"token": newToken})
			return
		}

		c.Next()
	}
}
