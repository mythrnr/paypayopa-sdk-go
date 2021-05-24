package paypayopa

import (
	"context"
	"net/http"
)

// PreAuthCapture provides an API for PayPay's PreAuth & Capture functionality.
//
// PreAuthCapture は PayPay の都度課金（出荷売上）機能の API を提供する.
//
// Docs
//
// https://developer.paypay.ne.jp/products/docs/preauthcapture
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture
type PreAuthCapture struct{ client opaClient }

// NewPreAuthCapture returns a client for PreAuth & Capture.
//
// NewPreAuthCapture は PreAuth & Capture のクライアントを返す.
func NewPreAuthCapture(creds *Credential) *PreAuthCapture {
	return &PreAuthCapture{client: newClient(creds)}
}

// NewPreAuthCaptureWithHTTPClient returns a PreAuth & Capture client
// that performs with a pre-configured *http.Client.
//
// NewPreAuthCaptureWithHTTPClient は設定済みの *http.Client を用いて通信を行う
// PreAuth & Capture のクライアントを返す.
func NewPreAuthCaptureWithHTTPClient(
	creds *Credential,
	client *http.Client,
) *PreAuthCapture {
	return &PreAuthCapture{client: newClientWithHTTPClient(creds, client)}
}

// ConsultExpectedCashbackInfo consult expected cashback info
// before placing the order.
//
// ConsultExpectedCashbackInfo は注文する前に, 予定のキャッシュバック情報を参照する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/consultExpectedCashbackInfo
//
// nolint:lll
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/consultExpectedCashbackInfo
func (p *PreAuthCapture) ConsultExpectedCashbackInfo(
	ctx context.Context,
	req *ConsultExpectedCashbackInfoPayload,
) (*CashbackInfoResponse, *ResultInfo, error) {
	return consultExpectedCashbackInfo(ctx, p.client, req)
}

// CreatePaymentAuthorization creates a payment authorization
// to block the money.
//
// CreatePaymentAuthorization は決済金額をブロックするための
// 支払いAuthorizationを作成する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/createAuth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/createAuth
func (p *PreAuthCapture) CreatePaymentAuthorization(
	ctx context.Context,
	req *CreatePaymentAuthorizationPayload,
) (*Payment, *ResultInfo, error) {
	return createPaymentAuthorization(ctx, p.client, req)
}

// GetPaymentDetails retrieves the details of a payment.
//
// GetPaymentDetails は決済の詳細を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/getPaymentDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/getPaymentDetails
func (p *PreAuthCapture) GetPaymentDetails(
	ctx context.Context,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	return getPaymentDetails(ctx, p.client, merchantPaymentID)
}

// CancelPaymentAuthorization cancels a payment request.
//
// CancelPaymentAuthorization は支払リクエストのキャンセルを行う.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/cancelPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/cancelPayment
func (p *PreAuthCapture) CancelPaymentAuthorization(
	ctx context.Context,
	merchantPaymentID string,
) (*ResultInfo, error) {
	return cancelPayment(ctx, p.client, merchantPaymentID)
}

// CapturePaymentAuthorization captures the payment authorization for a payment.
//
// CapturePaymentAuthorization は事前にユーザーの残高からブロックしていた
// 決済金額をキャプチャ（決済）する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/capturePaymentAuth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/capturePaymentAuth
func (p *PreAuthCapture) CapturePaymentAuthorization(
	ctx context.Context,
	req *CapturePaymentAuthorizationPayload,
) (*Payment, *ResultInfo, error) {
	return capturePaymentAuthorization(ctx, p.client, req)
}

// RevertPaymentAuthorization canels the payment authorization.
//
// RevertPaymentAuthorization はユーザーの残高から決済金額を
// ブロックしている状態をキャンセルする.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/revertAuth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/revertAuth
func (p *PreAuthCapture) RevertPaymentAuthorization(
	ctx context.Context,
	req *RevertPaymentAuthorizationPayload,
) (*RevertedPayment, *ResultInfo, error) {
	return revertPaymentAuthorization(ctx, p.client, req)
}

// RefundPayment refunds the payment.
//
// RefundPayment は返金を行う.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/refundPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/refundPayment
func (p *PreAuthCapture) RefundPayment(
	ctx context.Context,
	req *RefundPaymentPayload,
) (*Refund, *ResultInfo, error) {
	return refundPayment(ctx, p.client, req)
}

