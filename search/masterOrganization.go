package search

import "gorm.io/gorm"

type sOrganizationSearch struct{}

var OrganizationSearch = sOrganizationSearch{}

func (s *sOrganizationSearch) SearchOrganization(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Where("`flag_show` = ?", 1)
		return db
	}
}
