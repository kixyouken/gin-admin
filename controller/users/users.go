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

func View(c *gin.Context) {
	c.HTML(http.StatusOK, "users/view", gin.H{
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

func Read(c *gin.Context) {
	id := convert.GetUint(c, "id")
	modelUsersInfo, err := service.UsersService.GetByIDUsers(c, id)
	if err != nil {
		return
	}
	info := service.UsersService.UsersInfo(*modelUsersInfo)
	info = service.UsersService.GetUsersString(info)
	output.Dataful(c, info)
}

func Update(c *gin.Context) {
	id := convert.GetUint(c, "id")
	paramUsersInfo := param.UsersInfo{}
	c.ShouldBind(&paramUsersInfo)
	column := []string{"family_name", "name", "email", "mobile", "birthday", "master_university_id", "university_name_opt", "graduation_year", "graduation_month", "kana_family_name", "kana_name", "faculty", "department", "type", "status", "ap_informal_offer", "prospective_destination", "master_organization_id", "organization_name_opt", "master_science_id", "master_negotiate_id", "master_industry_id", "master_job_id", "master_welfare_id", "master_location_id", "master_scale_id", "postal_code", "address_area", "address_city", "address", "address_building", "self_pr", "extracurricular", "ap_comment", "ap_remark", "unsubscribe", "flag_telng", "flag_withdrawal"}
	service.UsersService.SetByIDUsers(c, id, paramUsersInfo, column)
	output.Successful(c, "success")
}
