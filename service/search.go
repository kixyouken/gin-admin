package service

import (
	"encoding/json"
	"fmt"
	"gin-admin/param"

	"gorm.io/gorm"
)

type sSearchService struct{}

var SearchService = sSearchService{}

// SearchUser 搜索用户
//
//	@receiver s
//	@param search
//	@return db
//	@return func(db *gorm.DB) *gorm.DB
func (s *sSearchService) SearchUser(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		body, _ := json.Marshal(search)
		users := param.UsersInfo{}
		json.Unmarshal(body, &users)
		if users.ID > 0 {
			db.Where("id = ?", users.ID)
		}
		if users.Name != "" {
			fmt.Println(111)
			db.Where("`name` LIKE ? OR `family_name` LIKE ?", "%"+users.Name+"%", "%"+users.Name+"%")
		}
		if users.KanaName != "" {
			db.Where("`kana_name` LIKE ? OR `kana_family_name` LIKE ?", "%"+users.KanaName+"%", "%"+users.KanaName+"%")
		}
		if users.UniversityNameOpt != "" {
			db.Where("`MasterUniversity`.`name` LIKE ? OR `university_name_opt` LIKE ?", "%"+users.UniversityNameOpt+"%", "%"+users.UniversityNameOpt+"%")
		}
		if users.Mobile != "" {
			db.Where("`mobile` LIKE ?", "%"+users.Mobile+"%")
		}
		if users.Email != "" {
			db.Where("`email` LIKE ?", "%"+users.Email+"%")
		}
		if users.Faculty != "" {
			db.Where("`faculty` LIKE ? OR `department` LIKE ?", "%"+users.Faculty+"%", "%"+users.Faculty+"%")
		}
		if users.Type > 0 {
			db.Where("`users`.`type` = ?", users.Type)
		}
		return db
	}
}
