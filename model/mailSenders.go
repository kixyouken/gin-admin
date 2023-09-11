package model

import "gorm.io/gorm"

type MailSendersModel struct {
	gorm.Model
	Name  string
	Email string
	Flag  int
}

func (MailSendersModel) TableName() string {
	return "mail_senders"
}
