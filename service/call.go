package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sCallService struct{}

var CallService = sCallService{}

var (
	modelCallList []model.CallModel
)

func (s *sCallService) CallInfo(in model.CallModel) (out param.CallInfo) {
	out.ID = in.ID
	out.Name = in.Name
	out.Status = in.Status
	out.Detail = in.Detail
	out.Remark = in.Remark
	out.CreatedBy = in.CreatedBy
	out.UpdatedBy = in.UpdatedBy
	out.DeletedBy = in.DeletedBy
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)
	if !in.DeletedAt.Time.IsZero() {
		out.DeletedAt = in.DeletedAt.Time.Format(format.YMDHI)
	}

	out.CallTargets = CallTargetsService.CallTargetsList(in.CallTargets)

	return out
}

func (s *sCallService) CallList(in []model.CallModel) (out []param.CallInfo) {
	for _, v := range in {
		info := s.CallInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sCallService) GetAllCall(c *gin.Context, search func(db *gorm.DB) *gorm.DB, order string) ([]model.CallModel, error) {
	err := BaseService.GetPagePreload(c, &modelCallList, search, "*", order, "CallTargets")
	if err != nil {
		return nil, err
	}

	return modelCallList, nil
}

func (s *sCallService) GetCountCall(c *gin.Context, search func(db *gorm.DB) *gorm.DB) int64 {
	return BaseService.GetCount(&modelCallList, search)
}

func (s *sCallService) SetByIDCall(c *gin.Context, id uint, updates interface{}, column interface{}) (*model.CallModel, error) {
	modelCallInfo := model.CallModel{}
	err := BaseService.UpdateByID(&modelCallInfo, id, updates, column)
	if err != nil {
		return nil, err
	}

	return &modelCallInfo, nil
}
