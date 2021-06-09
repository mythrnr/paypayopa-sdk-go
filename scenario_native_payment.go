package paypayopa

import (
	"context"
	"net/http"
)

// NativePayment provides an API for PayPay's Native Payment functionality.
//
// NativePayment は PayPay のネイティブペイメント機能の API を提供する.
//
// Docs
//
// https://developer.paypay.ne.jp/products/docs/nativepayment
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit
type NativePayment struct {
	client *opaClient
	creds  *Credentials
}

// NewNativePayment returns a client for Native Payment.
//
// NewNativePayment は Native Payment のクライアントを返す.
func NewNativePayment(creds *Credentials) *NativePayment {
	return &NativePayment{
		client: newClient(creds),
		creds:  creds,
	}
}

// NewNativePaymentWithHTTPClient returns a Native Payment client
// that performs with a pre-configured *http.Client.
//
// NewNativePaymentWithHTTPClient は設定済みの *http.Client を用いて通信を行う
// Native Payment のクライアントを返す.
func NewNativePaymentWithHTTPClient(
	creds *Credentials,
	client *http.Client,
) *NativePayment {
	return &NativePayment{
		client: newClientWithHTTPClient(creds, client),
		creds:  creds,
	}
}

// ConsultExpectedCashbackInfo consult expected cashback info
// before placing the order.
//
// ConsultExpectedCashbackInfo は注文する前に, 予定のキャッシュバック情報を参照する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/consultExpectedCashbackInfo
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/consultExpectedCashbackInfo
func (n *NativePayment) ConsultExpectedCashbackInfo(
	ctx context.Context,
	req *ConsultExpectedCashbackInfoPayload,
) (*CashbackInfoResponse, *ResultInfo, error) {
	return consultExpectedCashbackInfo(ctx, n.client, req)
}

// CreatePayment create a naitve payment and start the money transfer.
//
// CreatePayment は決済リクエストを作成して送金を開始する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/consultExpectedCashbackInfo
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/consultExpectedCashbackInfo
func (n *NativePayment) CreatePayment(
	ctx context.Context,
	req *CreatePaymentPayload,
) (*Payment, *ResultInfo, error) {
	return createPayment(ctx, n.client, req)
}

// GetPaymentDetails retrieves the details of a payment.
//
// GetPaymentDetails は決済の詳細を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/getPaymentDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/getPaymentDetails
func (n *NativePayment) GetPaymentDetails(
	ctx context.Context,
	merchantPaymentID string,
) (*Payment, *ResultInfo, error) {
	return getPaymentDetails(ctx, n.client, merchantPaymentID)
}

// CancelPayment cancels the payment.
//
// CancelPayment は支払いのキャンセルをする.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/cancelPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/cancelPayment
func (n *NativePayment) CancelPayment(
	ctx context.Context,
	merchantPaymentID string,
) (*ResultInfo, error) {
	return cancelPayment(ctx, n.client, merchantPaymentID)
}

// RefundPayment refunds the payment.
//
// RefundPayment は返金を行う.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/refundPayment
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/refundPayment
func (n *NativePayment) RefundPayment(
	ctx context.Context,
	req *RefundPaymentPayload,
) (*RefundResponse, *ResultInfo, error) {
	return refundPayment(ctx, n.client, req)
}

// GetRefundDetails gets the refund details.
//
// GetRefundDetails は返金の詳細を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/getRefundDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/getRefundDetails
func (n *NativePayment) GetRefundDetails(
	ctx context.Context,
	merchantRefundID string,
) (*RefundResponse, *ResultInfo, error) {
	return getRefundDetails(ctx, n.client, merchantRefundID)
}

// CreateTopupQRCode creates qr code for the user.
//
// CreateTopupQRCode はトップアップ用のQRコードを作成する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/createTopUpQRCode
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/createTopUpQRCode
func (n *NativePayment) CreateTopupQRCode(
	ctx context.Context,
	req *CreateTopupQRCodePayload,
) (*TopupQRCodeResponse, *ResultInfo, error) {
	return createTopupQRCode(ctx, n.client, req)
}

