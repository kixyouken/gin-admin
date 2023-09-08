package param

type MailDeliveryInfo struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Title              string `json:"title"`
	SenderID           uint   `json:"sender_id"`
	SendAt             string `json:"send_at"`
	SendCompletedAt    string `json:"send_completed_at"`
	FlagImportant      int    `json:"flag_important"`
	FlagImmediate      int    `json:"flag_immediate"`
	Content            string `json:"content"`
	MasterSignID       uint   `json:"master_sign_id"`
	Status             int    `json:"status"`
	MailContentConfirm int    `json:"mail_content_confirm"`
	MailAddressConfirm int    `json:"mail_address_confirm"`
	SummaryCategory    int    `json:"summary_category"`
	SummaryTargetID    int    `json:"summary_target_id"`
	MailConfirmBy      int    `json:"mail_confirm_by"`
	MailConfirmAt      string `json:"mail_confirm_at"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	DeletedAt          string `json:"deleted_at"`
}

type MailDeliverySearch struct {
}
