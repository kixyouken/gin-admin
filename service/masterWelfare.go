package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMasterWelfareService struct{}

var MasterWelfareService = sMasterWelfareService{}

var (
	modelWelfareList []model.MasterWelfareModel
)

func (s *sMasterWelfareService) WelfareInfo(in model.MasterWelfareModel) (out param.MasterWelfareInfo) {
	out.ID = in.ID
	out.Name = in.Name
	out.FlagShow = in.FlagShow
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)
	out.DeletedAt = in.DeletedAt.Time.Format(format.YMDHI)

	return out
}

func (s *sMasterWelfareService) WelfareList(in []model.MasterWelfareModel) (out []param.MasterWelfareInfo) {
	for _, v := range in {
		info := s.WelfareInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMasterWelfareService) GetAllWelfare(c *gin.Context, search func(db *gorm.DB) *gorm.DB) ([]model.MasterWelfareModel, error) {
	err := BaseService.GetAll(&modelWelfareList, search)
	if err != nil {
		return nil, err
	}

	return modelWelfareList, nil
}