// GetRefundDetails gets the refund details.
//
// GetRefundDetails は返金の詳細を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/getRefundDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/getRefundDetails
func (p *PreAuthCapture) GetRefundDetails(
	ctx context.Context,
	merchantRefundID string,
) (*Refund, *ResultInfo, error) {
	return getRefundDetails(ctx, p.client, merchantRefundID)
}

// CreateTopupQRCode creates qr code for the user.
//
// CreateTopupQRCode はトップアップ用のQRコードを作成する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/createTopUpQRCode
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/createTopUpQRCode
func (p *PreAuthCapture) CreateTopupQRCode(
	ctx context.Context,
	req *CreateTopupQRCodePayload,
) (*TopupQRCodeResponse, *ResultInfo, error) {
	return createTopupQRCode(ctx, p.client, req)
}

// GetTopupDetails gets the details of topup done using QR Code.
//
// GetTopupDetails はトップアップ用のQRコードの詳細を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/getTopUpQRDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/getTopUpQRDetails
func (p *PreAuthCapture) GetTopupDetails(
	ctx context.Context,
	merchantTopUpID string,
) (*TopupQRCodeDetailsResponse, *ResultInfo, error) {
	return getTopupDetails(ctx, p.client, merchantTopUpID)
}

// DeleteTopupQRCode deletes the topup QR code.
//
// DeleteTopupQRCode はトップアップ用のQRコードを削除する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/deleteTopUpQrCode
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/deleteTopUpQrCode
func (p *PreAuthCapture) DeleteTopupQRCode(
	ctx context.Context,
	codeID string,
) (*ResultInfo, error) {
	return deleteTopupQRCode(ctx, p.client, codeID)
}

// GetUserWalletBalance gets Total Amount in User Wallet.
//
// GetUserWalletBalance はユーザーの残高を参照する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/getBalance
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/getBalance
func (p *PreAuthCapture) GetUserWalletBalance(
	ctx context.Context,
	req *GetUserWalletBalancePayload,
) (*UserWalletBalanceResponse, *ResultInfo, error) {
	return getUserWalletBalance(ctx, p.client, req)
}

// CheckUserWalletBalance check sif user has enough balance
// to make a payment Total Amount in User Wallet.
//
// CheckUserWalletBalance はユーザーが支払いをするための十分な残高があるかを確認する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/checkWalletBalance
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/checkWalletBalance
func (p *PreAuthCapture) CheckUserWalletBalance(
	ctx context.Context,
	req *CheckUserWalletBalancePayload,
) (*CheckUserWalletBalance, *ResultInfo, error) {
	return checkUserWalletBalance(ctx, p.client, req)
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
func (p *PreAuthCapture) CreateAccountLinkQRCode(
	ctx context.Context,
	req *CreateAccountLinkQrCodePayload,
) (*CreateAccountLinkQrCodeResponse, *ResultInfo, error) {
	return createAccountLinkQrCode(ctx, p.client, req)
}

// UnlinkUser unlinks an user from the client.
//
// UnlinkUser はクライアントからユーザーのリンクを解除する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/unlinkUser
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/unlinkUser
func (p *PreAuthCapture) UnlinkUser(
	ctx context.Context,
	userAuthorizationID string,
) (*ResultInfo, error) {
	return unlinkUser(ctx, p.client, userAuthorizationID)
}

// GetUserAuthorizationStatus gets the authorization status of a user.
//
// GetUserAuthorizationStatus はユーザー認可状態を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/getUserAuthorizationStatus
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/getUserAuthorizationStatus
func (p *PreAuthCapture) GetUserAuthorizationStatus(
	ctx context.Context,
	userAuthorizationID string,
) (*GetUserAuthorizationStatusResponse, *ResultInfo, error) {
	return getUserAuthorizationStatus(ctx, p.client, userAuthorizationID)
}

// GetMaskedUserProfile retrieves the masked phone number of the user.
//
// GetMaskedUserProfile はマスクされたユーザーの電話番号を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/preauth_capture#operation/getMaskedUserProfile
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/preauth_capture#operation/getMaskedUserProfile
func (p *PreAuthCapture) GetMaskedUserProfile(
	ctx context.Context,
	userAuthorizationID string,
) (*MaskedUserProfileResponse, *ResultInfo, error) {
	return getMaskedUserProfile(ctx, p.client, userAuthorizationID)
}
