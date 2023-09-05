package param

type MasterNegotiateInfo struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	FlagShow  int    `json:"flag_show"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type MasterNegotiateSearch struct {
}
