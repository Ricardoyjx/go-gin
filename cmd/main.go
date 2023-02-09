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

	r.GET("/query", func(c *gin.Context) {
		name := c.Param("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	r.GET("/ask/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "i am %s !", name)
	})

	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.Query("page")
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"id":       id,
			"page":     page,
		})
	})

	//map传参
	r.POST("/map", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})

	})

	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "123456")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})

	})

	r.Run() //default listen and serve on :8080
}
