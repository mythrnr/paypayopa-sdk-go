package paypayopa

type WebhookCustomerAuthorizationSucceeded struct {
	NotificationType    string `json:"notification_type"`
	NotificationID      string `json:"notification_id"`
	CreatedAt           string `json:"createdAt"`
	ReferenceID         string `json:"referenceId"`
	Nonce               string `json:"nonce"`
	Scopes              string `json:"scopes"`
	UserAuthorizationID string `json:"userAuthorizationId"`
	ProfileIdentifier   string `json:"profileIdentifier"`
	Expiry              int    `json:"expiry"`
}

type WebhookCustomerAuthorizationFailed struct {
	NotificationType string `json:"notification_type"`
	NotificationID   string `json:"notification_id"`
	CreatedAt        string `json:"createdAt"`
	ReferenceID      string `json:"referenceId"`
	Nonce            string `json:"nonce"`
	Result           string `json:"result"`
	Reason           string `json:"reason"`
}

type WebhookCustomerAuthorizationRevoked struct {
	NotificationType    string `json:"notification_type"`
	NotificationID      string `json:"notification_id"`
	CreatedAt           string `json:"createdAt"`
	UserAuthorizationID string `json:"userAuthorizationId"`
	ReferenceID         string `json:"referenceId"`
}

type WebhookCustomerAuthorizationExtended struct {
	NotificationType    string `json:"notification_type"`
	NotificationID      string `json:"notification_id"`
	CreatedAt           string `json:"createdAt"`
	Scopes              string `json:"scopes"`
	UserAuthorizationID string `json:"userAuthorizationId"`
	Expiry              int64  `json:"expiry"`
}

type WebhookReconFile struct {
	NotificationType string `json:"notification_type"`
	NotificationID   string `json:"notification_id"`
	FileType         string `json:"fileType"`
	Path             string `json:"path"`
	RequestedAt      string `json:"requestedAt"`
}

type WebhookTransaction struct {
	NotificationType string `json:"notification_type"`
	MerchantID       string `json:"merchant_id"`
	StoreID          string `json:"store_id"`
	PosID            string `json:"pos_id"`
	OrderID          string `json:"order_id"`
	MerchantOrderID  string `json:"merchant_order_id"`
	AuthorizedAt     string `json:"authorized_at"`
	ExpiresAt        string `json:"expires_at"`
	PaidAt           string `json:"paid_at"`
	OrderAmount      int    `json:"order_amount"`
	State            string `json:"state"`
}
