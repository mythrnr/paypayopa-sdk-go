package paypayopa

import (
	"context"
	"net/http"
)

// ContinuousPayment provides an API for PayPay's
// Continuous Payment functionality.
//
// ContinuousPayment は PayPay の継続課金機能の API を提供する.
//
// Docs
//
// https://developer.paypay.ne.jp/products/docs/continuouspayment
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/continuous_payments
type ContinuousPayment struct{ client opaClient }

// NewContinuousPayment returns a client for Continuous Payment.
//
// NewContinuousPayment は Continuous Payment のクライアントを返す.
func NewContinuousPayment(creds *Credential) *ContinuousPayment {
	return &ContinuousPayment{client: newClient(creds)}
}

// NewContinuousPaymentWithHTTPClient returns a Continuous Payment client
// that performs with a pre-configured *http.Client.
//
// NewContinuousPaymentWithHTTPClient は設定済みの *http.Client を用いて通信を行う
// Continuous Payment のクライアントを返す.
func NewContinuousPaymentWithHTTPClient(
	creds *Credential,
	client *http.Client,
) *ContinuousPayment {
	return &ContinuousPayment{client: newClientWithHTTPClient(creds, client)}
}

// CreateContinuousPayment create a continuous payment
// and start the money transfer.
//
// CreateContinuousPayment は継続課金リクエストを作成して送金を開始する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/createPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/continuous_payments#operation/createPayment
func (c *ContinuousPayment) CreateContinuousPayment(
	ctx context.Context,
	req *CreateContinuousPaymentPayload,
) (*Payment, *ResultInfo, error) {
	return createContinuousPayment(ctx, c.client, req)
}

// GetPaymentDetails retrieves the details of a payment.
//
// GetPaymentDetails は決済の詳細を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/getPaymentDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/continuous_payments#operation/getPaymentDetails
func (c *ContinuousPayment) GetPaymentDetails(
	ctx context.Context,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	return getPaymentDetails(ctx, c.client, merchantPaymentID)
}

// CancelPayment cancels the payment.
//
// CancelPayment は支払いのキャンセルをする.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/cancelPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/continuous_payments#operation/cancelPayment
func (c *ContinuousPayment) CancelPayment(
	ctx context.Context,
	merchantPaymentID string,
) (*ResultInfo, error) {
	return cancelPayment(ctx, c.client, merchantPaymentID)
}

// RefundPayment refunds the payment.
//
// RefundPayment は返金を行う.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/refundPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/continuous_payments#operation/refundPayment
func (c *ContinuousPayment) RefundPayment(
	ctx context.Context,
	req *RefundPaymentPayload,
) (*Refund, *ResultInfo, error) {
	return refundPayment(ctx, c.client, req)
}

// GetRefundDetails gets the refund details.
//
// GetRefundDetails は返金の詳細を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/getRefundDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/continuous_payments#operation/getRefundDetails
func (c *ContinuousPayment) GetRefundDetails(
	ctx context.Context,
	merchantRefundID string,
) (*Refund, *ResultInfo, error) {
	return getRefundDetails(ctx, c.client, merchantRefundID)
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
func (c *ContinuousPayment) CreateAccountLinkQRCode(
	ctx context.Context,
	req *CreateAccountLinkQrCodePayload,
) (*CreateAccountLinkQrCodeResponse, *ResultInfo, error) {
	return createAccountLinkQrCode(ctx, c.client, req)
}

// UnlinkUser unlinks an user from the client.
//
// UnlinkUser はクライアントからユーザーのリンクを解除する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/unlinkUser
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/continuous_payments#operation/unlinkUser
func (c *ContinuousPayment) UnlinkUser(
	ctx context.Context,
	userAuthorizationID string,
) (*ResultInfo, error) {
	return unlinkUser(ctx, c.client, userAuthorizationID)
}

// GetUserAuthorizationStatus gets the authorization status of a user.
//
// GetUserAuthorizationStatus はユーザー認可状態を取得する.
//
// API Docs
//
// nolint:lll
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/getUserAuthorizationStatus
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/continuous_payments#operation/getUserAuthorizationStatus
func (c *ContinuousPayment) GetUserAuthorizationStatus(
	ctx context.Context,
	userAuthorizationID string,
) (*GetUserAuthorizationStatusResponse, *ResultInfo, error) {
	return getUserAuthorizationStatus(ctx, c.client, userAuthorizationID)
}

// GetMaskedUserProfile retrieves the masked phone number of the user.
//
// GetMaskedUserProfile はマスクされたユーザーの電話番号を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/continuous_payments#operation/getMaskedUserProfile
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/continuous_payments#operation/getMaskedUserProfile
func (c *ContinuousPayment) GetMaskedUserProfile(
	ctx context.Context,
	userAuthorizationID string,
) (*MaskedUserProfileResponse, *ResultInfo, error) {
	return getMaskedUserProfile(ctx, c.client, userAuthorizationID)
}
