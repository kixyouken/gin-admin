package search

import (
	"encoding/json"
	"gin-admin/param"

	"gorm.io/gorm"
)

type sLocationSearch struct{}

var LocationSearch = sLocationSearch{}

func (s *sLocationSearch) SearchLocation(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		body, _ := json.Marshal(search)
		location := param.MasterLocationSearch{}
		json.Unmarshal(body, &location)
		if location.ID != nil {
			db.Where("`id` = ?", location.ID)
		}
		if location.ChildID != nil {
			db.Where("`id` BETWEEN ? AND ?", location.ChildID[0], location.ChildID[1])
		}
		db.Where("`flag_show` = ?", 1)
		return db
	}
}
