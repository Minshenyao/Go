package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/example01", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "example01"})
	})
	r.Run(":5700") //默认端口为8080
}
