package model

import "gorm.io/gorm"

type MasterNegotiateModel struct {
	gorm.Model
	Name     string
	FlagShow int
}

func (MasterNegotiateModel) TableName() string {
	return "master_negotiate"
}
