package search

import "gorm.io/gorm"

type sSignSearch struct{}

var SignSearch = sSignSearch{}

func (s *sSignSearch) SearchSign(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
