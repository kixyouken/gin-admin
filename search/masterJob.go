package search

import "gorm.io/gorm"

type sJobSearch struct{}

var JobSearch = sJobSearch{}

func (s *sJobSearch) SearchJob(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Where("`flag_show` = ?", 1)
		return db
	}
}
