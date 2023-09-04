package master

import "gorm.io/gorm"

type sUniversitySearch struct{}

var UniversitySearch = sUniversitySearch{}

func (s *sUniversitySearch) SearchUniversity(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
