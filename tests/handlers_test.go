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
	env := config.ConfigureEnvironmentals()
	// Test case for when the API is down
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	urlPrefix := "http://"

	httpmock.RegisterResponder("GET", urlPrefix+env.AppHost+":"+env.AppPort+env.APIPath, httpmock.NewErrorResponder(http.ErrServerClosed))

	services := []model.ServiceInfo{
		{Name: "UI", URL: urlPrefix + env.AppHost + ":" + env.AppPort},
		{Name: "API", URL: urlPrefix + env.AppHost + ":" + env.AppPort + env.APIPath},
	}

	statuses := handlers.ReviewServiceStatus(services)

	if statuses[1]["status"] != "down" {
		tests.Errorf("Expected API status to be down, got %s", statuses[1]["status"])
	}

	assert.Equal(tests, "down", statuses[1]["status"])
}
