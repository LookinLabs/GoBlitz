package tests

import (
	"net/http"
	"os"
	"testing"

	model "web/model"
	httpTemplates "web/views/templates"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func WrapTestsIntoSuite(t *testing.T) {
	suite.Run(t, new(TestsSuite))
}

func (suite *TestsSuite) TestStatusPageWhenAllServicesDown() {
	// Arrange
	urlPrefix := "http://"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	services := []model.StatusPage{
		{
			Name: "API",
			URL:  urlPrefix + os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT") + os.Getenv("API_PATH") + "ping",
		},
	}

	// Act

	// Register mock responders for all service endpoints
	for _, service := range services {
		httpmock.RegisterResponder("GET", service.URL, httpmock.NewErrorResponder(http.ErrServerClosed))
	}

	statuses := httpTemplates.CheckServicesStatus(services)

	// Assert
	// Check the status of all services
	for i, service := range services {
		assert.Equalf(suite.T(), "down", statuses[i]["status"], "expected %s status to be down, got %s", service.Name, statuses[i]["status"])
	}
}
