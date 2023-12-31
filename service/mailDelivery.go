package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sMailDeliveryService struct{}

var MailDeliveryService = sMailDeliveryService{}

var (
	modelDeliveryList []model.MailDeliveryModel
)

func (s *sMailDeliveryService) DeliveryInfo(in model.MailDeliveryModel) (out param.MailDeliveryInfo) {
	out.ID = in.ID
	out.Name = in.Name
	out.Title = in.Title
	out.SenderID = in.SenderID
	out.SendAt = in.SendAt.Format(format.YMDHI)
	out.SendCompletedAt = in.SendCompletedAt.Format(format.YMDHI)
	out.FlagImportant = in.FlagImportant
	out.FlagImmediate = in.FlagImmediate
	out.Content = in.Content
	out.MasterSignID = in.MasterSignID
	out.Status = in.Status
	if in.Status == 1 {
		out.StatusText = "配信完了"
	} else {
		out.StatusText = "未実施"
	}
	out.MailContentConfirm = in.MailContentConfirm
	out.MailAddressConfirm = in.MailAddressConfirm
	out.SummaryCategory = in.SummaryCategory
	switch in.SummaryCategory {
	case 0:
		out.SummaryCategoryText = "-"
	case 1:
		out.SummaryCategoryText = "ジョブオファー"
	case 2:
		out.SummaryCategoryText = "イベント"
	}
	out.MailConfirmBy = in.MailConfirmBy
	out.MailConfirmAt = in.MailConfirmAt.Format(format.YMDHI)
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)
	if !in.DeletedAt.Time.IsZero() {
		out.DeletedAt = in.DeletedAt.Time.Format(format.YMDHI)
	}

	out.MailDeliveryLogs = MailDeliveryLogsService.MailDeliveryLogsList(in.MailDeliveryLogs)
	out.MailDeliveryTargets = MailDeliveryTargetsService.DeliveryTargetsList(in.MailDeliveryTargets)
	out.MailSenders = MailSendersService.SendersInfo(in.MailSenders)
	out.MasterSign = MasterSignService.SignInfo(in.MasterSign)

	return out
}

func (s *sMailDeliveryService) DeliveryList(in []model.MailDeliveryModel) (out []param.MailDeliveryInfo) {
	for _, v := range in {
		info := s.DeliveryInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sMailDeliveryService) GetAllDelivery(c *gin.Context, search func(db *gorm.DB) *gorm.DB, order string) ([]model.MailDeliveryModel, error) {
	err := BaseService.GetPagePreload(c, &modelDeliveryList, search, "*", order, "MailDeliveryLogs.Unscoped", "MailDeliveryTargets.Unscoped")
	if err != nil {
		return nil, err
	}

	return modelDeliveryList, nil
}

func (s *sMailDeliveryService) GetCountDelivery(c *gin.Context, search func(db *gorm.DB) *gorm.DB) int64 {
	return BaseService.GetCount(&modelDeliveryList, search)
}

func (s *sMailDeliveryService) GetByIDDelivery(c *gin.Context, id uint) (*model.MailDeliveryModel, error) {
	modelDeliveryInfo := model.MailDeliveryModel{}
	err := BaseService.GetByID(&modelDeliveryInfo, id, "MailSenders", "MasterSign")
	if err != nil {
		return nil, err
	}

	return &modelDeliveryInfo, nil
}
