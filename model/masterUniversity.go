package model

import "gorm.io/gorm"

type MasterUniversityModel struct {
	gorm.Model
	Name        string
	NameKana    string
	Area        string
	Prefectures string
	Class       int
	Type        string
	Rank        string
	FlagShow    int
}

func (MasterUniversityModel) TableName() string {
	return "master_university"
}