// GetTopupDetails gets the details of topup done using QR Code.
//
// GetTopupDetails はトップアップ用のQRコードの詳細を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/getTopUpQRDetails
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/getTopUpQRDetails
func (n *NativePayment) GetTopupDetails(
	ctx context.Context,
	merchantTopUpID string,
) (*TopupQRCodeDetailsResponse, *ResultInfo, error) {
	return getTopupDetails(ctx, n.client, merchantTopUpID)
}

// DeleteTopupQRCode deletes the topup QR code.
//
// DeleteTopupQRCode はトップアップ用のQRコードを削除する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/deleteTopUpQrCode
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/deleteTopUpQrCode
func (n *NativePayment) DeleteTopupQRCode(
	ctx context.Context,
	codeID string,
) (*ResultInfo, error) {
	return deleteTopupQRCode(ctx, n.client, codeID)
}

// GetUserWalletBalance gets Total Amount in User Wallet.
//
// GetUserWalletBalance はユーザーの残高を参照する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/getBalance
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/getBalance
func (n *NativePayment) GetUserWalletBalance(
	ctx context.Context,
	req *GetUserWalletBalancePayload,
) (*UserWalletBalanceResponse, *ResultInfo, error) {
	return getUserWalletBalance(ctx, n.client, req)
}

// CheckUserWalletBalance check sif user has enough balance
// to make a payment Total Amount in User Wallet.
//
// CheckUserWalletBalance はユーザーが支払いをするための十分な残高があるかを確認する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/checkWalletBalance
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/checkWalletBalance
func (n *NativePayment) CheckUserWalletBalance(
	ctx context.Context,
	req *CheckUserWalletBalancePayload,
) (*CheckUserWalletBalance, *ResultInfo, error) {
	return checkUserWalletBalance(ctx, n.client, req)
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
func (n *NativePayment) CreateAccountLinkQRCode(
	ctx context.Context,
	req *CreateAccountLinkQRCodePayload,
) (*CreateAccountLinkQRCodeResponse, *ResultInfo, error) {
	return createAccountLinkQRCode(ctx, n.client, req)
}

// UnlinkUser unlinks an user from the client.
//
// UnlinkUser はクライアントからユーザーのリンクを解除する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/unlinkUser
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/unlinkUser
func (n *NativePayment) UnlinkUser(
	ctx context.Context,
	userAuthorizationID string,
) (*ResultInfo, error) {
	return unlinkUser(ctx, n.client, userAuthorizationID)
}

// GetUserAuthorizationStatus gets the authorization status of a user.
//
// GetUserAuthorizationStatus はユーザー認可状態を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/getUserAuthorizationStatus
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/getUserAuthorizationStatus
func (n *NativePayment) GetUserAuthorizationStatus(
	ctx context.Context,
	userAuthorizationID string,
) (*GetUserAuthorizationStatusResponse, *ResultInfo, error) {
	return getUserAuthorizationStatus(ctx, n.client, userAuthorizationID)
}

// GetMaskedUserProfile retrieves the masked phone number of the user.
//
// GetMaskedUserProfile はマスクされたユーザーの電話番号を取得する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/direct_debit#operation/getMaskedUserProfile
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/direct_debit#operation/getMaskedUserProfile
func (n *NativePayment) GetMaskedUserProfile(
	ctx context.Context,
	userAuthorizationID string,
) (*GetMaskedUserProfileResponse, *ResultInfo, error) {
	return getMaskedUserProfile(ctx, n.client, userAuthorizationID)
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
func (n *NativePayment) DecodeResponseToken(
	ctx context.Context,
	token string,
) (*AuthorizationResponseToken, error) {
	return decodeAuthorizationResponseToken(n.creds, token)
}
