package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMasterNegotiateService struct{}

var MasterNegotiateService = sMasterNegotiateService{}

var (
	modelNegotiateList []model.MasterNegotiateModel
)

func (s *sMasterNegotiateService) NegotiateInfo(in model.MasterNegotiateModel) (out param.MasterNegotiateInfo) {
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

func (s *sMasterNegotiateService) NegotiateList(in []model.MasterNegotiateModel) (out []param.MasterNegotiateInfo) {
	for _, v := range in {
		info := s.NegotiateInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMasterNegotiateService) GetAllNegotiate(c *gin.Context, search func(db *gorm.DB) *gorm.DB) ([]model.MasterNegotiateModel, error) {
	err := BaseService.GetAll(&modelNegotiateList, search)
	if err != nil {
		return nil, err
	}

	return modelNegotiateList, nil
}
