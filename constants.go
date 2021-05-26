package paypayopa

// CodeType is the code type required for QR Code creation requests.
//
// CodeType は QR コード作成リクエストに必要なコード種別.
type CodeType string

const (
	// CodeTypeOrderQR is a fixed value that needs to be entered
	// when sending a QR Code creation request.
	//
	// CodeTypeOrderQR は QR コード作成リクエストを送信するときに入力が必要な固定値.
	CodeTypeOrderQR CodeType = "ORDER_QR"

	// CodeTypeTopupQR is a fixed value that needs to be entered
	// when sending a request to create a QR code for top-up.
	//
	// CodeTypeTopupQR はトップアップ用の QR コード作成リクエストを
	// 送信するときに入力が必要な固定値.
	CodeTypeTopupQR CodeType = "TOPUP_QR"
)

// Currency is the type of currency.
//
// Currency は通貨の種別.
type Currency string

// CurrencyJPY is the currency unit for the Japanese yen.
//
// CurrencyJPY は日本円を表す通貨単位.
const CurrencyJPY Currency = "JPY"

// Lang is a value specified in the lang header to set
// the language of the cashback message text.
//
// Lang はキャッシュバックメッセージテキストの言語を設定するために
// lang ヘッダに指定する値.
type Lang string

const (
	// LangEN is a value to set the language of
	// the cashback message text to English.
	//
	// LangEN はキャッシュバックメッセージテキストの言語の設定を英語にするための値.
	LangEN = "EN"

	// LangJA is a value to set the language of
	// the cashback message text to Japanese.
	// It is the default value and need not be specified.
	//
	// LangJA はキャッシュバックメッセージテキストの言語の設定を日本語にするための値.
	// デフォルト値の為, 指定する必要は無い.
	LangJA = "JA"

	// headerNameLang is the name of the language header
	// to set to specify the language of the response.
	//
	// headerNameLang はレスポンスの言語を指定する為にセットする言語ヘッダの名前.
	headerNameLang = "lang"
)

// RedirectType is the type of redirection to specify
// when sending a QR Code creation request.
//
// RedirectType は QR コード作成リクエストを
// 送信するときに指定するリダイレクトの種別.
type RedirectType string

const (
	// RedirectTypeWebLink is specified when the payment is occurring
	// in a web browser.
	//
	// RedirectTypeWebLink は支払いがウェブブラウザで発生しているときに指定する.
	RedirectTypeWebLink RedirectType = "WEB_LINK"

	// RedirectTypeDeepLink is specified when the payment is occurring
	// in the app.
	//
	// RedirectTypeDeepLink は支払いがアプリで発生しているときに指定する.
	RedirectTypeDeepLink RedirectType = "APP_DEEP_LINK"
)

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
