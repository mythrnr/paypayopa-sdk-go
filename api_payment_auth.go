package paypayopa

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
)

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

func createPaymentAuthorization(
	ctx context.Context,
	client *opaClient,
	req *CreatePaymentAuthorizationPayload,
) (*Payment, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &Payment{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v2/payments/preauthorize?agreeSimilarTransaction="+
			strconv.FormatBool(req.AgreeSimilarTransaction),
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

type CapturePaymentAuthorizationPayload struct {
	MerchantPaymentID string       `json:"merchantPaymentId"`
	Amount            *MoneyAmount `json:"amount"`
	MerchantCaptureID string       `json:"merchantCaptureId"`
	RequestedAt       int64        `json:"requestedAt"`
	OrderDescription  string       `json:"orderDescription"`
}

func capturePaymentAuthorization(
	ctx context.Context,
	client *opaClient,
	req *CapturePaymentAuthorizationPayload,
) (*Payment, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &Payment{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v2/payments/capture",
		res,
		req,
	)

	if err != nil ||
		!info.Success() ||
		info.Code == "USER_CONFIRMATION_REQUIRED" {
		return nil, info, err
	}

	return res, info, nil
}

type RevertPaymentAuthorizationPayload struct {
	MerchantRevertID string `json:"merchantRevertId"`
	PaymentID        string `json:"paymentId"`
	RequestedAt      int64  `json:"requestedAt"`
	Reason           string `json:"reason"`
}

type RevertedPaymentResponse struct {
	Status      string `json:"status"`
	AcceptedAt  int64  `json:"acceptedAt"`
	PaymentID   string `json:"paymentId"`
	RequestedAt int64  `json:"requestedAt"`
	Reason      string `json:"reason"`
}

func revertPaymentAuthorization(
	ctx context.Context,
	client *opaClient,
	req *RevertPaymentAuthorizationPayload,
) (*RevertedPaymentResponse, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &RevertedPaymentResponse{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v2/payments/preauthorize/revert",
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}
