package call

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
	modelCallList []model.CallModel
	paramCallList []param.CallInfo
	err           error
)

func Index(c *gin.Context) {
	menuTree, _ := c.Get("menuTree")
	c.HTML(http.StatusOK, "call/index", gin.H{
		"title": "会員管理",
		"menu":  menuTree,
		"uri":   "call/list",
	})
}

func Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "call/edit", gin.H{
		"title": "Main website",
	})
}

func Data(c *gin.Context) {
	paramCallSearch := param.CallSearch{}
	c.ShouldBind(&paramCallSearch)
	searchDB := search.CallSearch.SearchCall(paramCallSearch)

	count := service.CallService.GetCountCall(c, searchDB)
	if count > 0 {
		modelCallList, err = service.CallService.GetAllCall(c, searchDB, "")
		if err != nil {
			return
		}
		paramCallList = service.CallService.CallList(modelCallList)
	} else {
		paramCallList = nil
	}

	output.Pageful(c, paramCallList, count)
}

func Update(c *gin.Context) {
	id := convert.GetUint(c, "id")
	paramCallInfo := param.CallInfo{}
	c.ShouldBind(&paramCallInfo)
	column := []string{"name", "status", "detail", "remark"}
	service.CallService.SetByIDCall(c, id, paramCallInfo, column)
	output.Successful(c, "success")
}
