package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello-world", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!!!!",
		})
	})

	router.Run()

}
