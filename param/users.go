package param

type UsersInfo struct {
	ID                     uint                 `json:"id" form:"id"`
	Email                  string               `json:"email" form:"email"`
	Password               string               `json:"password"`
	FamilyName             string               `json:"family_name"`
	Name                   string               `json:"name" form:"name"`
	KanaFamilyName         string               `json:"kana_family_name"`
	KanaName               string               `json:"kana_name" form:"kana_name"`
	Birthday               string               `json:"birthday"`
	UniversityNameSort     int                  `json:"university_name_sort"`
	MasterUniversityID     uint                 `json:"master_university_id"`
	UniversityNameOpt      string               `json:"university_name_opt" form:"university_name_opt"`
	Faculty                string               `json:"faculty" form:"faculty"`
	Department             string               `json:"department"`
	Type                   int                  `json:"type" form:"type"`
	GraduationYear         int                  `json:"graduation_year"`
	GraduationMonth        int                  `json:"graduation_month"`
	Mobile                 string               `json:"mobile" form:"mobile"`
	PostalCode             string               `json:"postal_code"`
	AddressArea            string               `json:"address_area"`
	AddressCity            string               `json:"address_city"`
	Address                string               `json:"address"`
	AddressBuilding        string               `json:"address_building"`
	MasterOrganizationID   uint                 `json:"master_organization_id"`
	Agreement              int                  `json:"agreement"`
	Status                 int                  `json:"status"`
	ProspectiveDestination string               `json:"prospective_destination"`
	SelfPr                 string               `json:"self_pr"`
	Extracurricular        string               `json:"extracurricular"`
	MasterNegotiateID      string               `json:"master_negotiate_id"`
	MasterScaleID          uint                 `json:"master_scale_id"`
	MasterIndustryID       uint                 `json:"master_industry_id"`
	MasterJobID            uint                 `json:"master_job_id"`
	MasterWelfareID        uint                 `json:"master_welfare_id"`
	MasterLocationID       uint                 `json:"master_location_id"`
	MasterScienceID        uint                 `json:"master_science_id"`
	Unsubscribe            int                  `json:"unsubscribe"`
	FlagWithdrawal         int                  `json:"flag_withdrawal"`
	WithdrawaledAt         string               `json:"withdrawaled_at"`
	FlagTelng              int                  `json:"flag_telng"`
	ApInformalOffer        string               `json:"ap_informal_offer"`
	ApComment              string               `json:"ap_comment"`
	ApRemark               string               `json:"ap_remark"`
	EmailVerifiedAt        string               `json:"email_verified_at"`
	RememberToken          string               `json:"remember_token"`
	LastLoginAt            string               `json:"last_login_at"`
	OrganizationNameOpt    string               `json:"organization_name_opt"`
	WithdrawalBy           int                  `json:"withdrawal_by"`
	CreatedAt              string               `json:"created_at"`
	UpdatedAt              string               `json:"updated_at"`
	DeletedAt              string               `json:"deleted_at"`
	MasterUniversity       MasterUniversityInfo `json:"master_university"`
}
