package university

import (
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"

	"github.com/gin-gonic/gin"
)

var (
	modelUniversityList []model.MasterUniversityModel
	paramUniversityList []param.MasterUniversityInfo
	err                 error
)

// Area 大学地区
//
//	@param c
func Area(c *gin.Context) {
	modelUniversityList, err = service.MasterUniversityService.GetMasterUniversityArea("area")
	if err != nil {
		return
	}

	paramUniversityList = service.MasterUniversityService.MasterUniversityList(modelUniversityList)
	output.Dataful(c, paramUniversityList)
}

func Data(c *gin.Context) {
	paramUniversitySearch := param.MasterUniversitySearch{}
	c.ShouldBind(&paramUniversitySearch)
	searchDB := search.UniversitySearch.SearchUniversity(paramUniversitySearch)
	modelUniversityList, err = service.MasterUniversityService.GetAllUniversity(c, searchDB)
	if err != nil {
		return
	}

	paramUniversityList = service.MasterUniversityService.MasterUniversityList(modelUniversityList)
	output.Dataful(c, paramUniversityList)
}
