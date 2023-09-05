package job

import (
	"gin-admin/common/output"
	"gin-admin/model"
	"gin-admin/param"
	"gin-admin/search"
	"gin-admin/service"

	"github.com/gin-gonic/gin"
)

var (
	modelJobList []model.MasterJobModel
	paramJobList []param.MasterJobInfo
	err          error
)

func Data(c *gin.Context) {
	paramJobSearch := param.MasterJobSearch{}
	c.ShouldBind(&paramJobSearch)
	searchDB := search.JobSearch.SearchJob(paramJobSearch)
	modelJobList, err = service.MasterJobService.GetAllJob(c, searchDB)
	if err != nil {
		return
	}
	paramJobList = service.MasterJobService.JobList(modelJobList)
	output.Dataful(c, paramJobList)
}
