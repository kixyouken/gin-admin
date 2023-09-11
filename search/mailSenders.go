package search

import "gorm.io/gorm"

type sSendersSearch struct{}

var SendersSearch = sSendersSearch{}

func (s *sSendersSearch) SearchSenders(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
