package model

import (
	"time"

	"gorm.io/gorm"
)

type UsersModel struct {
	gorm.Model
	Email                  string
	Password               string
	FamilyName             string
	Name                   string
	KanaFamilyName         string
	KanaName               string
	Birthday               time.Time
	UniversityNameSort     int
	MasterUniversityID     uint
	UniversityNameOpt      string
	Faculty                string
	Department             string
	Type                   int
	GraduationYear         int
	GraduationMonth        int
	Mobile                 string
	PostalCode             string
	AddressArea            string
	AddressCity            string
	Address                string
	AddressBuilding        string
	MasterOrganizationID   uint
	Agreement              int
	Status                 int
	ProspectiveDestination string
	SelfPr                 string
	Extracurricular        string
	MasterNegotiateID      string
	MasterScaleID          uint
	MasterIndustryID       uint
	MasterJobID            uint
	MasterWelfareID        uint
	MasterLocationID       uint
	MasterScienceID        uint
	Unsubscribe            int
	FlagWithdrawal         int
	WithdrawaledAt         *time.Time
	FlagTelng              int
	ApInformalOffer        string
	ApComment              string
	ApRemark               string
	EmailVerifiedAt        time.Time
	RememberToken          string
	LastLoginAt            time.Time
	OrganizationNameOpt    string
	WithdrawalBy           int
	MasterUniversityName   string
	MasterUniversity       MasterUniversityModel   `gorm:"foreignKey:ID;references:MasterUniversityID"`
	MasterOrganization     MasterOrganizationModel `gorm:"foreignKey:ID;references:MasterOrganizationID"`
	MasterScience          MasterScienceModel      `gorm:"foreignKey:ID;references:MasterScienceID"`
	MasterIndustry         MasterIndustryModel     `gorm:"foreignKey:ID;references:MasterIndustryID"`
	MasterJob              MasterJobModel          `gorm:"foreignKey:ID;references:MasterJobID"`
	MasterWelfare          MasterWelfareModel      `gorm:"foreignKey:ID;references:MasterWelfareID"`
	MasterLocation         MasterLocationModel     `gorm:"foreignKey:ID;references:MasterLocationID"`
	MasterScale            MasterScaleModel        `gorm:"foreignKey:ID;references:MasterScaleID"`
}

func (UsersModel) TableName() string {
	return "users"
}
