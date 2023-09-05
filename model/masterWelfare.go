package model

import "gorm.io/gorm"

type MasterWelfareModel struct {
	gorm.Model
	Name     string
	FlagShow int
}

func (MasterWelfareModel) TableName() string {
	return "master_welfare"
}
