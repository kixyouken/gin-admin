package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"
)

type sMasterSignService struct{}

var MasterSignService = sMasterSignService{}

func (s *sMasterSignService) SignInfo(in model.MasterSignModel) (out param.MasterSignInfo) {
	out.ID = in.ID
	out.Name = in.Name
	out.Content = in.Content
	out.FlagShow = in.FlagShow
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)
	if !in.DeletedAt.Time.IsZero() {
		out.DeletedAt = in.DeletedAt.Time.Format(format.YMDHI)
	}

	return out
}

func (s *sMasterSignService) SignList(in []model.MasterSignModel) (out []param.MasterSignInfo) {
	for _, v := range in {
		info := s.SignInfo(v)
		out = append(out, info)
	}

	return out
}
