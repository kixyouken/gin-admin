package search

import "gorm.io/gorm"

type sDeliveryLogsSearch struct{}

var DeliveryLogsSearch = sDeliveryLogsSearch{}

func (s *sDeliveryLogsSearch) SearchDeliveryLogs(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
