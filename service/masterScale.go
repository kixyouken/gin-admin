package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMasterScaleService struct{}

var MasterScaleService = sMasterScaleService{}

var (
	modelScaleList []model.MasterScaleModel
)

func (s *sMasterScaleService) ScaleInfo(in model.MasterScaleModel) (out param.MasterScaleInfo) {
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

func (s *sMasterScaleService) ScaleList(in []model.MasterScaleModel) (out []param.MasterScaleInfo) {
	for _, v := range in {
		info := s.ScaleInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMasterScaleService) GetAllScale(c *gin.Context, search func(db *gorm.DB) *gorm.DB) ([]model.MasterScaleModel, error) {
	err := BaseService.GetAll(&modelScaleList, search)
	if err != nil {
		return nil, err
	}

	return modelScaleList, nil
}
