package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": getPong(),
		})
	})

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run()
}

func getPong() string {
	return "pong"
}
