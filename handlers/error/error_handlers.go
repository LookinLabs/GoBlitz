package errorhandler

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Show 404 Not Found error page
func StatusNotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
	file, err := os.Open("./public/error/404.html")
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()
	if _, err := io.Copy(c.Writer, file); err != nil {
		if err := c.Error(err); err != nil {
			log.Printf("Error adding error to Gin context: %v", err)
		}
	}
}

// Show 502 Bad Gateway error page
func StatusBadGateway() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if request := recover(); request != nil {
				c.Status(http.StatusBadGateway)
				file, err := os.Open("./public/error/502.html")
				if err != nil {
					log.Printf("Error opening file: %v", err)
					return
				}
				defer file.Close()
				if _, err := io.Copy(c.Writer, file); err != nil {
					if err := c.Error(err); err != nil {
						log.Printf("Error adding error to Gin context: %v", err)
					}
				}
			}
		}()
		c.Next()
	}
}
