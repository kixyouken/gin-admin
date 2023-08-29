package model

import "gorm.io/gorm"

type MenuModel struct {
	gorm.Model
	ParentID  uint
	Order     int
	Title     string
	Icon      string
	Uri       string
	Extension string
	Show      int
}

func (MenuModel) TableName() string {
	return "admin_menu"
}
