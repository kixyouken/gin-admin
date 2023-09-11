package search

import "gorm.io/gorm"

type sDeliveryTargetsSearch struct{}

var DeliveryTargetsSearch = sDeliveryTargetsSearch{}

func (s *sDeliveryTargetsSearch) SearchDeliveryTargets(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
