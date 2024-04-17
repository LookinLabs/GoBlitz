package templates

import (
	"net/http"
	"os"
	model "web/model"

	"github.com/gin-gonic/gin"
)

func serviceHealthHandler(serviceInfo model.ServiceStatusInfo) map[string]string {
	resp, err := http.Get(serviceInfo.URL)
	service := map[string]string{
		"name": serviceInfo.Name,
	}

	if err != nil {
		service["status"] = "down"
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			service["status"] = "up"
		} else {
			service["status"] = "down"
		}
	}

	return service
}

func CheckServicesStatus(services []model.ServiceStatusInfo) []map[string]string {
	statuses := make([]map[string]string, 0)
	for _, service := range services {
		status := serviceHealthHandler(service)
		statuses = append(statuses, status)
	}
	return statuses
}

func StatusPageResponse() gin.HandlerFunc {
	urlPrefix := os.Getenv("FORCE_TLS")
	if urlPrefix == "" {
		urlPrefix = "http://"
	}
	return func(c *gin.Context) {
		services := []model.ServiceStatusInfo{
			{
				Name: "API",
				URL:  urlPrefix + os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT") + os.Getenv("API_PATH") + "ping",
			},
		}

		statuses := CheckServicesStatus(services)

		c.Set("statuses", statuses)
		c.Next()
	}
}
