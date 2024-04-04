package handlers

import (
	"net/http"
	"web/model"

	"github.com/gin-gonic/gin"
)

func checkServiceStatus(service model.ServiceStatusInfo) map[string]string {
	resp, err := http.Get(service.URL)
	status := map[string]string{
		"name": service.Name,
	}

	if err != nil {
		status["status"] = "down"
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			status["status"] = "up"
		} else {
			status["status"] = "down"
		}
	}

	return status
}

func ReviewServiceStatus(services []model.ServiceStatusInfo) []map[string]string {
	statuses := make([]map[string]string, 0)
	for _, service := range services {
		status := checkServiceStatus(service)
		statuses = append(statuses, status)
	}
	return statuses
}

func StatusHandler(env model.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		services := []model.ServiceStatusInfo{
			{Name: "API", URL: env.URLPrefix + env.AppHost + ":" + env.AppPort + env.APIPath + "ping"},
		}

		statuses := ReviewServiceStatus(services)

		c.Set("statuses", statuses)
		c.Next()
	}
}
