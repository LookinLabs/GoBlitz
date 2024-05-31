package helper

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func ServePageAssets(router *gin.Engine) {
	if CheckFileExists("./public/index.html") {
		router.Use(static.Serve("/assets", static.LocalFile("./public/assets", true)))
	} else {
		router.Use(static.Serve("/assets", static.LocalFile("./views/assets", true)))
		assetsPerPage(router)
	}
}

func assetsPerPage(router *gin.Engine) {
	assetsDir, err := os.ReadDir("./views")
	if err != nil {
		log.Fatal(err)
	}

	for _, directory := range assetsDir {
		if directory.IsDir() && strings.Contains(directory.Name(), "_page") {
			router.Use(static.Serve("/"+directory.Name()+"/assets", static.LocalFile("./views/"+directory.Name()+"/assets", true)))
		}
	}
}
