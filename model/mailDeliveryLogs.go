package model

import (
	"time"

	"gorm.io/gorm"
)

type MailDeliveryLogsModel struct {
	gorm.Model
	MailDeliveryID uint
	Title          string
	SendAt         time.Time
	Sender         string
	Content        string
	UserID         uint
	FlagOpened     int
	LinkClickTimes int
	Recycled       int
	DeletedAt      gorm.DeletedAt `gorm:"-"`
}

func (MailDeliveryLogsModel) TableName() string {
	return "mail_delivery_logs"
}
