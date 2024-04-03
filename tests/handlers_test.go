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

func TestReviewServiceStatusWhenAllServicesDown(tests *testing.T) {
	env := config.ConfigureEnvironmentals()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	urlPrefix := "http://"

	services := []model.ServiceInfo{
		{Name: "API", URL: urlPrefix + env.AppHost + ":" + env.AppPort + env.APIPath + "ping"},
		{Name: "Users", URL: urlPrefix + env.AppHost + ":" + env.AppPort + env.APIPath + "users"},
	}

	// Register mock responders for all service endpoints
	for _, service := range services {
		httpmock.RegisterResponder("GET", service.URL, httpmock.NewErrorResponder(http.ErrServerClosed))
	}

	statuses := handlers.ReviewServiceStatus(services)

	// Check the status of all services
	for i, service := range services {
		if statuses[i]["status"] != "down" {
			tests.Errorf("Expected %s status to be down, got %s", service.Name, statuses[i]["status"])
		}

		assert.Equal(tests, "down", statuses[i]["status"])
	}
}
