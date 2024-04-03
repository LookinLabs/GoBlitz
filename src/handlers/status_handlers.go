package handlers

import (
	"net/http"
	"web/src/model"

	"github.com/gin-gonic/gin"
)

func checkServiceStatus(service model.ServiceInfo) map[string]string {
	resp, err := http.Get(service.URL)
	status := map[string]string{
		"name": service.Name,
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	status["status"] = "up"

	if err != nil {
		status["status"] = "down"
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

func statusHandler(env model.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPrefix := "http://"

		if env.ForceSSL == "true" {
			urlPrefix = "https://"
		}

		services := []model.ServiceInfo{
			{Name: "UI", URL: urlPrefix + env.AppHost + ":" + env.AppPort},
			{Name: "API", URL: urlPrefix + env.AppHost + ":" + env.AppPort + env.APIPath},
		}

		statuses := ReviewServiceStatus(services)

		c.HTML(http.StatusOK, "status.html", gin.H{
			"services": statuses,
		})
	}
}
