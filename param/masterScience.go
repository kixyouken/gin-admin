package param

type MasterScienceInfo struct {
	ID            uint   `json:"id"`
	ResearchAreas string `json:"research_areas"`
	FlagShow      int    `json:"flag_show"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}

type MasterScienceSearch struct {
}
