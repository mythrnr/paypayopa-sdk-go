package paypayopa

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

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

func createPayment(
	ctx context.Context,
	client *opaClient,
	req *CreatePaymentPayload,
) (*Payment, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &Payment{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v2/payments?agreeSimilarTransaction="+
			strconv.FormatBool(req.AgreeSimilarTransaction),
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func cancelPayment(
	ctx context.Context,
	client *opaClient,
	merchantPaymentID string,
) (*ResultInfo, error) {
	const timeout = 15 * time.Second

	return client.DELETE(
		ctxWithTimeout(ctx, timeout),
		"/v2/payments/"+merchantPaymentID,
	)
}

func getPaymentDetails(
	ctx context.Context,
	client *opaClient,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &Payment{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v2/payments/"+merchantPaymentID,
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
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

func createContinuousPayment(
	ctx context.Context,
	client *opaClient,
	req *CreateContinuousPaymentPayload,
) (*Payment, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &Payment{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v1/subscription/payments",
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

type CashbackInfoResponse struct {
	Campaignmessage string `json:"campaignMessage"`
}

func consultExpectedCashbackInfo(
	ctx context.Context,
	client *opaClient,
	req *ConsultExpectedCashbackInfoPayload,
) (*CashbackInfoResponse, *ResultInfo, error) {
	const timeout = 15 * time.Second

	rq, err := client.Request(
		ctxWithTimeout(ctx, timeout),
		http.MethodPost,
		"/v1/payments/cashback/expected",
		req,
	)

	if err != nil {
		return nil, nil, err
	}

	if req.Lang != "" {
		rq.Header.Set(headerNameLang, string(req.Lang))
	}

	res := &CashbackInfoResponse{}
	info, err := client.Do(rq, res)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}
