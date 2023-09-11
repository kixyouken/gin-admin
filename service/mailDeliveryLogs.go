package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMailDeliveryLogsService struct{}

var MailDeliveryLogsService = sMailDeliveryLogsService{}

var (
	modelDeliveryLogsList []model.MailDeliveryLogsModel
)

func (s *sMailDeliveryLogsService) MailDeliveryLogsInfo(in model.MailDeliveryLogsModel) (out param.MailDeliveryLogsInfo) {
	out.ID = in.ID
	out.MailDeliveryID = in.MailDeliveryID
	out.Title = in.Title
	out.SendAt = in.SendAt.Format(format.YMDHI)
	out.Sender = in.Sender
	out.Content = in.Content
	out.UserID = in.UserID
	out.FlagOpened = in.FlagOpened
	out.LinkClickTimes = in.LinkClickTimes
	out.Recycled = in.Recycled
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)

	return out
}

func (s *sMailDeliveryLogsService) MailDeliveryLogsList(in []model.MailDeliveryLogsModel) (out []param.MailDeliveryLogsInfo) {
	for _, v := range in {
		info := s.MailDeliveryLogsInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMailDeliveryLogsService) GetAllDeliveryLogs(c *gin.Context, search func(db *gorm.DB) *gorm.DB, order string) ([]model.MailDeliveryLogsModel, error) {
	err := BaseService.GetPage(c, &modelDeliveryLogsList, search, "*", order)
	if err != nil {
		return nil, err
	}

	return modelDeliveryLogsList, nil
}

func (s *sMailDeliveryLogsService) GetCountDeliveryLogs(c *gin.Context, search func(db *gorm.DB) *gorm.DB) int64 {
	return BaseService.GetCount(&modelDeliveryLogsList, search)
}
