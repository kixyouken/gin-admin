package router

import (
	"gin-admin/controller"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// 模板文件(默认单模板)
	// router.LoadHTMLGlob("templates/**/*")
	// 模板文件(开启多模板)
	router.HTMLRender = createRender()
	// 静态文件
	router.Static("/assets", "./assets")

	router.GET("", controller.Index)
	router.GET("/read", controller.Read)

	return router
}

func createRender() multitemplate.Renderer {
	tmpl := multitemplate.NewRenderer()
	tmpl.AddFromFiles("index", "templates/base/layout.tmpl", "templates/index/index.tmpl")
	tmpl.AddFromFiles("read", "templates/base/layout.tmpl", "templates/index/read.tmpl")
	return tmpl
}
