package router

import (
	"gin-admin/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("", controller.Index)

	return router
}
