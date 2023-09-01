package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"
)

type sMasterUniversityService struct{}

var MasterUniversityService = sMasterUniversityService{}

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

func (s *sMasterUniversityService) MasterUniversityList(in []model.MasterUniversityModel) (out []param.MasterUniversityInfo) {
	for _, v := range in {
		info := s.MasterUniversityInfo(v)
		out = append(out, info)
	}

	return out
}
