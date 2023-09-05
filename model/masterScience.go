package model

import "gorm.io/gorm"

type MasterScienceModel struct {
	gorm.Model
	ResearchAreas string
	FlagShow      int
}

func (MasterScienceModel) TableName() string {
	return "master_science"
}
