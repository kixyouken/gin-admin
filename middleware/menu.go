package middleware

import (
	"gin-admin/service"

	"github.com/gin-gonic/gin"
)

func Menu() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuList, err := service.MenuService.GetAllMenu(c)
		if err != nil {
			return
		}
		list := service.MenuService.MenuList(menuList)
		menuTree := service.MenuService.BuildTree(list, 0)
		c.Set("menuTree", menuTree)
		c.Next()
	}
}
