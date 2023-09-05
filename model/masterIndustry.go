package model

import "gorm.io/gorm"

type MasterIndustryModel struct {
	gorm.Model
	Name     string
	FlagShow int
}

func (MasterIndustryModel) TableName() string {
	return "master_industry"
}
