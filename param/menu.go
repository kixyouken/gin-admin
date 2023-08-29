package param

type MenuInfo struct {
	ID        uint        `json:"id"`
	ParentID  uint        `json:"parent_id"`
	Order     int         `json:"order"`
	Title     string      `json:"title"`
	Icon      string      `json:"icon"`
	Uri       string      `json:"uri"`
	Extension string      `json:"extension"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	Children  []*MenuInfo `json:"children"`
}
