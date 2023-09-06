package location

import (
	"gin-admin/common/convert"
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"

	"github.com/gin-gonic/gin"
)

var (
	modelLocationList []model.MasterLocationModel
	paramLocationList []param.MasterLocationInfo
	err               error
)

func Data(c *gin.Context) {
	paramLocationSearch := param.MasterLocationSearch{}
	c.ShouldBind(&paramLocationSearch)
	searchDB := search.LocationSearch.SearchLocation(paramLocationSearch)
	modelLocationList, err = service.MasterLocationService.GetAllLocation(c, searchDB)
	if err != nil {
		return
	}

	paramLocationList = service.MasterLocationService.LocationList(modelLocationList)
	output.Dataful(c, paramLocationList)
}

func Detail(c *gin.Context) {
	id := convert.GetUint(c, "id")
	modelLocationInfo, err := service.MasterLocationService.GetByIDLocation(c, id)
	if err != nil {
		return
	}
	paramLocationInfo := service.MasterLocationService.LocationInfo(*modelLocationInfo)
	output.Dataful(c, paramLocationInfo)
}
