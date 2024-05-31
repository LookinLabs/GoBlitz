package views

import (
	"log"
	"net/http"
	"os"

	model "web/model"

	"github.com/gin-gonic/gin"
)

func CheckServicesStatus(services []model.StatusPage) []map[string]string {
	serviceStatuses := make([]map[string]string, 0)
	for _, service := range services {
		status := servicesHealthHandler(service)
		serviceStatuses = append(serviceStatuses, status)
	}

	return serviceStatuses
}

func StatusPageResponse() gin.HandlerFunc {
	urlPrefix := "http://"
	if os.Getenv("FORCE_TLS") == "true" {
		urlPrefix = "https://"
	}

	return func(c *gin.Context) {
		services := []model.StatusPage{
			{
				Name: "API",
				URL:  urlPrefix + os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT") + os.Getenv("API_PATH") + "ping",
			},
			{
				Name: "Users API",
				URL:  urlPrefix + os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT") + os.Getenv("API_PATH") + "users",
			},
		}

		statuses := CheckServicesStatus(services)

		c.Set("statuses", statuses)
		c.Next()
	}
}

func servicesHealthHandler(serviceInfo model.StatusPage) map[string]string {
	resp, err := http.Get(serviceInfo.URL)

	service := map[string]string{
		"name":   serviceInfo.Name,
		"status": "up",
	}

	if err != nil {
		log.Println("Error checking service status:", err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			service["status"] = "down"
		}
	}

	return service
}
