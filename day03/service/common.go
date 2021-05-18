package service_common

import (
	"github.com/gin-gonic/gin"
)

func state(c *gin.Context) {
	c.JSON(200, gin.H{
		"MSG": "PONG",
	})
}
