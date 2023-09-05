package model

import "gorm.io/gorm"

type MasterOrganizationModel struct {
	gorm.Model
	Name     string
	FlagShow int
}

func (MasterOrganizationModel) TableName() string {
	return "master_organization"
}
