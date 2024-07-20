package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"web/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type authMiddlewareTestSuite struct {
	suite.Suite
	router *gin.Engine
}

type authMiddleTestCase struct {
	description    string
	setupRequest   func(req *http.Request, router *gin.Engine)
	mockSetup      func()
	expectedStatus int
}

func (suite *authMiddlewareTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.router = gin.New()
	store := cookie.NewStore([]byte("testing_secret"))
	suite.router.Use(sessions.Sessions("test_session", store))
	suite.router.Use(middleware.Authentication())
	suite.router.GET("/test", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
}

func (suite *authMiddlewareTestSuite) TestAuthenticationMiddleware() {
	tests := []authMiddleTestCase{
		{
			description: "Invalid API Key",
			setupRequest: func(request *http.Request, _ *gin.Engine) {
				request.Header.Set("X-API-Key", "hello-world")
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			description: "Valid API key",
			setupRequest: func(request *http.Request, _ *gin.Engine) {
				request.Header.Set("X-API-Key", os.Getenv("STATUSPAGE_API_KEY"))
			},
			expectedStatus: http.StatusOK,
		},
		{
			description: "Valid session with userID",
			setupRequest: func(_ *http.Request, router *gin.Engine) {
				// Assuming you have a way to directly manipulate the session store or mock it
				// This is a conceptual approach; actual implementation may vary
				router.Use(func(ctx *gin.Context) {
					session := sessions.Default(ctx)
					session.Set("userID", 1) // Set a valid userID
					if err := session.Save(); err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{
							"error": "Failed to save session",
						})
						ctx.Abort()
						return
					}
					ctx.Next()
				})
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, testCase := range tests {
		testCase := testCase // Capture range variable
		suite.Run(testCase.description, func() {
			suite.T().Parallel()
			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			testCase.setupRequest(req, suite.router)
			if testCase.mockSetup != nil {
				testCase.mockSetup()
			}
			testRecorder := httptest.NewRecorder()
			suite.router.ServeHTTP(testRecorder, req)
			assert.Equal(suite.T(), testCase.expectedStatus, testRecorder.Code, testCase.description)
		})
	}
}

func TestAuthMiddlewareTestSuite(tests *testing.T) {
	suite.Run(tests, new(authMiddlewareTestSuite))
}
