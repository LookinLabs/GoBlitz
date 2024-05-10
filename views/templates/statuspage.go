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
	req, err := http.NewRequest(http.MethodGet, serviceInfo.URL, nil)
	if os.Getenv("API_PUBLIC_ACCESS") != "true" {
		token := os.Getenv("AWS_COGNITO_JWT_TOKEN")
		if token == "" {
			log.Println("AWS_COGNITO_JWT_TOKEN is not set")
			log.Println("Please set the environment variable AWS_COGNITO_JWT_TOKEN")

			return map[string]string{
				"name":   serviceInfo.Name,
				"status": "down",
			}
		}
		req.Header.Add("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	if err != nil {
		log.Println("Error creating request:", err)
		return map[string]string{
			"name":   serviceInfo.Name,
			"status": "down",
		}
	}

	resp, err := client.Do(req)

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
