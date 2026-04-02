# paypayopa-sdk-go

[Japanese (日本語)](./README.jp.md)

> **Warning**: This SDK is **unofficial**.
> For the official documentation, see <https://developer.paypay.ne.jp>.

Package `paypayopa` provides a Go SDK for the
[PayPay Open Payment API](https://developer.paypay.ne.jp).

[![Go Reference](https://pkg.go.dev/badge/github.com/mythrnr/paypayopa-sdk-go.svg)](https://pkg.go.dev/github.com/mythrnr/paypayopa-sdk-go)
[![Check codes](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/check-code.yaml/badge.svg)](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/check-code.yaml)
[![Scan Vulnerabilities](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/scan-vulnerabilities.yaml/badge.svg)](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/scan-vulnerabilities.yaml)

## Requirements

- Go 1.24 or above

## Installation

```bash
go get github.com/mythrnr/paypayopa-sdk-go
```

## Design

Unlike REST-API-based SDKs, this SDK provides clients per
[integration scenario](https://developer.paypay.ne.jp/products),
so each client exposes only the methods relevant to that payment flow.

## Quick Start

### 1. Create Credentials

Use `paypayopa.NewCredentials` with the API key generated on the
[PayPay developer page](https://developer.paypay.ne.jp).

```go
creds := paypayopa.NewCredentials(
    paypayopa.EnvSandbox,
    "YOUR_API_KEY",
    "YOUR_API_KEY_SECRET",
    "YOUR_MERCHANT_ID",
)
```

### 2. Create a Client and Call an API

The following example creates a QR code for Web Payment and then deletes it.

```go
package main

import (
    "context"
    "log"

    "github.com/google/uuid"
    "github.com/mythrnr/paypayopa-sdk-go"
)

func main() {
    creds := paypayopa.NewCredentials(
        paypayopa.EnvSandbox,
        "YOUR_API_KEY",
        "YOUR_API_KEY_SECRET",
        "YOUR_MERCHANT_ID",
    )

    wp := paypayopa.NewWebPayment(creds)
    ctx := context.Background()

    res, info, err := wp.CreateQRCode(ctx, &paypayopa.CreateQRCodePayload{
        MerchantPaymentID: uuid.NewString(),
        Amount: &paypayopa.MoneyAmount{
            Amount:   1000,
            Currency: paypayopa.CurrencyJPY,
        },
        CodeType:     paypayopa.CodeTypeOrderQR,
        RedirectURL:  "https://example.com/callback",
        RedirectType: paypayopa.RedirectTypeWebLink,
    })

    if err != nil {
        log.Fatalf("request failed: %+v", err)
    }

    if !info.Success() {
        log.Fatalf("API error: %s (code: %s)", info.Message, info.Code)
    }

    log.Printf("QR Code URL: %s", res.URL)

    if _, err = wp.DeleteQRCode(ctx, res.CodeID); err != nil {
        log.Fatalf("delete failed: %+v", err)
    }
}
```

For a more complete example application, see
[paypay-sample-ecommerce-backend-go](https://github.com/mythrnr/paypay-sample-ecommerce-backend-go).

## Supported Integrations

| Integration | Document | Client |
| ----------- | -------- | ------ |
| Web Payment | [Doc](https://developer.paypay.ne.jp/products/docs/webpayment) | `NewWebPayment` |
| Native Payment | [Doc](https://developer.paypay.ne.jp/products/docs/nativepayment) | `NewNativePayment` |
| Dynamic QR | [Doc](https://developer.paypay.ne.jp/products/docs/qrcode) | `NewDynamicQR` |
| App Invoke | [Doc](https://developer.paypay.ne.jp/products/docs/appinvoke) | `NewAppInvoke` |
| Continuous Payment | [Doc](https://developer.paypay.ne.jp/products/docs/continuouspayment) | `NewContinuousPayment` |
| PreAuth & Capture | [Doc](https://developer.paypay.ne.jp/products/docs/preauthcapture) | `NewPreAuthCapture` |
| Request Money | [Doc](https://developer.paypay.ne.jp/products/docs/pendingpayment) | `NewRequestMoney` |

## Environments

Use one of the following constants to specify the target environment.

| Constant | Environment |
| -------- | ----------- |
| `paypayopa.EnvProduction` | Production |
| `paypayopa.EnvStaging` | Staging |
| `paypayopa.EnvSandbox` | Sandbox (for development) |

## Error Handling

All API methods return `(response, *ResultInfo, error)`.

- `error` is non-nil when the request itself fails
  (e.g. network error, marshaling failure).
- `ResultInfo.Success()` returns `true` when the HTTP status code is less than 400.
- `ResultInfo.Code` and `ResultInfo.Message` contain API-level error details.

```go
res, info, err := client.CreateQRCode(ctx, payload)
if err != nil {
    // Transport or marshaling error
    log.Fatal(err)
}

if !info.Success() {
    // API returned an error response
    log.Printf("code=%s message=%s", info.Code, info.Message)
}
```

## Custom HTTP Client

Each client constructor has a `WithHTTPClient` variant that accepts
a pre-configured `*http.Client`. This is useful for setting custom
timeouts, proxies, or transport-level configurations.

```go
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

wp := paypayopa.NewWebPaymentWithHTTPClient(creds, httpClient)
```

## Webhooks

The SDK provides typed structs for webhook payloads.
See the `Webhook*` types in the
[package documentation](https://pkg.go.dev/github.com/mythrnr/paypayopa-sdk-go).

## License

[MIT](./LICENSE)
