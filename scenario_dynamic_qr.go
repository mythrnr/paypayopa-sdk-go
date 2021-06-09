package paypayopa

import "net/http"

// DynamicQR provides an API for PayPay's Dynamic QR functionality.
//
// DynamicQR は PayPay の動的ユーザスキャン機能の API を提供する.
//
// Docs
//
// https://developer.paypay.ne.jp/products/docs/qrcode
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/dynamicqrcode
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/dynamicqrcode
type DynamicQR WebPayment

// NewDynamicQR returns a client for Dynamic QR.
//
// NewDynamicQR は Dynamic QR のクライアントを返す.
func NewDynamicQR(creds *Credentials) *DynamicQR {
	return &DynamicQR{client: newClient(creds)}
}

// NewDynamicQRWithHTTPClient returns a Dynamic QR client
// that performs with a pre-configured *http.Client.
//
// NewDynamicQRWithHTTPClient は設定済みの *http.Client を用いて通信を行う
// Dynamic QR のクライアントを返す.
func NewDynamicQRWithHTTPClient(
	creds *Credentials,
	client *http.Client,
) *DynamicQR {
	return &DynamicQR{client: newClientWithHTTPClient(creds, client)}
}
