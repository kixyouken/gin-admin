package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMasterJobService struct{}

var MasterJobService = sMasterJobService{}

var (
	modelJobList []model.MasterJobModel
)

func (s *sMasterJobService) JobInfo(in model.MasterJobModel) (out param.MasterJobInfo) {
	out.ID = in.ID
	out.Name = in.Name
	out.FlagShow = in.FlagShow
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)
	if !in.DeletedAt.Time.IsZero() {
		out.DeletedAt = in.DeletedAt.Time.Format(format.YMDHI)
	}
	return out
}

func (s *sMasterJobService) JobList(in []model.MasterJobModel) (out []param.MasterJobInfo) {
	for _, v := range in {
		info := s.JobInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMasterJobService) GetAllJob(c *gin.Context, search func(db *gorm.DB) *gorm.DB) ([]model.MasterJobModel, error) {
	err := BaseService.GetAll(&modelJobList, search)
	if err != nil {
		return nil, err
	}

	return modelJobList, nil
}
