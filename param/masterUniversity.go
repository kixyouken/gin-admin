package param

type MasterUniversityInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	NameKana    string `json:"name_kana"`
	Area        string `json:"area"`
	Prefectures string `json:"prefectures"`
	Class       int    `json:"class"`
	Type        string `json:"type"`
	Rank        string `json:"rank"`
	FlagShow    int    `json:"flag_show"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
