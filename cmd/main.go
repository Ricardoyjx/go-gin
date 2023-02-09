package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	r.GET("/ask", func(c *gin.Context) {
		c.String(http.StatusOK, "who you are?")
	})

	r.GET("/ask/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "i am %s !", name)
	})

	r.Run() //default listen and serve on :8080
}
