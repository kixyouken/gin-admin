package industry

import (
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"

	"github.com/gin-gonic/gin"
)

var (
	modelIndustryList []model.MasterIndustryModel
	paramIndustryList []param.MasterIndustryInfo
	err               error
)

func Data(c *gin.Context) {
	paramIndustrySearch := param.MasterIndustrySearch{}
	c.ShouldBind(&paramIndustrySearch)
	searchDB := search.IndustrySearch.SearchIndustry(paramIndustrySearch)
	modelIndustryList, err = service.MasterIndustryService.GetAllIndustry(c, searchDB)
	if err != nil {
		return
	}

	paramIndustryList = service.MasterIndustryService.IndustryList(modelIndustryList)
	output.Dataful(c, paramIndustryList)
}
