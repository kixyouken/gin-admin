package search

import "gorm.io/gorm"

type sScienceSearch struct{}

var ScienceSearch = sScienceSearch{}

func (s *sScienceSearch) SearchScience(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Where("`flag_show` = ?", 1)
		return db
	}
}
