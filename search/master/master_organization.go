package master

import "gorm.io/gorm"

type sOrganizationSearch struct{}

var OrganizationSearch = sOrganizationSearch{}

func (s *sOrganizationSearch) SearchOrganization(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
