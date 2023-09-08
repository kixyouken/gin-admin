package param

type CallInfo struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name" form:"name"`
	Status      int               `json:"status" form:"status"`
	Detail      string            `json:"detail" form:"detail"`
	Remark      string            `json:"remark" form:"remark"`
	CreatedBy   int               `json:"created_by"`
	UpdatedBy   int               `json:"updated_by"`
	DeletedBy   int               `json:"deleted_by"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	DeletedAt   string            `json:"deleted_at"`
	CallTargets []CallTargetsInfo `json:"call_targets"`
}

type CallSearch struct {
	Name string `form:"name"`
}
