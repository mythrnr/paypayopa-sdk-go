package paypayopa

import (
	"encoding/json"
	"net/http"
)

// response is the root object of the PayPay API response.
//
// response は PayPay API のレスポンスのルートオブジェクト.
type response struct {
	Result *ResultInfo     `json:"resultInfo"`
	Data   json.RawMessage `json:"data,omitempty"`
}

// ResultInfo is the result object of the PayPay API processing.
// The JSON of the response does not contain the HTTP status,
// but the SDK sets it in this object and returns it for convenience.
//
// ResultInfo は PayPay API の処理結果オブジェクト.
// レスポンスの JSON には HTTP ステータスは含まれないが,
// SDK は利便性の為にこのオブジェクトにセットして返す.
type ResultInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	CodeID  string `json:"codeId"`

	StatusCode int `json:"-"`
}

// Success returns the success or failure of the operation.
// Based on HTTP status, less than 400 is considered success.
// 3xx is not returned now.
//
// Success は処理の成否を返す.
// HTTP ステータスを基準に, 400 未満を成功と判定する.
// 3xx は現時点では返されない.
func (r *ResultInfo) Success() bool {
	return r.StatusCode < http.StatusBadRequest
}

type Capture struct {
	AcceptedAt        int64        `json:"acceptedAt"`
	MerchantCaptureID string       `json:"merchantCaptureId"`
	Amount            *MoneyAmount `json:"amount"`
	OrderDescription  string       `json:"orderDescription"`
	RequestedAt       int64        `json:"requestedAt"`
	ExpiresAt         *int64       `json:"expiresAt"`
	Status            string       `json:"status"`
}

type MerchantOrderItem struct {
	Name      string       `json:"name"`
	Category  string       `json:"category"`
	Quantity  int          `json:"quantity"`
	ProductID string       `json:"productId"`
	UnitPrice *MoneyAmount `json:"unitPrice"`
}

type MoneyAmount struct {
	Amount   int      `json:"amount"`
	Currency Currency `json:"currency"`
}

type Revert struct {
	AcceptedAt       int64  `json:"acceptedAt"`
	MerchantRevertID string `json:"merchantRevertId"`
	RequestedAt      int64  `json:"requestedAt"`
	Reason           string `json:"reason"`
}
