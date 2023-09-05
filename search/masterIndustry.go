package search

import "gorm.io/gorm"

type sIndustrySearch struct{}

var IndustrySearch = sIndustrySearch{}

func (s *sIndustrySearch) SearchIndustry(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Where("`flag_show` = ?", 1)
		return db
	}
}
