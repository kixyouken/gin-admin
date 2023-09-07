package service

import (
	"gin-admin/common/format"
	"gin-admin/model"
	"gin-admin/param"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sUsersService struct{}

var UsersService = sUsersService{}

var (
	modelUsersList []model.UsersModel
)

func (s *sUsersService) UsersInfo(in model.UsersModel) (out param.UsersInfo) {
	out.ID = in.ID
	out.Email = in.Email
	out.Password = in.Password
	out.FamilyName = in.FamilyName
	out.Name = in.Name
	out.KanaFamilyName = in.KanaFamilyName
	out.KanaName = in.KanaName
	out.Birthday = in.Birthday.Format(format.YMD)
	out.UniversityNameSort = in.UniversityNameSort
	out.MasterUniversityID = in.MasterUniversityID
	out.UniversityNameOpt = in.UniversityNameOpt
	out.Faculty = in.Faculty
	out.Department = in.Department
	out.Type = in.Type
	out.GraduationYear = in.GraduationYear
	out.GraduationMonth = in.GraduationMonth
	out.Mobile = in.Mobile
	out.PostalCode = in.PostalCode
	out.AddressArea = in.AddressArea
	out.AddressCity = in.AddressCity
	out.Address = in.Address
	out.AddressBuilding = in.AddressBuilding
	out.MasterOrganizationID = in.MasterOrganizationID
	out.Agreement = in.Agreement
	out.Status = in.Status
	out.ProspectiveDestination = in.ProspectiveDestination
	out.SelfPr = in.SelfPr
	out.Extracurricular = in.Extracurricular
	out.MasterNegotiateID = in.MasterNegotiateID
	out.MasterScaleID = in.MasterScaleID
	out.MasterIndustryID = in.MasterIndustryID
	out.MasterJobID = in.MasterJobID
	out.MasterWelfareID = in.MasterWelfareID
	out.MasterLocationID = in.MasterLocationID
	out.MasterScienceID = in.MasterScienceID
	out.Unsubscribe = in.Unsubscribe
	out.FlagWithdrawal = in.FlagWithdrawal
	if in.WithdrawaledAt != nil {
		out.WithdrawaledAt = in.WithdrawaledAt.Format(format.YMDHI)
	}
	out.FlagTelng = in.FlagTelng
	out.ApInformalOffer = in.ApInformalOffer
	out.ApComment = in.ApComment
	out.ApRemark = in.ApRemark
	out.EmailVerifiedAt = in.EmailVerifiedAt.Format(format.YMDHI)
	out.RememberToken = in.RememberToken
	out.LastLoginAt = in.LastLoginAt.Format(format.YMDHI)
	out.OrganizationNameOpt = in.OrganizationNameOpt
	out.WithdrawalBy = in.WithdrawalBy
	out.CreatedAt = in.CreatedAt.Format(format.YMDHI)
	out.UpdatedAt = in.UpdatedAt.Format(format.YMDHI)
	if !in.DeletedAt.Time.IsZero() {
		out.DeletedAt = in.DeletedAt.Time.Format(format.YMDHI)
	}

	out.MasterUniversity = MasterUniversityService.MasterUniversityInfo(in.MasterUniversity)
	out.MasterOrganization = MasterOrganizationService.OrganizationInfo(in.MasterOrganization)
	out.MasterScience = MasterScienceService.ScienceInfo(in.MasterScience)
	out.MasterIndustry = MasterIndustryService.IndustryInfo(in.MasterIndustry)
	out.MasterJob = MasterJobService.JobInfo(in.MasterJob)
	out.MasterWelfare = MasterWelfareService.WelfareInfo(in.MasterWelfare)
	out.MasterLocation = MasterLocationService.LocationInfo(in.MasterLocation)
	out.MasterScale = MasterScaleService.ScaleInfo(in.MasterScale)

	return out
}

func (s *sUsersService) UsersList(in []model.UsersModel) (out []param.UsersInfo) {
	for _, v := range in {
		info := s.UsersInfo(v)
		out = append(out, info)
	}

	return out
}

func (s *sUsersService) GetUsersString(in param.UsersInfo) param.UsersInfo {
	switch in.Type {
	case 1:
		in.TypeText = "文科"
	case 2:
		in.TypeText = "理科"
	}

	switch in.Status {
	case 0:
		in.StatusText = "未入力"
	case 1:
		in.StatusText = "内定"
	case 2:
		in.StatusText = "重複"
	case 3:
		in.StatusText = "公務員"
	case 4:
		in.StatusText = "教職"
	case 5:
		in.StatusText = "既卒"
	case 6:
		in.StatusText = "卒年延長"
	case 7:
		in.StatusText = "留学"
	case 8:
		in.StatusText = "進学"
	case 9:
		in.StatusText = "ダミー登録"
	case 91:
		in.StatusText = "メールエラー"
	}

	if in.ApInformalOffer == "1" {
		in.ApInformalOfferText = "ON"
	}

	if in.Unsubscribe == 1 {
		in.UnsubscribeText = "ON"
	}

	if in.FlagTelng == 1 {
		in.FlagTelngText = "ON"
	}

	if in.FlagWithdrawal == 1 {
		in.FlagWithdrawalText = "ON"
	}

	return in
}

// GetAllUser 分页查询
//
//	@receiver s
//	@param c
//	@param search
//	@param order
//	@return []model.UsersModel
//	@return error
func (s *sUsersService) GetAllUsers(c *gin.Context, search func(db *gorm.DB) *gorm.DB, order string) ([]model.UsersModel, error) {
	err := BaseService.GetPageJoins(c, &modelUsersList, search, "`users`.*, `MasterUniversity`.`name` AS `master_university_name`", "`users`.id DESC", "MasterUniversity")
	if err != nil {
		return nil, err
	}

	return modelUsersList, nil
}

// GetCountUser 查询数量
//
//	@receiver s
//	@param c
//	@param search
//	@return int64
func (s *sUsersService) GetCountUsers(c *gin.Context, search func(db *gorm.DB) *gorm.DB) int64 {
	return BaseService.GetCountJoins(&modelUsersList, search, "MasterUniversity")
}

// GetByIDUsers 根据ID获取
//
//	@receiver s
//	@param c
//	@param id
//	@return *model.UsersModel
//	@return error
func (s *sUsersService) GetByIDUsers(c *gin.Context, id uint) (*model.UsersModel, error) {
	modelUsersInfo := model.UsersModel{}
	err := BaseService.GetByIDUnscoped(&modelUsersInfo, id, "MasterUniversity", "MasterOrganization", "MasterScience", "MasterIndustry", "MasterJob", "MasterWelfare", "MasterLocation", "MasterScale")
	if err != nil {
		return nil, err
	}

	return &modelUsersInfo, nil
}

// SetByIDUsers 根据ID更新
//
//	@receiver s
//	@param c
//	@param id
//	@param updates
//	@param column
//	@return *model.UsersModel
//	@return error
func (s *sUsersService) SetByIDUsers(c *gin.Context, id uint, updates interface{}, column interface{}) (*model.UsersModel, error) {
	modelUsersInfo := model.UsersModel{}
	err := BaseService.UpdateByID(&modelUsersInfo, id, updates, column)
	if err != nil {
		return nil, err
	}

	return &modelUsersInfo, nil
}
