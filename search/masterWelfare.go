package search

import "gorm.io/gorm"

type sWelfareSearch struct{}

var WelfareSearch = sWelfareSearch{}

func (s *sWelfareSearch) SearchWelfare(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Where("`flag_show` = ?", 1)
		return db
	}
}
