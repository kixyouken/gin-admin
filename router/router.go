package router

import (
	"gin-admin/controller/index"
	"gin-admin/controller/users"
	"gin-admin/middleware"

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

	router.GET("", index.Index)

	router.GET("user/data", users.Data)

	router.Use(middleware.Menu())
	router.GET("user", users.Index)

	return router
}

func createRender() multitemplate.Renderer {
	tmpl := multitemplate.NewRenderer()
	tmpl.AddFromFiles("index", "templates/base/layout.tmpl", "templates/index/index.tmpl")
	tmpl.AddFromFiles("users/index", "templates/base/layout.tmpl", "templates/users/index.tmpl")
	return tmpl
}
