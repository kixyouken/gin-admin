package param

// UsersInfo 详情
type UsersInfo struct {
	ID                     uint                   `json:"id"`
	Email                  string                 `json:"email" form:"email"`
	Password               string                 `json:"password"`
	FamilyName             string                 `json:"family_name" form:"family_name"`
	Name                   string                 `json:"name" form:"name"`
	KanaFamilyName         string                 `json:"kana_family_name" form:"kana_family_name"`
	KanaName               string                 `json:"kana_name" form:"kana_name"`
	Birthday               string                 `json:"birthday" form:"birthday"`
	UniversityNameSort     int                    `json:"university_name_sort"`
	MasterUniversityID     uint                   `json:"master_university_id" form:"master_university_id"`
	UniversityNameOpt      string                 `json:"university_name_opt" form:"university_name_opt"`
	Faculty                string                 `json:"faculty" form:"faculty"`
	Department             string                 `json:"department" form:"department"`
	Type                   int                    `json:"type" form:"type"`
	TypeText               string                 `json:"type_text"`
	GraduationYear         int                    `json:"graduation_year" form:"graduation_year"`
	GraduationMonth        int                    `json:"graduation_month" form:"graduation_month"`
	Mobile                 string                 `json:"mobile" form:"mobile"`
	PostalCode             string                 `json:"postal_code" form:"postal_code"`
	AddressArea            string                 `json:"address_area" form:"address_area"`
	AddressCity            string                 `json:"address_city" form:"address_city"`
	Address                string                 `json:"address" form:"address"`
	AddressBuilding        string                 `json:"address_building" form:"address_building"`
	MasterOrganizationID   uint                   `json:"master_organization_id" form:"master_organization_id"`
	Agreement              int                    `json:"agreement"`
	Status                 int                    `json:"status" form:"status"`
	StatusText             string                 `json:"status_text"`
	ProspectiveDestination string                 `json:"prospective_destination" form:"prospective_destination"`
	SelfPr                 string                 `json:"self_pr" form:"self_pr"`
	Extracurricular        string                 `json:"extracurricular" form:"extracurricular"`
	MasterNegotiateID      string                 `json:"master_negotiate_id" form:"master_negotiate_id"`
	MasterScaleID          uint                   `json:"master_scale_id" form:"master_scale_id"`
	MasterIndustryID       uint                   `json:"master_industry_id" form:"master_industry_id"`
	MasterJobID            uint                   `json:"master_job_id" form:"master_job_id"`
	MasterWelfareID        uint                   `json:"master_welfare_id" form:"master_welfare_id"`
	MasterLocationID       uint                   `json:"master_location_id" form:"master_location_id"`
	MasterScienceID        uint                   `json:"master_science_id" form:"master_science_id"`
	Unsubscribe            int                    `json:"unsubscribe" form:"unsubscribe"`
	UnsubscribeText        string                 `json:"unsubscribe_text"`
	FlagWithdrawal         int                    `json:"flag_withdrawal" form:"flag_withdrawal"`
	FlagWithdrawalText     string                 `json:"flag_withdrawal_text"`
	WithdrawaledAt         string                 `json:"withdrawaled_at"`
	FlagTelng              int                    `json:"flag_telng" form:"flag_telng"`
	FlagTelngText          string                 `json:"flag_telng_text"`
	ApInformalOffer        string                 `json:"ap_informal_offer" form:"ap_informal_offer"`
	ApInformalOfferText    string                 `json:"ap_informal_offer_text"`
	ApComment              string                 `json:"ap_comment" form:"ap_comment"`
	ApRemark               string                 `json:"ap_remark" form:"ap_remark"`
	EmailVerifiedAt        string                 `json:"email_verified_at"`
	RememberToken          string                 `json:"remember_token"`
	LastLoginAt            string                 `json:"last_login_at"`
	OrganizationNameOpt    string                 `json:"organization_name_opt" form:"organization_name_opt"`
	WithdrawalBy           int                    `json:"withdrawal_by"`
	CreatedAt              string                 `json:"created_at"`
	UpdatedAt              string                 `json:"updated_at"`
	DeletedAt              string                 `json:"deleted_at" form:"deleted_at"`
	MasterUniversity       MasterUniversityInfo   `json:"master_university"`
	MasterOrganization     MasterOrganizationInfo `json:"master_organization"`
	MasterScience          MasterScienceInfo      `json:"master_science"`
	MasterIndustry         MasterIndustryInfo     `json:"master_industry"`
	MasterJob              MasterJobInfo          `json:"master_job"`
	MasterWelfare          MasterWelfareInfo      `json:"master_welfare"`
	MasterLocation         MasterLocationInfo     `json:"master_location"`
	MasterScale            MasterScaleInfo        `json:"master_scale"`
}

// UsersSearch 搜索
type UsersSearch struct {
	ID                     uint     `form:"id"`
	Name                   string   `form:"name"`
	KanaName               string   `form:"kana_name"`
	UniversityName         string   `form:"university_name"`
	Mobile                 string   `form:"mobile"`
	Email                  string   `form:"email"`
	Faculty                string   `form:"faculty"`
	Type                   int      `form:"type"`
	UniversityArea         []string `form:"university_area[]"`
	Rank                   []string `form:"rank[]"`
	Graduation             string   `form:"graduation"`
	ProspectiveDestination string   `form:"prospective_destination"`
	Status                 []string `form:"status[]"`
	CreatedAt              string   `form:"created_at"`
	Unsubscribe            string   `form:"unsubscribe"`
	FlagTelng              string   `form:"flag_telng"`
	FlagWithdrawal         string   `form:"flag_withdrawal"`
	ApInformalOffer        string   `form:"ap_informal_offer"`
	CallUserID             string   `form:"call_user_id"`
}
