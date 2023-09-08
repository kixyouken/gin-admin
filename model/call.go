package model

import "gorm.io/gorm"

type CallModel struct {
	gorm.Model
	Name        string
	Status      int
	Detail      string
	Remark      string
	CreatedBy   int
	UpdatedBy   int
	DeletedBy   int
	CallTargets []CallTargetsModel `gorm:"foreignKey:CallID;references:ID"`
}

func (CallModel) TableName() string {
	return "call"
}
