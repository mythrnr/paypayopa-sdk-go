package paypayopa

import "net/http"

// AppInvoke provides an API for PayPay's App Invoke functionality.
//
// AppInvoke は PayPay のアプリコール機能の API を提供する.
//
// # Docs
//
// https://developer.paypay.ne.jp/products/docs/appinvoke
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/appinvoke
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/appinvoke
type AppInvoke WebPayment

// NewAppInvoke returns a client for App Invoke.
//
// NewAppInvoke は App Invoke のクライアントを返す.
func NewAppInvoke(creds *Credentials) *AppInvoke {
	return &AppInvoke{client: newClient(creds)}
}

// NewAppInvokeWithHTTPClient returns a App Invoke client
// that performs with a pre-configured *http.Client.
//
// NewAppInvokeWithHTTPClient は設定済みの *http.Client を用いて通信を行う
// App Invoke のクライアントを返す.
func NewAppInvokeWithHTTPClient(
	creds *Credentials,
	client *http.Client,
) *AppInvoke {
	return &AppInvoke{client: newClientWithHTTPClient(creds, client)}
}
