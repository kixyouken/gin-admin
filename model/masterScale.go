package model

import "gorm.io/gorm"

type MasterScaleModel struct {
	gorm.Model
	Name     string
	FlagShow int
}

func (MasterScaleModel) TableName() string {
	return "master_scale"
}
