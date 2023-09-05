package organization

import (
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"

	"github.com/gin-gonic/gin"
)

var (
	modelOrganizationList []model.MasterOrganizationModel
	paramOrganizationList []param.MasterOrganizationInfo
	err                   error
)

func Data(c *gin.Context) {
	paramOrganizationSearch := param.MasterOrganizationSearch{}
	c.ShouldBind(&paramOrganizationSearch)
	searchDB := search.OrganizationSearch.SearchOrganization(&paramOrganizationSearch)
	modelOrganizationList, err = service.MasterOrganizationService.GetAllOrganization(c, searchDB)
	if err != nil {
		return
	}

	paramOrganizationList = service.MasterOrganizationService.OrganizationList(modelOrganizationList)
	output.Dataful(c, paramOrganizationList)
}
