package paypayopa

import (
	"context"
	"encoding/json"
	"time"
)

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

func createPendingPayment(
	ctx context.Context,
	client *opaClient,
	req *CreatePendingPaymentPayload,
) (*PendingPayment, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &PendingPayment{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v1/requestOrder",
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func cancelPendingOrder(
	ctx context.Context,
	client *opaClient,
	merchantPaymentID string,
) (*ResultInfo, error) {
	const timeout = 15 * time.Second

	return client.DELETE(
		ctxWithTimeout(ctx, timeout),
		"/v1/requestOrder/"+merchantPaymentID,
	)
}

func getRequestedPaymentDetails(
	ctx context.Context,
	client *opaClient,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &Payment{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v1/requestOrder/"+merchantPaymentID,
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}
