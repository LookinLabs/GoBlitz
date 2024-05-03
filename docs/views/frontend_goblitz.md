# FrontEnd development via GoBlitz

## Assets

Assets (images, gifs, css, js files etc) are served from public folder. The relative path is `/assets`.

You can look an example of importing assets in the [Welcome Page Code](../../views/welcome_page/welcome.html).

## Using HTML Templates to setup pages

In GoBlitz Views are similar to react, where you can define template components in views/components folder and then use them on the page.
You can take a look for the [Welcome Page](../../views/welcome_page/welcome.html)

The pages are separated into folders like views/welcome_page/; views/status_page etc.

HTML Template values are generated via [views/templates folder](../../views/templates/).

The middleware is responsible for serving the HTML Template, you can take a look at [Middleware serving HTML Template Sites](../../middleware/sites.go) and [HTTP Router](../../middleware/router.go) on line 40 & 61

## Error pages

Error pages are handled by the [error controller](../../controller/error/http_errors.go). 

Controller look whether the path is available or not and returns the view into middleware, which returns the response to user.

## An example

1. Set up and Middleware `sites.go` which checks for the template

```
// middleware/sites.go


// StatusPageMiddleware handles the status page
func StatusPageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		statuses := c.MustGet("statuses").([]map[string]string)
		c.HTML(http.StatusOK, "status.html", gin.H{
			"services": statuses,
		})
	}
}
```
2. Setup an route for the status_page

```
// middleware/router.go

router.GET("/status", httpTemplates.StatusPageResponse(), StatusPageMiddleware())
```

3. Configure HTML Template values generator

```
// views/templates/statuspage.go

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
```

4. Setup an status_page

```
// views/status_page/status.html

<!DOCTYPE html>
<html>
<head>
    <title>Status Page</title>
    <link rel="stylesheet" href="/assets/main.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.2/css/all.min.css">
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-gray-100 flex flex-col min-h-screen justify-between">
    {{template "header.html"}}
    <main class="px-4 py-8">
        <h1 class="text-2xl font-bold mb-4 text-center">Status Page</h1>
        {{range .services}}
        <div class="service bg-white shadow rounded-lg p-4 mb-4 flex justify-between items-center w-full">
            <h2 class="text-xl font-bold mb-2">{{.name}}</h2>
            {{if eq .status "up"}}
            <div class="status text-green-500 mb-2">
                <span>&#10003;</span>
            </div>
            {{else}}
            <div class="status text-red-500">
                <span>&#10007;</span>
            </div>
            {{end}}
        </div>
        {{end}}
    </main>
    {{template "footer.html"}}
</body>
</html>
```

Navigate into "APP_HOST:APP_PORT/status" for the results.

In development case it can be `http://localhost:8000/status".