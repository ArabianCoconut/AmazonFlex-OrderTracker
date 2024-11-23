package main

import "github.com/gin-gonic/gin"


func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFile("/index.html", "./index.html")
	r.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})
	r.Run(":8080")
}