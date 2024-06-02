package helper

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func ServePageAssets(router *gin.Engine) {
	if CheckIfFileExists("./public/index.html") {
		router.Use(static.Serve("/assets", static.LocalFile("./public/assets", true)))
	} else {
		router.Use(static.Serve("/assets", static.LocalFile("./views/assets", false)))
		assetsPerPage(router)
	}
}

func assetsPerPage(router *gin.Engine) {
	views, err := os.ReadDir("./views")
	if err != nil {
		log.Fatal(err)
	}

	for _, directory := range views {
		if directory.IsDir() && strings.Contains(directory.Name(), "_page") {
			assetsPath := "./views/" + directory.Name() + "/assets"
			if CheckFileNotExists(assetsPath) {
				continue
			}

			// Serve page-specific assets
			router.Use(static.Serve("/"+directory.Name()+"/assets", static.LocalFile(assetsPath, false)))
		}
	}
}
