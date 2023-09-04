package users

import (
	"gin-admin/common/convert"
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	modelUsersList []model.UsersModel
	paramUsersList []param.UsersInfo
	err            error
)

// Index 列表模板
//
//	@param c
func Index(c *gin.Context) {
	menuTree, _ := c.Get("menuTree")
	c.HTML(http.StatusOK, "users/index", gin.H{
		"title": "Main website",
		"menu":  menuTree,
	})
}

// Data 列表数据
//
//	@param c
func Data(c *gin.Context) {
	paramUsersSearch := param.UsersSearch{}
	c.ShouldBind(&paramUsersSearch)
	search := search.UsersSearch.SearchUsers(paramUsersSearch)
	count := service.UsersService.GetCountUsers(c, search)
	if count > 0 {
		modelUsersList, err = service.UsersService.GetAllUsers(c, search, "")
		if err != nil {
			return
		}
		paramUsersList = service.UsersService.UsersList(modelUsersList)
	} else {
		paramUsersList = nil
	}
	output.Pageful(c, paramUsersList, count)
}

func Read(c *gin.Context) {
	c.HTML(http.StatusOK, "users/read", gin.H{
		"title": "Main website",
	})
}

func Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "users/edit", gin.H{
		"title": "Main website",
	})
}

func Detail(c *gin.Context) {
	id := convert.GetUint(c, "id")
	modelUsersInfo, err := service.UsersService.GetByIDUsers(c, id)
	if err != nil {
		return
	}
	info := service.UsersService.UsersInfo(*modelUsersInfo)
	output.Dataful(c, info)
}
