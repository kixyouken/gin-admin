package main

import (
	"gin-admin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	router := router.GetRouter()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 监听并在 0.0.0.0:9000 上启动服务
	router.Run(":9000")
}
