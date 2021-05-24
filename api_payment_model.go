package paypayopa

import "encoding/json"

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

type CapturePaymentAuthorizationPayload struct {
	MerchantPaymentID string       `json:"merchantPaymentId"`
	Amount            *MoneyAmount `json:"amount"`
	MerchantCaptureID string       `json:"merchantCaptureId"`
	RequestedAt       int64        `json:"requestedAt"`
	OrderDescription  string       `json:"orderDescription"`
}

type ConsultExpectedCashbackInfoPayload struct {
	RequestID           string               `json:"requestId"`
	MerchantPaymentID   string               `json:"merchantPaymentId"`
	UserAuthorizationID string               `json:"userAuthorizationId"`
	Amount              *MoneyAmount         `json:"amount"`
	RequestedAt         int64                `json:"requestedAt"`
	StoreID             string               `json:"storeId"`
	OrderItems          []*MerchantOrderItem `json:"orderItems"`
	ProductType         string               `json:"productType"`

	Lang Lang `json:"-"`
}

type CreateContinuousPaymentPayload struct {
	MerchantPaymentID   string               `json:"merchantPaymentId"`
	UserAuthorizationID string               `json:"userAuthorizationId"`
	Amount              *MoneyAmount         `json:"amount"`
	RequestedAt         int64                `json:"requestedAt"`
	StoreID             string               `json:"storeId"`
	TerminalID          string               `json:"terminalId"`
	OrderReceiptNumber  string               `json:"orderReceiptNumber"`
	OrderDescription    string               `json:"orderDescription"`
	OrderItems          []*MerchantOrderItem `json:"orderItems"`
	Metadata            *json.RawMessage     `json:"metadata"`
}

type CreatePaymentPayload struct {
	MerchantPaymentID   string               `json:"merchantPaymentId"`
	UserAuthorizationID string               `json:"userAuthorizationId"`
	Amount              *MoneyAmount         `json:"amount"`
	RequestedAt         int64                `json:"requestedAt"`
	StoreID             string               `json:"storeId"`
	TerminalID          string               `json:"terminalId"`
	OrderReceiptNumber  string               `json:"orderReceiptNumber"`
	OrderDescription    string               `json:"orderDescription"`
	OrderItems          []*MerchantOrderItem `json:"orderItems"`
	Metadata            *json.RawMessage     `json:"metadata"`
	ProductType         string               `json:"productType"`

	AgreeSimilarTransaction bool `json:"-"`
}

type CreatePaymentAuthorizationPayload struct {
	MerchantPaymentID   string               `json:"merchantPaymentId"`
	UserAuthorizationID string               `json:"userAuthorizationId"`
	Amount              *MoneyAmount         `json:"amount"`
	RequestedAt         int64                `json:"requestedAt"`
	ExpiresAt           int64                `json:"expiresAt"`
	StoreID             string               `json:"storeId"`
	TerminalID          string               `json:"terminalId"`
	OrderReceiptNumber  string               `json:"orderReceiptNumber"`
	OrderDescription    string               `json:"orderDescription"`
	OrderItems          []*MerchantOrderItem `json:"orderItems"`
	Metadata            *json.RawMessage     `json:"metadata"`

	AgreeSimilarTransaction bool `json:"-"`
}

type CreatePendingPaymentPayload struct {
	MerchantPaymentID   string               `json:"merchantPaymentId"`
	UserAuthorizationID string               `json:"userAuthorizationId"`
	Amount              *MoneyAmount         `json:"amount"`
	RequestedAt         int64                `json:"requestedAt"`
	ExpiryDate          *int64               `json:"expiryDate"`
	StoreID             string               `json:"storeId"`
	TerminalID          string               `json:"terminalId"`
	OrderReceiptNumber  string               `json:"orderReceiptNumber"`
	OrderDescription    string               `json:"orderDescription"`
	OrderItems          []*MerchantOrderItem `json:"orderItems"`
	Metadata            *json.RawMessage     `json:"metadata"`
}

type CreateQrCodePayload struct {
	MerchantPaymentID   string               `json:"merchantPaymentId"`
	Amount              *MoneyAmount         `json:"amount"`
	OrderDescription    string               `json:"orderDescription"`
	OrderItems          []*MerchantOrderItem `json:"orderItems"`
	Metadata            interface{}          `json:"metadata"`
	CodeType            CodeType             `json:"codeType"`
	StoreInfo           string               `json:"storeInfo"`
	StoreID             string               `json:"storeId"`
	TerminalID          string               `json:"terminalId"`
	RequestedAt         int64                `json:"requestedAt"`
	RedirectURL         string               `json:"redirectUrl"`
	RedirectType        RedirectType         `json:"redirectType"`
	UserAgent           string               `json:"userAgent"`
	IsAuthorization     bool                 `json:"isAuthorization"`
	AuthorizationExpiry *int64               `json:"authorizationExpiry"`
}

