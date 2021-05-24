package paypayopa

import "encoding/json"

// Scope is the scope of the user authorization.
//
// Scope はユーザー認可のスコープ.
type Scope string

const (

	// ScopeDirectDebit represents the scope direct_debit
	// of the user authorization.
	//
	// ScopeDirectDebit はユーザー認可のスコープの direct_debit を表す.
	ScopeDirectDebit Scope = "direct_debit"

	// ScopeCashback represents the scope cashback
	// of the user authorization.
	//
	// ScopeCashback はユーザー認可のスコープの cashback を表す.
	ScopeCashback Scope = "cashback"

	// ScopeGetBalance represents the scope get_balance
	// of the user authorization.
	//
	// ScopeGetBalance はユーザー認可のスコープの get_balance を表す.
	ScopeGetBalance Scope = "get_balance"

	// ScopeQuickPay represents the scope quick_pay
	// of the user authorization.
	//
	// ScopeQuickPay はユーザー認可のスコープの quick_pay を表す.
	ScopeQuickPay Scope = "quick_pay"

	// ScopeContinuousPayments represents the scope
	// continuous_payments of the user authorization.
	//
	// ScopeContinuousPayments はユーザー認可のスコープの
	// continuous_payments を表す.
	ScopeContinuousPayments Scope = "continuous_payments"

	// ScopeMerchantTopup represents the scope merchant_topup
	// of the user authorization.
	//
	// ScopeMerchantTopup はユーザー認可のスコープの merchant_topup を表す.
	ScopeMerchantTopup Scope = "merchant_topup"

	// ScopePendingPayments represents the scope pending_payments
	// of the user authorization.
	//
	// ScopePendingPayments はユーザー認可のスコープの pending_payments を表す.
	ScopePendingPayments Scope = "pending_payments"

	// ScopeUserNotification represents the scope user_notification
	// of the user authorization.
	//
	// ScopeUserNotification はユーザー認可のスコープの user_notification を表す.
	ScopeUserNotification Scope = "user_notification"

	// ScopeUserTopup represents the scope user_topup
	// of the user authorization.
	//
	// ScopeUserTopup はユーザー認可のスコープの user_topup を表す.
	ScopeUserTopup Scope = "user_topup"

	// ScopeUserProfile represents the scope user_profile
	// of the user authorization.
	//
	// ScopeUserProfile はユーザー認可のスコープの user_profile を表す.
	ScopeUserProfile Scope = "user_profile"

	// ScopePreauthCaptureNative represents the scope
	// preauth_capture_native of the user authorization.
	//
	// ScopePreauthCaptureNative はユーザー認可のスコープの
	// preauth_capture_native を表す.
	ScopePreauthCaptureNative Scope = "preauth_capture_native"

	// ScopePreauthCaptureTransaction represents the scope
	// preauth_capture_transaction of the user authorization.
	//
	// ScopePreauthCaptureTransaction はユーザー認可のスコープの
	// preauth_capture_transaction を表す.
	ScopePreauthCaptureTransaction Scope = "preauth_capture_transaction"

	// ScopePushNotification represents the scope push_notification
	// of the user authorization.
	//
	// ScopePushNotification はユーザー認可のスコープの push_notification を表す.
	ScopePushNotification Scope = "push_notification"

	// ScopeNotificationCenterOg represents the scope
	// notification_center_og of the user authorization.
	//
	// ScopeNotificationCenterOg はユーザー認可のスコープの
	// notification_center_og を表す.
	ScopeNotificationCenterOg Scope = "notification_center_og"

	// ScopeNotificationCenterAb represents the scope
	// notification_center_ab of the user authorization.
	//
	// ScopeNotificationCenterAb はユーザー認可のスコープの
	// notification_center_ab を表す.
	ScopeNotificationCenterAb Scope = "notification_center_ab"

	// ScopeNotificationCenterTl represents the scope
	// notification_center_tl of the user authorization.
	//
	// ScopeNotificationCenterTl はユーザー認可のスコープの
	// notification_center_tl を表す.
	ScopeNotificationCenterTl Scope = "notification_center_tl"

	// ScopeBankRegistration represents the scope
	// bank_registration of the user authorization.
	//
	// ScopeBankRegistration はユーザー認可のスコープの
	// bank_registration を表す.
	ScopeBankRegistration Scope = "bank_registration"
)

type CreateAccountLinkQrCodePayload struct {
	Scopes       []Scope `json:"scopes"`
	Nonce        string  `json:"nonce"`
	RedirectType string  `json:"redirectType"`
	RedirectURL  string  `json:"redirectUrl"`
	ReferenceID  string  `json:"referenceId"`
	PhoneNumber  string  `json:"phoneNumber"`
	DeviceID     string  `json:"deviceId"`
	UserAgent    string  `json:"userAgent"`
}

type GetUserAuthorizationStatusResponse struct {
	UserAuthorizationID string           `json:"userAuthorizationId"`
	ReferenceIDs        *json.RawMessage `json:"referenceIds"`
	Status              string           `json:"status"`
	Scopes              []string         `json:"scopes"`
	ExpireAt            int64            `json:"expireAt"`
	IssuedAt            int64            `json:"issuedAt"`
}

type MaskedUserProfileResponse struct {
	PhoneNumber string `json:"phoneNumber"`
}

type CreateAccountLinkQrCodeResponse struct {
	LinkQRCodeURL string `json:"linkQRCodeURL"`
}
