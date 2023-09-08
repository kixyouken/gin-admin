package service

import (
	"gin-admin/model"
	"gin-admin/param"
)

type sCallTargetsService struct{}

var CallTargetsService = sCallTargetsService{}

func (s *sCallTargetsService) CallTargetsInfo(in model.CallTargetsModel) (out param.CallTargetsInfo) {
	out.ID = in.ID
	out.CallID = in.CallID
	out.UserID = in.UserID

	return out
}

func (s *sCallTargetsService) CallTargetsList(in []model.CallTargetsModel) (out []param.CallTargetsInfo) {
	for _, v := range in {
		info := s.CallTargetsInfo(v)
		out = append(out, info)
	}

	return out
}
