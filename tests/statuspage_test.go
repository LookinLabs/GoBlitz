package handlers

import (
	"log"
	"net/http"
	"os"
	"testing"
	model "web/model"
	httpTemplates "web/views/templates"

	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestStatusPageWhenAllServicesDown(tests *testing.T) {
	if err := godotenv.Load("../.env.tests"); err != nil {
		log.Fatalf("Error loading .env.tests file")
	}

	urlPrefix := "http://"
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

	statuses := httpTemplates.CheckServicesStatus(services)

	// Check the status of all services
	for i, service := range services {
		assert.Equalf(tests, "down", statuses[i]["status"], "Expected %s status to be down, got %s", service.Name, statuses[i]["status"])
	}
}
