package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMasterOrganizationService struct{}

var MasterOrganizationService = sMasterOrganizationService{}

var (
	modelOrganizationList []model.MasterOrganizationModel
)

func (s *sMasterOrganizationService) OrganizationInfo(in model.MasterOrganizationModel) (out param.MasterOrganizationInfo) {
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

func (s *sMasterOrganizationService) OrganizationList(in []model.MasterOrganizationModel) (out []param.MasterOrganizationInfo) {
	for _, v := range in {
		info := s.OrganizationInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMasterOrganizationService) GetAllOrganization(c *gin.Context, search func(db *gorm.DB) *gorm.DB) ([]model.MasterOrganizationModel, error) {
	err := BaseService.GetAll(&modelOrganizationList, search)
	if err != nil {
		return nil, err
	}

	return modelOrganizationList, nil
}
