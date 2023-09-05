package welfare

import (
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"

	"github.com/gin-gonic/gin"
)

var (
	modelWelfareList []model.MasterWelfareModel
	paramWelfareList []param.MasterWelfareInfo
	err              error
)

func Data(c *gin.Context) {
	paramWelfareSearch := param.MasterWelfareSearch{}
	c.ShouldBind(&paramWelfareSearch)
	searchDB := search.WelfareSearch.SearchWelfare(paramWelfareSearch)
	modelWelfareList, err = service.MasterWelfareService.GetAllWelfare(c, searchDB)
	if err != nil {
		return
	}

	paramWelfareList = service.MasterWelfareService.WelfareList(modelWelfareList)
	output.Dataful(c, paramWelfareList)
}
