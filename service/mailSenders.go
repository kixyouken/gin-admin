package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"
)

type sMailSendersService struct{}

var MailSendersService = sMailSendersService{}

func (s *sMailSendersService) SendersInfo(in model.MailSendersModel) (out param.MailSendersInfo) {
	out.ID = in.ID
	out.Name = in.Name
	out.Email = in.Email
	out.Flag = in.Flag
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)
	if !in.DeletedAt.Time.IsZero() {
		out.DeletedAt = in.DeletedAt.Time.Format(format.YMDHI)
	}

	return out
}

func (s *sMailSendersService) SendersList(in []model.MailSendersModel) (out []param.MailSendersInfo) {
	for _, v := range in {
		info := s.SendersInfo(v)
		out = append(out, info)
	}

	return out
}
