package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMasterLocationService struct{}

var MasterLocationService = sMasterLocationService{}

var (
	modelLocationList []model.MasterLocationModel
)

func (s *sMasterLocationService) LocationInfo(in model.MasterLocationModel) (out param.MasterLocationInfo) {
	out.ID = in.ID
	out.Name = in.Name
	out.FlagShow = in.FlagShow
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)
	out.DeletedAt = in.DeletedAt.Time.Format(format.YMDHI)

	return out
}

func (s *sMasterLocationService) LocationList(in []model.MasterLocationModel) (out []param.MasterLocationInfo) {
	for _, v := range in {
		info := s.LocationInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMasterLocationService) GetAllLocation(c *gin.Context, search func(db *gorm.DB) *gorm.DB) ([]model.MasterLocationModel, error) {
	err := BaseService.GetAll(&modelLocationList, search)
	if err != nil {
		return nil, err
	}

	return modelLocationList, nil
}

func (s *sMasterLocationService) GetByIDLocation(c *gin.Context, id uint) (*model.MasterLocationModel, error) {
	modelLocationInfo := model.MasterLocationModel{}
	err := BaseService.GetByID(&modelLocationInfo, id)
	if err != nil {
		return nil, err
	}

	return &modelLocationInfo, nil
}
