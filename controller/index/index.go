package index

import (
	"gin-admin/common/output"
	"gin-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	menuTree, _ := c.Get("menuTree")
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "ダッシュボード",
		"menu":  menuTree,
		"uri":   "/",
	})
}

func Menu(c *gin.Context) {
	menuList, err := service.MenuService.GetAllMenu(c)
	if err != nil {
		return
	}
	list := service.MenuService.MenuList(menuList)
	menuTree := service.MenuService.BuildTree(list, 0)
	output.Dataful(c, menuTree)
}
