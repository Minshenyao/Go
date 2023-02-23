package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/User", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "User Page", "method": "Get"})
	})
	r.POST("/User", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "User Page", "method": "Post"})
	})
	r.PUT("/User", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "User Page", "method": "Put"})
	})
	r.DELETE("/User", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "User Page", "method": "Delete"})
	})
	r.Run()
}
