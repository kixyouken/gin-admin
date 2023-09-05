package model

import "gorm.io/gorm"

type MasterJobModel struct {
	gorm.Model
	Name     string
	FlagShow int
}

func (MasterJobModel) TableName() string {
	return "master_job"
}
