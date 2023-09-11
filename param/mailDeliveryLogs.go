package param

type MailDeliveryLogsInfo struct {
	ID             uint   `json:"id"`
	MailDeliveryID uint   `json:"mail_delivery_id"`
	Title          string `json:"title"`
	SendAt         string `json:"send_at"`
	Sender         string `json:"sender"`
	Content        string `json:"content"`
	UserID         uint   `json:"user_id"`
	FlagOpened     int    `json:"flag_opened"`
	LinkClickTimes int    `json:"link_click_times"`
	Recycled       int    `json:"recycled"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type MailDeliveryLogsSearch struct {
}
