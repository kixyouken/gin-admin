package scale

import (
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"

	"github.com/gin-gonic/gin"
)

var (
	modelScaleList []model.MasterScaleModel
	paramScaleList []param.MasterScaleInfo
	err            error
)

func Data(c *gin.Context) {
	paramScaleSearch := param.MasterScaleSearch{}
	c.ShouldBind(&paramScaleSearch)
	searchDB := search.ScaleSearch.SearchScale(paramScaleSearch)
	modelScaleList, err = service.MasterScaleService.GetAllScale(c, searchDB)
	if err != nil {
		return
	}

	paramScaleList = service.MasterScaleService.ScaleList(modelScaleList)
	output.Dataful(c, paramScaleList)
}
