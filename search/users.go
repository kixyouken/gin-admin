package search

import (
	"encoding/json"
	"gin-admin/param"
	"strings"

	"gorm.io/gorm"
)

type sUsersSearch struct{}

var UsersSearch = sUsersSearch{}

func (s *sUsersSearch) SearchUsers(search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		body, _ := json.Marshal(search)
		users := param.UsersSearch{}
		json.Unmarshal(body, &users)
		if users.ID > 0 {
			db.Where("`users`.`id` = ?", users.ID)
		}
		if users.Name != "" {
			db.Where("`name` LIKE ? OR `family_name` LIKE ?", "%"+users.Name+"%", "%"+users.Name+"%")
		}
		if users.KanaName != "" {
			db.Where("`kana_name` LIKE ? OR `kana_family_name` LIKE ?", "%"+users.KanaName+"%", "%"+users.KanaName+"%")
		}
		if users.UniversityName != "" {
			db.Where("`MasterUniversity`.`name` LIKE ? OR `university_name_opt` LIKE ?", "%"+users.UniversityName+"%", "%"+users.UniversityName+"%")
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
		if users.UniversityArea != nil {
			db.Where("`MasterUniversity`.`area` IN ?", users.UniversityArea)
		}
		if users.Rank != nil {
			db.Where("`MasterUniversity`.`rank` IN ?", users.Rank)
		}
		if users.Graduation != "" {
			graduation := strings.Split(users.Graduation, " - ")
			db.Where("CONCAT(`graduation_year`, '-', LPAD(`graduation_month`, 2, '0')) BETWEEN ? AND ?", graduation[0], graduation[1])
		}
		if users.ProspectiveDestination != "" {
			db.Where("`prospective_destination` LIKE ?", "%"+users.ProspectiveDestination+"%")
		}
		if users.Status != nil {
			db.Where("`status` IN ?", users.Status)
		}
		if users.CreatedAt != "" {
			createAt := strings.Split(users.CreatedAt, " - ")
			db.Where("`users`.`created_at` BETWEEN ? AND ?", createAt[0], createAt[1]+" 23:59:59")
		}
		if users.Unsubscribe != "" {
			db.Where("`unsubscribe` NOT IN (?) OR `unsubscribe` IS NULL", users.Unsubscribe)
		}
		if users.FlagTelng != "" {
			db.Where("`flag_telng` NOT IN (?) OR `flag_telng` IS NULL", users.FlagTelng)
		}
		if users.FlagWithdrawal != "" {
			db.Where("`flag_withdrawal` NOT IN (?) OR `flag_withdrawal` IS NULL", users.FlagWithdrawal)
		}
		if users.ApInformalOffer != "" {
			db.Where("`ap_informal_offer` NOT IN (?) OR `ap_informal_offer` IS NULL", users.ApInformalOffer)
		}
		if users.CallUserID != "" {
			userID := strings.Split(users.CallUserID, ",")
			db.Where("`users`.`id` IN ?", userID)
		}
		db.Unscoped()
		return db
	}
}
