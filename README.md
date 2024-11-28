# paypayopa-sdk-go

[日本語](./README.jp.md)

- ⚠️ This SDK is **Unofficial**. ⚠️
- Package `paypayopa` provides SDK to use PayPay API for Go.
- Official: https://developer.paypay.ne.jp

## Status

[![Check codes](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/check-code.yaml/badge.svg)](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/check-code.yaml)

[![Scan Vulnerabilities](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/scan-vulnerabilities.yaml/badge.svg)](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/scan-vulnerabilities.yaml)

## Requirements

- Go 1.21 or above.
- Docker (for Development)

### For development

- [golangci-lint](https://golangci-lint.run)

- [mockery](https://github.com/vektra/mockery)

## Install

Get it with `go get`.

```bash
go get github.com/mythrnr/paypayopa-sdk-go
```

## Feature

### Integration based SDK

The official `paypay/paypayopa-sdk-*` will obviously do the job.
But, to focus to integrate SDK to your system, `mythrnr/paypayopa-sdk-go`
provides clients each integration such as Web Payment or Native Payment,
not based REST API.

## Usage

Available integration here.

|Integration|Document|
|-|-|
|Web Payment|https://developer.paypay.ne.jp/products/docs/webpayment|
|Native Payment|https://developer.paypay.ne.jp/products/docs/nativepayment|
|Dynamic QR|https://developer.paypay.ne.jp/products/docs/qrcode|
|App Invoke|https://developer.paypay.ne.jp/products/docs/appinvoke|
|Continuous Payment|https://developer.paypay.ne.jp/products/docs/continuouspayment|
|PreAuth & Capture|https://developer.paypay.ne.jp/products/docs/preauthcapture|
|Request Money|https://developer.paypay.ne.jp/products/docs/pendingpayment|

### Create Credentials

- Use `paypayopa.NewCredentials` , and set the API key generated on the developer page.
- Use `paypayopa.Env***` to switch the connection destination.

```go
creds := paypayopa.NewCredentials(
    paypayopa.EnvSandbox,
    "YOUR_API_KEY",
    "YOUR_API_KEY_SECRET",
    "YOUR_MERCHANT_ID",
)
```

### Example to create and delete QR code using Web Payment

```go
package main

import (
    "context"
    "encoding/json"
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

    res, info, err := wp.CreateCode(ctx, &paypayopa.CreateQrCodePayload{
        MerchantPaymentID: uuid.NewString(),
        Amount: &paypayopa.MoneyAmount{
            Amount:   1000,
            Currency: paypayopa.CurrencyJPY,
        },
        CodeType:     paypayopa.CodeTypeOrderQR,
        RedirectURL:  "https://localhost",
        RedirectType: paypayopa.RedirectTypeWebLink,
    })

    if err != nil {
        log.Fatalf("%+v", err)
    }

    if !info.Success() {
        log.Fatalf("%+v", info)
    }

    info, err = wp.DeleteCode(ctx, res.CodeID)
    if err != nil {
        log.Fatalf("%+v", err)
    }
}
```
