package users

import (
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	modelUserList  []model.UsersModel
	paramUsersList []param.UsersInfo
	err            error
)

func Index(c *gin.Context) {
	menuTree, _ := c.Get("menuTree")
	c.HTML(http.StatusOK, "users/index", gin.H{
		"title": "Main website",
		"menu":  menuTree,
	})
}

func Data(c *gin.Context) {
	paramUsersInfo := param.UsersInfo{}
	c.BindQuery(&paramUsersInfo)
	search := service.SearchService.SearchUser(paramUsersInfo)
	count := service.UsersService.GetCountUsers(c, search)
	if count > 0 {
		modelUserList, err = service.UsersService.GetAllUsers(c, search, "")
		if err != nil {
			return
		}
		paramUsersList = service.UsersService.UsersList(modelUserList)
	} else {
		paramUsersList = nil
	}
	output.Pageful(c, paramUsersList, count)
}
