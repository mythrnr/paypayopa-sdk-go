package paypayopa

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

func cancelPayment(
	ctx context.Context,
	client opaClient,
	merchantPaymentID string,
) (*ResultInfo, error) {
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	return client.DELETE(ctx, "/v2/payments/"+merchantPaymentID)
}

func cancelPendingOrder(
	ctx context.Context,
	client opaClient,
	merchantPaymentID string,
) (*ResultInfo, error) {
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	return client.DELETE(ctx, "/v1/requestOrder/"+merchantPaymentID)
}

func capturePaymentAuthorization(
	ctx context.Context,
	client opaClient,
	req *CapturePaymentAuthorizationPayload,
) (*Payment, *ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &Payment{}
	info, err := client.POST(
		ctx,
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
	client opaClient,
	req *ConsultExpectedCashbackInfoPayload,
) (*CashbackInfoResponse, *ResultInfo, error) {
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	rq, err := client.Request(
		ctx,
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
	client opaClient,
	req *CreateQrCodePayload,
) (*QRCodeResponse, *ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &QRCodeResponse{}

	info, err := client.POST(
		ctx,
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
	client opaClient,
	req *CreateContinuousPaymentPayload,
) (*Payment, *ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &Payment{}
	info, err := client.POST(
		ctx,
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
	client opaClient,
	req *CreatePaymentPayload,
) (*Payment, *ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &Payment{}
	info, err := client.POST(
		ctx,
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
	client opaClient,
	req *CreatePaymentAuthorizationPayload,
) (*Payment, *ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &Payment{}
	info, err := client.POST(
		ctx,
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
	client opaClient,
	req *CreatePendingPaymentPayload,
) (*PendingPayment, *ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &PendingPayment{}
	info, err := client.POST(
		ctx,
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
	client opaClient,
	codeID string,
) (*ResultInfo, error) {
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	return client.DELETE(ctx, "/v2/codes/"+codeID)
}

func getCodePaymentDetails(
	ctx context.Context,
	client opaClient,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &Payment{}
	info, err := client.GET(
		ctx,
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
	client opaClient,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &Payment{}
	info, err := client.GET(
		ctx,
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
	client opaClient,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &Payment{}
	info, err := client.GET(
		ctx,
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
	client opaClient,
	merchantRefundID string,
) (*Refund, *ResultInfo, error) {
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &Refund{}
	info, err := client.GET(
		ctx,
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
	client opaClient,
	req *RefundPaymentPayload,
) (*Refund, *ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &Refund{}
	info, err := client.POST(
		ctx,
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
	client opaClient,
	req *RevertPaymentAuthorizationPayload,
) (*RevertedPayment, *ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &RevertedPayment{}
	info, err := client.POST(
		ctx,
		"/v2/payments/preauthorize/revert",
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}
