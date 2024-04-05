package templates

import (
	"net/http"
	envModel "web/model/config"
	templatesModel "web/model/templates"

	"github.com/gin-gonic/gin"
)

func serviceHealthHandler(serviceInfo templatesModel.ServiceStatusInfo) map[string]string {
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

func CheckServicesStatus(services []templatesModel.ServiceStatusInfo) []map[string]string {
	statuses := make([]map[string]string, 0)
	for _, service := range services {
		status := serviceHealthHandler(service)
		statuses = append(statuses, status)
	}
	return statuses
}

func StatusPageResponse(env envModel.AppEnv) gin.HandlerFunc {
	return func(c *gin.Context) {
		services := []templatesModel.ServiceStatusInfo{
			{Name: "API", URL: env.URLPrefix + env.AppHost + ":" + env.AppPort + env.APIPath + "ping"},
		}

		statuses := CheckServicesStatus(services)

		c.Set("statuses", statuses)
		c.Next()
	}
}
