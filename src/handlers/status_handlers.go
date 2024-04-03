package handlers

import (
	"net/http"
	"web/src/config"
	"web/src/model"

	"github.com/gin-gonic/gin"
)

func checkServiceStatus(service model.ServiceInfo) map[string]string {
	_, err := http.Get(service.URL)
	status := map[string]string{
		"name": service.Name,
	}
	if err != nil {
		status["status"] = "down"
	} else {
		status["status"] = "up"
	}
	return status
}

func ReviewServiceStatus(services []model.ServiceInfo) []map[string]string {
	statuses := make([]map[string]string, 0)
	for _, service := range services {
		status := checkServiceStatus(service)
		statuses = append(statuses, status)
	}
	return statuses
}

func statusHandler(c *gin.Context) {
	apiPort, apiPath, _ := config.ConfigureEnvironmentals()
	services := []model.ServiceInfo{
		{Name: "UI", URL: "http://localhost:" + apiPort},
		{Name: "API", URL: "http://localhost:" + apiPort + apiPath},
	}

	statuses := ReviewServiceStatus(services)

	c.HTML(http.StatusOK, "status.html", gin.H{
		"services": statuses,
	})
}
