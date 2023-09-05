package router

import (
	"gin-admin/controller/index"
	"gin-admin/controller/master/industry"
	"gin-admin/controller/master/job"
	"gin-admin/controller/master/negotiate"
	"gin-admin/controller/master/organization"
	"gin-admin/controller/master/science"
	"gin-admin/controller/master/university"
	"gin-admin/controller/master/welfare"
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

	router.GET("user/data", users.Data)
	router.GET("user/read", users.Read)
	router.GET("user/edit", users.Edit)
	router.POST("user/detail/:id", users.Detail)
	router.GET("university/area", university.Area)
	router.GET("university/data", university.Data)
	router.GET("organization/data", organization.Data)
	router.GET("science/data", science.Data)
	router.GET("negotiate/data", negotiate.Data)
	router.GET("industry/data", industry.Data)
	router.GET("job/data", job.Data)
	router.GET("welfare/data", welfare.Data)

	router.Use(middleware.Menu())
	router.GET("user", users.Index)
	router.GET("", index.Index)

	return router
}

func createRender() multitemplate.Renderer {
	tmpl := multitemplate.NewRenderer()
	tmpl.AddFromFiles("index", "templates/base/layout.tmpl", "templates/index/index.tmpl")
	tmpl.AddFromFiles("users/index", "templates/base/layout.tmpl", "templates/users/index.tmpl")
	tmpl.AddFromFiles("users/edit", "templates/base/iframe.tmpl", "templates/users/edit.tmpl")
	return tmpl
}
