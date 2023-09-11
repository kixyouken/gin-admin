package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMailDeliveryTargetsService struct{}

var MailDeliveryTargetsService = sMailDeliveryTargetsService{}

var (
	modelDeliveryTargetsList []model.MailDeliveryTargetsModel
)

func (s *sMailDeliveryTargetsService) DeliveryTargetsInfo(in model.MailDeliveryTargetsModel) (out param.MailDeliveryTargetsInfo) {
	out.ID = in.ID
	out.UserID = in.UserID
	out.MailDeliveryID = in.MailDeliveryID
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)

	return out
}

func (s *sMailDeliveryTargetsService) DeliveryTargetsList(in []model.MailDeliveryTargetsModel) (out []param.MailDeliveryTargetsInfo) {
	for _, v := range in {
		info := s.DeliveryTargetsInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMailDeliveryTargetsService) GetAllDeliveryTargets(c *gin.Context, search func(db *gorm.DB) *gorm.DB, order string) ([]model.MailDeliveryTargetsModel, error) {
	err := BaseService.GetPage(c, &modelDeliveryTargetsList, search, "*", order)
	if err != nil {
		return nil, err
	}

	return modelDeliveryTargetsList, nil
}

func (s *sMailDeliveryTargetsService) GetCountDeliveryTargets(c *gin.Context, search func(db *gorm.DB) *gorm.DB) int64 {
	return BaseService.GetCount(&modelDeliveryTargetsList, search)
}
