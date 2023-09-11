package param

type MailDeliveryTargetsInfo struct {
	ID             uint   `json:"id"`
	UserID         uint   `json:"user_id"`
	MailDeliveryID uint   `json:"mail_delivery_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type MailDeliveryTargetsSearch struct {
}
