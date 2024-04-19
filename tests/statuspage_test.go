package handlers

import (
	"net/http"
	"os"
	"testing"
	model "web/model"
	templates "web/views/view_templates"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestStatusPageWhenAllServicesDown(tests *testing.T) {
	urlPrefix := os.Getenv("FORCE_TLS")
	if urlPrefix == "" {
		urlPrefix = "http://"
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	services := []model.ServiceStatusInfo{
		{
			Name: "API",
			URL:  urlPrefix + os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT") + os.Getenv("API_PATH") + "ping",
		},
	}

	// Register mock responders for all service endpoints
	for _, service := range services {
		httpmock.RegisterResponder("GET", service.URL, httpmock.NewErrorResponder(http.ErrServerClosed))
	}

	statuses := templates.CheckServicesStatus(services)

	// Check the status of all services
	for i, service := range services {
		assert.Equalf(tests, "down", statuses[i]["status"], "Expected %s status to be down, got %s", service.Name, statuses[i]["status"])
	}
}
