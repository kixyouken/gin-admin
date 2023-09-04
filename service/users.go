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
	modelUsersInfo model.UsersModel
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
	out.DeletedAt = in.DeletedAt.Time.Format(format.YMDHI)

	out.MasterUniversity = MasterUniversityService.MasterUniversityInfo(in.MasterUniversity)

	return out
}

func (s *sUsersService) UsersList(in []model.UsersModel) (out []param.UsersInfo) {
	for _, v := range in {
		info := s.UsersInfo(v)
		out = append(out, info)
	}

	return out
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
	err := BaseService.GetByID(&modelUsersInfo, id)
	if err != nil {
		return nil, err
	}

	return &modelUsersInfo, nil
}
