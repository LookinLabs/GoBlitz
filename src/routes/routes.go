package routes

import (
	"net/http"
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
	router.StaticFS("/public", http.Dir("./public"))
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
}
