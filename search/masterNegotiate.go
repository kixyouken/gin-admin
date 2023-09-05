package search

import "gorm.io/gorm"

type sNegotiateSearch struct{}

var NegotiateSearch = sNegotiateSearch{}

func (s *sNegotiateSearch) SearchNegotiate(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Where("`flag_show` = ?", 1)
		return db
	}
}
