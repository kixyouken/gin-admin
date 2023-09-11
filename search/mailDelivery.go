package search

import "gorm.io/gorm"

type sDeliverySearch struct{}

var DeliverySearch = sDeliverySearch{}

func (s *sDeliverySearch) SearchDelivery(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Unscoped()
		return db
	}
}
