package search

import "gorm.io/gorm"

type sScaleSearch struct{}

var ScaleSearch = sScaleSearch{}

func (s *sScaleSearch) SearchScale(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Where("`flag_show` = ?", 1)
		return db
	}
}
