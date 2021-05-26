package paypayopa

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

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

func createQRCode(
	ctx context.Context,
	client *opaClient,
	req *CreateQrCodePayload,
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

func getRefundDetails(
	ctx context.Context,
	client *opaClient,
	merchantRefundID string,
) (*Refund, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &Refund{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v2/refunds/"+merchantRefundID,
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func refundPayment(
	ctx context.Context,
	client *opaClient,
	req *RefundPaymentPayload,
) (*Refund, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &Refund{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v2/refunds",
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func revertPaymentAuthorization(
	ctx context.Context,
	client *opaClient,
	req *RevertPaymentAuthorizationPayload,
) (*RevertedPayment, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &RevertedPayment{}
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
