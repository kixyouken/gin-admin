package model

import "gorm.io/gorm"

type MasterSignModel struct {
	gorm.Model
	Name     string
	Content  string
	FlagShow int
}

func (MasterSignModel) TableName() string {
	return "master_sign"
}
