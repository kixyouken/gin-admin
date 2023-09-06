package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMasterUniversityService struct{}

var MasterUniversityService = sMasterUniversityService{}

var (
	modelUniversityList []model.MasterUniversityModel
)

// MasterUniversityInfo 大学详情
//
//	@receiver s
//	@param in
//	@return out
func (s *sMasterUniversityService) MasterUniversityInfo(in model.MasterUniversityModel) (out param.MasterUniversityInfo) {
	out.ID = in.ID
	out.Name = in.Name
	out.NameKana = in.NameKana
	out.Area = in.Area
	out.Prefectures = in.Prefectures
	out.Class = in.Class
	out.Type = in.Type
	out.Rank = in.Rank
	out.FlagShow = in.FlagShow
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)

	return out
}

// MasterUniversityList 大学列表
//
//	@receiver s
//	@param in
//	@return out
func (s *sMasterUniversityService) MasterUniversityList(in []model.MasterUniversityModel) (out []param.MasterUniversityInfo) {
	for _, v := range in {
		info := s.MasterUniversityInfo(v)
		out = append(out, info)
	}

	return out
}

// GetMasterUniversityArea 大学地区
//
//	@receiver s
//	@param column
//	@return []model.MasterUniversityModel
//	@return error
func (s *sMasterUniversityService) GetMasterUniversityArea(column interface{}) ([]model.MasterUniversityModel, error) {
	err := BaseService.GetDistinct(&modelUniversityList, column)
	if err != nil {
		return nil, err
	}
	return modelUniversityList, nil
}

// GetAllUniversity 大学列表
//
//	@receiver s
//	@param c
//	@param search
//	@return []model.MasterUniversityModel
//	@return error
func (s *sMasterUniversityService) GetAllUniversity(c *gin.Context, search func(db *gorm.DB) *gorm.DB) ([]model.MasterUniversityModel, error) {
	err := BaseService.GetAll(&modelUniversityList, search)
	if err != nil {
		return nil, err
	}

	return modelUniversityList, nil
}
