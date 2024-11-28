# paypayopa-sdk-go

[English](./README.md)

- ⚠️ この SDK は **非公式** . ⚠️
- `paypayopa` は PayPay API を Go で使うための SDK を提供する.
- 公式サイト: https://developer.paypay.ne.jp

## Status

[![Check codes](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/check-code.yaml/badge.svg)](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/check-code.yaml)

[![Scan Vulnerabilities](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/scan-vulnerabilities.yaml/badge.svg)](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/scan-vulnerabilities.yaml)

## Requirements

- Go 1.21 以上.
- Docker (開発時)

### 開発に必要

- [golangci-lint](https://golangci-lint.run)

- [mockery](https://github.com/vektra/mockery)

## Install

`go get` で取得する.

```bash
go get github.com/mythrnr/paypayopa-sdk-go
```

## Feature

### インテグレーションベースの SDK

公式の `paypay/paypayopa-sdk-*` も当然ながら十分に役割を果たすが, 
`mythrnr/paypayopa-sdk-go` は SDK の統合に更に集中するために
Web Payment や Native Payment などのインテグレーションごとに
クライアントを生成して利用する.

## Usage

使用可能なインテグレーションは下記の通り.

|インテグレーション|ドキュメント|
|-|-|
|Web Payment|https://developer.paypay.ne.jp/products/docs/webpayment|
|Native Payment|https://developer.paypay.ne.jp/products/docs/nativepayment|
|Dynamic QR|https://developer.paypay.ne.jp/products/docs/qrcode|
|App Invoke|https://developer.paypay.ne.jp/products/docs/appinvoke|
|Continuous Payment|https://developer.paypay.ne.jp/products/docs/continuouspayment|
|PreAuth & Capture|https://developer.paypay.ne.jp/products/docs/preauthcapture|
|Request Money|https://developer.paypay.ne.jp/products/docs/pendingpayment|

### 資格情報の生成

- `paypayopa.NewCredentials` を使い, 開発者ページで生成した API キーなどを設定する.
- `paypayopa.Env***` を指定し, 接続先を切り替える.

```go
creds := paypayopa.NewCredentials(
    paypayopa.EnvSandbox,
    "YOUR_API_KEY",
    "YOUR_API_KEY_SECRET",
    "YOUR_MERCHANT_ID",
)
```

### Web Payment を用いた QR コードの作成と削除の例

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