type RefundPaymentPayload struct {
	MerchantRefundID string       `json:"merchantRefundId"`
	PaymentID        string       `json:"paymentId"`
	Amount           *MoneyAmount `json:"amount"`
	RequestedAt      int64        `json:"requestedAt"`
	Reason           string       `json:"reason"`
}

type RevertPaymentAuthorizationPayload struct {
	MerchantRevertID string `json:"merchantRevertId"`
	PaymentID        string `json:"paymentId"`
	RequestedAt      int64  `json:"requestedAt"`
	Reason           string `json:"reason"`
}

type CashbackInfoResponse struct {
	Campaignmessage string `json:"campaignMessage"`
}

type MerchantOrderItemResponse struct {
	Name      string       `json:"name"`
	Category  string       `json:"category"`
	Quantity  int          `json:"quantity"`
	ProductID string       `json:"productId"`
	UnitPrice *MoneyAmount `json:"unit_price"`
}

type Payment struct {
	PaymentID  string `json:"paymentId"`
	Status     string `json:"status"`
	AcceptedAt int64  `json:"acceptedAt"`
	Refunds    struct {
		Data []*Refund `json:"data"`
	} `json:"refunds"`
	Captures struct {
		Data []*Capture `json:"data"`
	} `json:"captures"`
	Revert             *Revert              `json:"revert"`
	MerchantPaymentID  string               `json:"merchantPaymentId"`
	Amount             *MoneyAmount         `json:"amount"`
	RequestedAt        int64                `json:"requestedAt"`
	ExpiresAt          *int64               `json:"expiresAt"`
	CanceledAt         *int64               `json:"canceledAt"`
	StoreID            string               `json:"storeId"`
	TerminalID         string               `json:"terminalId"`
	OrderReceiptNumber string               `json:"orderReceiptNumber"`
	OrderDescription   string               `json:"orderDescription"`
	OrderItems         []*MerchantOrderItem `json:"orderItems"`
	Metadata           *json.RawMessage     `json:"metadata"`
}

type PendingPayment struct {
	MerchantPaymentID   string               `json:"merchantPaymentId"`
	UserAuthorizationID string               `json:"userAuthorizationId"`
	Amount              *MoneyAmount         `json:"amount"`
	RequestedAt         int64                `json:"requestedAt"`
	ExpiryDate          *int64               `json:"expiryDate"`
	StoreID             string               `json:"storeId"`
	TerminalID          string               `json:"terminalId"`
	OrderReceiptNumber  string               `json:"orderReceiptNumber"`
	OrderDescription    string               `json:"orderDescription"`
	OrderItems          []*MerchantOrderItem `json:"orderItems"`
	Metadata            *json.RawMessage     `json:"metadata"`
}

type QRCodeResponse struct {
	CodeID              string                       `json:"codeId"`
	URL                 string                       `json:"url"`
	DeepLink            string                       `json:"deeplink"`
	ExpiryDate          int64                        `json:"expiryDate"`
	MerchantPaymentID   string                       `json:"merchantPaymentId"`
	Amount              *MoneyAmount                 `json:"amount"`
	OrderDescription    string                       `json:"orderDescription"`
	OrderItems          []*MerchantOrderItemResponse `json:"orderItems"`
	Metadata            *json.RawMessage             `json:"metadata"`
	CodeType            string                       `json:"codeType"`
	StoreInfo           string                       `json:"storeInfo"`
	StoreID             string                       `json:"storeId"`
	TerminalID          string                       `json:"terminalId"`
	RequestedAt         int64                        `json:"requestedAt"`
	RedirectURL         string                       `json:"redirectUrl"`
	RedirectType        RedirectType                 `json:"redirectType"`
	Isauthorization     bool                         `json:"isAuthorization"`
	Authorizationexpiry *int64                       `json:"authorizationExpiry"`
}

type Refund struct {
	Status           string       `json:"status"`
	AcceptedAt       int64        `json:"acceptedAt"`
	MerchantRefundID string       `json:"merchantRefundId"`
	PaymentID        string       `json:"paymentId"`
	Amount           *MoneyAmount `json:"amount"`
	RequestedAt      int64        `json:"requestedAt"`
	Reason           string       `json:"reason"`
	AssumeMerchant   string       `json:"assumeMerchant"`
}

type RevertedPayment struct {
	Status      string `json:"status"`
	AcceptedAt  int64  `json:"acceptedAt"`
	PaymentID   string `json:"paymentId"`
	RequestedAt int64  `json:"requestedAt"`
	Reason      string `json:"reason"`
}
