package search

import (
	"encoding/json"
	"gin-admin/param"
	"strings"

	"gorm.io/gorm"
)

type sNegotiateSearch struct{}

var NegotiateSearch = sNegotiateSearch{}

func (s *sNegotiateSearch) SearchNegotiate(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		body, _ := json.Marshal(search)
		negotiate := param.MasterNegotiateSearch{}
		json.Unmarshal(body, &negotiate)
		if negotiate.ID != "" {
			ids := strings.Split(negotiate.ID, ",")
			db.Where("`id` IN ?", ids)
		}
		db.Where("`flag_show` = ?", 1)
		return db
	}
}
