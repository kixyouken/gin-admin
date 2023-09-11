package model

import (
	"time"

	"gorm.io/gorm"
)

type MailDeliveryModel struct {
	gorm.Model
	Name                string
	Title               string
	SenderID            uint
	SendAt              time.Time
	SendCompletedAt     time.Time
	FlagImportant       int
	FlagImmediate       int
	Content             string
	MasterSignID        uint
	Status              int
	MailContentConfirm  int
	MailAddressConfirm  int
	SummaryCategory     int
	SummaryTargetID     int
	MailConfirmBy       int
	MailConfirmAt       time.Time
	MailDeliveryLogs    []MailDeliveryLogsModel    `gorm:"foreignKey:MailDeliveryID;references:ID"`
	MailDeliveryTargets []MailDeliveryTargetsModel `gorm:"foreignKey:MailDeliveryID;references:ID"`
	MailSenders         MailSendersModel           `gorm:"foreignKey:ID;references:SenderID"`
	MasterSign          MasterSignModel            `gorm:"foreignKey:ID;references:MasterSignID"`
}

func (MailDeliveryModel) TableName() string {
	return "mail_delivery"
}
