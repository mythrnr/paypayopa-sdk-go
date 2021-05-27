package paypayopa

import (
	"context"
	"net/http"
)

// RequestMoney provides an API for PayPay's payment request functionality.
//
// RequestMoney は PayPay の支払リクエスト機能の API を提供する.
//
// Docs
//
// https://developer.paypay.ne.jp/products/docs/pendingpayment
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/pending_payments
type RequestMoney struct {
	client *opaClient
	creds  *Credential
}

// NewRequestMoney returns a client for Request Money.
//
// NewRequestMoney は Request Money のクライアントを返す.
func NewRequestMoney(creds *Credential) *RequestMoney {
	return &RequestMoney{
		client: newClient(creds),
		creds:  creds,
	}
}

// NewRequestMoneyWithHTTPClient returns a Request Money client
// that performs with a pre-configured *http.Client.
//
// NewRequestMoneyWithHTTPClient は設定済みの *http.Client を用いて通信を行う
// Request Money のクライアントを返す.
func NewRequestMoneyWithHTTPClient(
	creds *Credential,
	client *http.Client,
) *RequestMoney {
	return &RequestMoney{
		client: newClientWithHTTPClient(creds, client),
		creds:  creds,
	}
}

// CreatePendingPayment sends a push notification
// to the user requesting payment.
//
// CreatePendingPayment はユーザーに支払いを求める通知を送信する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments#operation/createPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/pending_payments#operation/createPayment
func (r *RequestMoney) CreatePendingPayment(
	ctx context.Context,
	req *CreatePendingPaymentPayload,
) (*PendingPayment, *ResultInfo, error) {
	return createPendingPayment(ctx, r.client, req)
}

// GetPaymentDetails retrieves the details of a payment.
//
// GetPaymentDetails は決済の詳細を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments#operation/getPaymentDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/pending_payments#operation/getPaymentDetails
func (r *RequestMoney) GetPaymentDetails(
	ctx context.Context,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	return getRequestedPaymentDetails(ctx, r.client, merchantPaymentID)
}

// CancelPendingOrder deletes a pending payment.
//
// CancelPendingOrder は保留中の支払いを削除する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments#operation/cancelPendingOrder
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/pending_payments#operation/cancelPendingOrder
func (r *RequestMoney) CancelPendingOrder(
	ctx context.Context,
	merchantPaymentID string,
) (*ResultInfo, error) {
	return cancelPendingOrder(ctx, r.client, merchantPaymentID)
}

// RefundPayment refunds the payment.
//
// RefundPayment は返金を行う.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments#operation/refundPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/pending_payments#operation/refundPayment
func (r *RequestMoney) RefundPayment(
	ctx context.Context,
	req *RefundPaymentPayload,
) (*RefundResponse, *ResultInfo, error) {
	return refundPayment(ctx, r.client, req)
}

// GetRefundDetails gets the refund details.
//
// GetRefundDetails は返金の詳細を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments#operation/getRefundDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/pending_payments#operation/getRefundDetails
func (r *RequestMoney) GetRefundDetails(
	ctx context.Context,
	merchantRefundID string,
) (*RefundResponse, *ResultInfo, error) {
	return getRefundDetails(ctx, r.client, merchantRefundID)
}

// CreateAccountLinkQRCode creates a ACCOUNT LINK QR
// and display it to the user.
//
// CreateAccountLinkQRCode はアカウントリンクQRを作成し, ユーザーに表示する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/account_link.html#operation/createQRSession
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/account_link.html#operation/createQRSession
func (r *RequestMoney) CreateAccountLinkQRCode(
	ctx context.Context,
	req *CreateAccountLinkQRCodePayload,
) (*CreateAccountLinkQRCodeResponse, *ResultInfo, error) {
	return createAccountLinkQRCode(ctx, r.client, req)
}

// GetMaskedUserProfile retrieves the masked phone number of the user.
//
// GetMaskedUserProfile はマスクされたユーザーの電話番号を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/pending_payments#operation/getMaskedUserProfile
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/pending_payments#operation/getMaskedUserProfile
func (r *RequestMoney) GetMaskedUserProfile(
	ctx context.Context,
	userAuthorizationID string,
) (*GetMaskedUserProfileResponse, *ResultInfo, error) {
	return getMaskedUserProfile(ctx, r.client, userAuthorizationID)
}

// DecodeResponseToken decodes and returns the JWT of
// the user's authorization result.
//
// DecodeResponseToken はユーザーの認可の結果の JWT をデコードして返す.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/account_link.html
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/account_link.html
func (r *RequestMoney) DecodeResponseToken(
	ctx context.Context,
	token string,
) (*AuthorizationResponseToken, error) {
	return decodeAuthorizationResponseToken(r.creds, token)
}
