package model

import "gorm.io/gorm"

type MailDeliveryTargetsModel struct {
	gorm.Model
	UserID         uint
	MailDeliveryID uint
	DeletedAt      gorm.DeletedAt `gorm:"-"`
}

func (MailDeliveryTargetsModel) TableName() string {
	return "mail_delivery_targets"
}
