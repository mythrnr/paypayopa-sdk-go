package paypayopa

import (
	"context"
	"net/http"
)

// WebPayment provides an API for PayPay's web payment functionality.
//
// WebPayment は PayPay のウェブペイメント機能の API を提供する.
//
// # Docs
//
// https://developer.paypay.ne.jp/products/docs/webpayment
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier
type WebPayment struct{ client *opaClient }

// NewWebPayment returns a client for Web Payment.
//
// NewWebPayment は Web Payment のクライアントを返す.
func NewWebPayment(creds *Credentials) *WebPayment {
	return &WebPayment{client: newClient(creds)}
}

// NewWebPaymentWithHTTPClient returns a Web Payment client
// that performs with a pre-configured *http.Client.
//
// NewWebPaymentWithHTTPClient は設定済みの *http.Client を用いて通信を行う
// Web Payment のクライアントを返す.
func NewWebPaymentWithHTTPClient(
	creds *Credentials,
	client *http.Client,
) *WebPayment {
	return &WebPayment{client: newClientWithHTTPClient(creds, client)}
}

// CreateQRCode creates a QR code for payment.
//
// CreateQRCode は支払い用の QR コードを作成する.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/createQRCode
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#operation/createQRCode
func (w *WebPayment) CreateQRCode(
	ctx context.Context,
	req *CreateQRCodePayload,
) (*QRCodeResponse, *ResultInfo, error) {
	return createQRCode(ctx, w.client, req)
}

// DeleteQRCode deletes a QR Code that has already been created.
//
// DeleteQRCode は作成済みの QR コードを削除する.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/deleteQRCode
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#operation/deleteQRCode
func (w *WebPayment) DeleteQRCode(
	ctx context.Context,
	codeID string,
) (*ResultInfo, error) {
	return deleteQRCode(ctx, w.client, codeID)
}

// GetPaymentDetails retrieves the details of a payment.
//
// GetPaymentDetails は決済の詳細を取得する.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/getPaymentDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#operation/getPaymentDetails
func (w *WebPayment) GetPaymentDetails(
	ctx context.Context,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	return getCodePaymentDetails(ctx, w.client, merchantPaymentID)
}

// CancelPayment cancels the payment.
//
// CancelPayment は支払いのキャンセルをする.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/cancelPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#operation/cancelPayment
func (w *WebPayment) CancelPayment(
	ctx context.Context,
	merchantPaymentID string,
) (*ResultInfo, error) {
	return cancelPayment(ctx, w.client, merchantPaymentID)
}

// CapturePaymentAuthorization captures
// the payment authorization for a payment.
//
// CapturePaymentAuthorization は支払い承認を取得する.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/capturePaymentAuth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#operation/capturePaymentAuth
func (w *WebPayment) CapturePaymentAuthorization(
	ctx context.Context,
	req *CapturePaymentAuthorizationPayload,
) (*Payment, *ResultInfo, error) {
	return capturePaymentAuthorization(ctx, w.client, req)
}

// RevertPaymentAuthorization cancels the payment authorization
// if the user cancels the order.
//
// RevertPaymentAuthorization はユーザーが注文のキャンセルをした場合に
// 支払い承認の取り消しをする.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/revertAuth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#operation/revertAuth
func (w *WebPayment) RevertPaymentAuthorization(
	ctx context.Context,
	req *RevertPaymentAuthorizationPayload,
) (*RevertedPaymentResponse, *ResultInfo, error) {
	return revertPaymentAuthorization(ctx, w.client, req)
}

// RefundPayment refunds the payment.
//
// RefundPayment は返金を行う.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/refundPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#operation/refundPayment
func (w *WebPayment) RefundPayment(
	ctx context.Context,
	req *RefundPaymentPayload,
) (*RefundResponse, *ResultInfo, error) {
	return refundPayment(ctx, w.client, req)
}

// GetRefundDetails gets the refund details.
//
// GetRefundDetails は返金の詳細を取得する.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#operation/getRefundDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#operation/getRefundDetails
func (w *WebPayment) GetRefundDetails(
	ctx context.Context,
	merchantRefundID string,
) (*RefundResponse, *ResultInfo, error) {
	return getRefundDetails(ctx, w.client, merchantRefundID)
}
