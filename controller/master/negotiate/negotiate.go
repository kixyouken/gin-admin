package negotiate

import (
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"

	"github.com/gin-gonic/gin"
)

var (
	modelNegotiateList []model.MasterNegotiateModel
	paramNegotiateList []param.MasterNegotiateInfo
	err                error
)

func Data(c *gin.Context) {
	paramNegotiateSearch := param.MasterNegotiateSearch{}
	c.ShouldBind(&paramNegotiateSearch)
	searchDB := search.NegotiateSearch.SearchNegotiate(paramNegotiateSearch)
	modelNegotiateList, err = service.MasterNegotiateService.GetAllNegotiate(c, searchDB)
	if err != nil {
		return
	}

	paramNegotiateList = service.MasterNegotiateService.NegotiateList(modelNegotiateList)
	output.Dataful(c, paramNegotiateList)
}
