package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMasterIndustryService struct{}

var MasterIndustryService = sMasterIndustryService{}

var (
	modelIndustryList []model.MasterIndustryModel
)

func (s *sMasterIndustryService) IndustryInfo(in model.MasterIndustryModel) (out param.MasterIndustryInfo) {
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

func (s *sMasterIndustryService) IndustryList(in []model.MasterIndustryModel) (out []param.MasterIndustryInfo) {
	for _, v := range in {
		info := s.IndustryInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMasterIndustryService) GetAllIndustry(c *gin.Context, search func(db *gorm.DB) *gorm.DB) ([]model.MasterIndustryModel, error) {
	err := BaseService.GetAll(&modelIndustryList, search)
	if err != nil {
		return nil, err
	}

	return modelIndustryList, nil
}
