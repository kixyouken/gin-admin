package model

import "gorm.io/gorm"

type MasterLocationModel struct {
	gorm.Model
	Name     string
	FlagShow int
}

func (MasterLocationModel) TableName() string {
	return "master_location"
}
