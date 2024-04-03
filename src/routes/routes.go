package routes

import (
	"web/src/handlers"
	model "web/src/model"

	"github.com/gin-gonic/gin"
)

func APIRoutes(router *gin.Engine, env model.Config) {
	router.GET(env.APIPath+"/ping", handlers.RootAPIPath)
	router.GET(env.APIPath+"/users", handlers.GetUserIDHandler)
}

func StaticRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("./public/views/*.html")
	router.NoRoute(func(c *gin.Context) {
		c.File("./public" + c.Request.URL.Path)
	})
}
