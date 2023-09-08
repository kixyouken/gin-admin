package search

import (
	"encoding/json"
	"gin-admin/param"

	"gorm.io/gorm"
)

type sCallSearch struct{}

var CallSearch = sCallSearch{}

func (s *sCallSearch) SearchCall(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		body, _ := json.Marshal(search)
		call := param.CallSearch{}
		json.Unmarshal(body, &call)
		if call.Name != "" {
			db.Where("`name` LIKE ?", "%"+call.Name+"%")
		}
		db.Unscoped()
		return db
	}
}
