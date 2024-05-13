package helpers

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"web/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTInjector() gin.HandlerFunc {
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
		token, err := FetchJWTToken(awsCognitoAPIUserEmail, awsCognitoAPIPassword)
		if err != nil {
			fmt.Println("Error fetching JWT token:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Add JWT Token to request
		c.Request.Header.Add("Authorization", "Bearer "+token)

		// Add JWT Token to response
		c.Header("Authorization", "JWT")
		c.Next()
	}
}

func FetchJWTToken(username, password string) (string, error) {
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		log.Fatal("environment AWS_REGION must be set")
		os.Exit(1)
	}

	awsCognitoAppClientId := os.Getenv("AWS_COGNITO_APP_CLIENT_ID")
	if awsCognitoAppClientId == "" {
		log.Fatal("environment AWS_COGNITO_APP_CLIENT_ID must be set")
		os.Exit(1)
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(awsRegion))
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := cognitoidentityprovider.NewFromConfig(cfg)

	params := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},
		ClientId: aws.String(awsCognitoAppClientId),
	}

	if resp, err := client.InitiateAuth(context.Background(), params); err != nil {
		return "", fmt.Errorf("failed to authenticate: %v", err)
	} else {
		return *resp.AuthenticationResult.IdToken, nil
	}
}

func FetchRefreshToken(refreshToken string) (string, error) {
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		log.Fatal("environment AWS_REGION must be set")
		os.Exit(1)
	}

	awsCognitoAppClientId := os.Getenv("AWS_COGNITO_APP_CLIENT_ID")
	if awsCognitoAppClientId == "" {
		log.Fatal("environment AWS_COGNITO_APP_CLIENT_ID must be set")
		os.Exit(1)
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(awsRegion))
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := cognitoidentityprovider.NewFromConfig(cfg)

	params := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeRefreshTokenAuth,
		AuthParameters: map[string]string{
			"REFRESH_TOKEN": refreshToken,
		},
		ClientId: aws.String(awsCognitoAppClientId),
	}

	resp, err := client.InitiateAuth(context.Background(), params)
	if err != nil {
		return "", fmt.Errorf("failed to fetch the refresh token: %v", err)
	}

	return *resp.AuthenticationResult.IdToken, nil
}

// parseToken is a function that takes a JWT token and returns the parsed RSA public key and an error.
func ParseRSAKeys(token *jwt.Token) (interface{}, error) {
	// Validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		log.Fatal("environment AWS_REGION must be set")
		os.Exit(1)
	}

	awsCognitoUserPoolId := os.Getenv("AWS_COGNITO_USER_POOL_ID")
	if awsCognitoUserPoolId == "" {
		log.Fatal("environment AWS_COGNITO_USER_POOL_ID must be set")
		os.Exit(1)
	}

	// Fetch the JWKS from AWS Cognito
	resp, err := http.Get(fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", awsRegion, awsCognitoUserPoolId))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jwks = &model.JWKS{}
	if err := json.NewDecoder(resp.Body).Decode(jwks); err != nil {
		return nil, err
	}

	// Find the key with the matching kid (where kid is the key identifier)
	for _, key := range jwks.Keys {
		if key.Kid == token.Header["kid"] {
			// Decode the N and E fields
			nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
			if err != nil {
				return nil, err
			}
			eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
			if err != nil {
				return nil, err
			}

			// Construct the RSA public key
			n := new(big.Int).SetBytes(nBytes)
			e := new(big.Int).SetBytes(eBytes).Int64()
			parsedKey := &rsa.PublicKey{N: n, E: int(e)}

			return parsedKey, nil
		}
	}

	return nil, fmt.Errorf("unable to find appropriate key")
}
