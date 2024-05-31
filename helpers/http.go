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
	dirs, err := os.ReadDir("./views")
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() && strings.Contains(dir.Name(), "_page") {
			router.Use(static.Serve("/"+dir.Name()+"/assets", static.LocalFile("./views/"+dir.Name()+"/assets", true)))
		}
	}
}
