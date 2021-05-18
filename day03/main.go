package main

import (
	"github.com/gin-gonic/gin"
)

//gin框架，基本的get响应
func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"MSG": "PONG",
		})
	})
	router.Run(":8082")
}
