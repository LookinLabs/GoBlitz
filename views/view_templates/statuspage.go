package view

import (
	"log"
	"net/http"
	"os"
	model "web/model"

	"github.com/gin-gonic/gin"
)

func serviceHealthHandler(serviceInfo model.ServiceStatusInfo) map[string]string {
	log.Println("Checking service status at URL:", serviceInfo.URL)
	resp, err := http.Get(serviceInfo.URL)
	service := map[string]string{
		"name": serviceInfo.Name,
	}

	if err != nil {
		log.Println("Error checking service status:", err)
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
	urlPrefix := "http://"
	if os.Getenv("FORCE_TLS") == "true" {
		urlPrefix = "https://"
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
