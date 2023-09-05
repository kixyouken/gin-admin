package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMasterScienceService struct{}

var MasterScienceService = sMasterScienceService{}

var (
	modelScienceList []model.MasterScienceModel
)

func (s *sMasterScienceService) ScienceInfo(in model.MasterScienceModel) (out param.MasterScienceInfo) {
	out.ID = in.ID
	out.ResearchAreas = in.ResearchAreas
	out.FlagShow = in.FlagShow
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)
	out.DeletedAt = in.DeletedAt.Time.Format(format.YMDHI)

	return out
}

func (s *sMasterScienceService) ScienceList(in []model.MasterScienceModel) (out []param.MasterScienceInfo) {
	for _, v := range in {
		info := s.ScienceInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMasterScienceService) GetAllScience(c *gin.Context, search func(db *gorm.DB) *gorm.DB) ([]model.MasterScienceModel, error) {
	err := BaseService.GetAll(&modelScienceList, search)
	if err != nil {
		return nil, err
	}

	return modelScienceList, nil
}
