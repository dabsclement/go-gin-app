package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// initializing gin default server
	server := gin.Default()
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok!",
		})
	})
	server.Run(":8080")
}
