package handlers

import (
	"net/http"
	"testing"
	"web/src/config"
	"web/src/handlers"
	"web/src/model"

	"github.com/go-playground/assert/v2"
	"github.com/jarcoal/httpmock"
)

func TestReviewServiceStatusWhenAPIDown(tests *testing.T) {
	apiPort, apiPath, _ := config.ConfigureEnvironmentals()

	// Test case for when the API is down
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://localhost:"+apiPort+apiPath, httpmock.NewErrorResponder(http.ErrServerClosed))
	services := []model.ServiceInfo{
		{Name: "UI", URL: "http://localhost:" + apiPort},
		{Name: "API", URL: "http://localhost:" + apiPort + apiPath},
	}

	statuses := handlers.ReviewServiceStatus(services)

	if statuses[1]["status"] != "down" {
		tests.Errorf("Expected API status to be down, got %s", statuses[1]["status"])
	}

	assert.Equal(tests, "down", statuses[1]["status"])
}
