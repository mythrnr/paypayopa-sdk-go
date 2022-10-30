package paypayopa

import (
	"context"
	"encoding/json"
	"time"
)

type CreateQRCodePayload struct {
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
	IsAuthorization     bool                         `json:"isAuthorization"`
	Authorizationexpiry *int64                       `json:"authorizationExpiry"`
}

type MerchantOrderItemResponse struct {
	Name      string       `json:"name"`
	Category  string       `json:"category"`
	Quantity  int          `json:"quantity"`
	ProductID string       `json:"productId"`
	UnitPrice *MoneyAmount `json:"unit_price"`
}

func createQRCode(
	ctx context.Context,
	client *opaClient,
	req *CreateQRCodePayload,
) (*QRCodeResponse, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &QRCodeResponse{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v2/codes",
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func deleteQRCode(
	ctx context.Context,
	client *opaClient,
	codeID string,
) (*ResultInfo, error) {
	const timeout = 15 * time.Second

	return client.DELETE(
		ctxWithTimeout(ctx, timeout),
		"/v2/codes/"+codeID,
	)
}

func getCodePaymentDetails(
	ctx context.Context,
	client *opaClient,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &Payment{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v2/codes/payments/"+merchantPaymentID,
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}
