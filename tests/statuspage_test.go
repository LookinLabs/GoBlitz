package handlers

import (
	"net/http"
	"testing"
	"web/config"
	templatesModel "web/model/templates"
	templates "web/templates"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestStatusPageWhenAllServicesDown(tests *testing.T) {
	env, _ := config.LoadEnvironmentals()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	services := []templatesModel.ServiceStatusInfo{
		{
			Name: "API",
			URL:  env.URLPrefix + env.AppHost + ":" + env.AppPort + env.APIPath + "ping",
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
