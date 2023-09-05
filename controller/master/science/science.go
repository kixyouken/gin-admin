package science

import (
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"

	"github.com/gin-gonic/gin"
)

var (
	modelScienceList []model.MasterScienceModel
	paramScienceList []param.MasterScienceInfo
	err              error
)

func Data(c *gin.Context) {
	paramMasterScienceSearch := param.MasterScienceSearch{}
	c.ShouldBind(&paramMasterScienceSearch)
	searchDB := search.ScienceSearch.SearchScience(paramMasterScienceSearch)
	modelScienceList, err = service.MasterScienceService.GetAllScience(c, searchDB)
	if err != nil {
		return
	}

	paramScienceList = service.MasterScienceService.ScienceList(modelScienceList)
	output.Dataful(c, paramScienceList)
}
