# paypayopa-sdk-go

[English](./README.md)

> **注意**: この SDK は **非公式** である.
> 公式ドキュメントは <https://developer.paypay.ne.jp> を参照すること.

パッケージ `paypayopa` は
[PayPay Open Payment API](https://developer.paypay.ne.jp)
の Go SDK を提供する.

[![Go Reference](https://pkg.go.dev/badge/github.com/mythrnr/paypayopa-sdk-go.svg)](https://pkg.go.dev/github.com/mythrnr/paypayopa-sdk-go)
[![Check codes](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/check-code.yaml/badge.svg)](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/check-code.yaml)
[![Scan Vulnerabilities](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/scan-vulnerabilities.yaml/badge.svg)](https://github.com/mythrnr/paypayopa-sdk-go/actions/workflows/scan-vulnerabilities.yaml)

## 要件

- Go 1.24 以上

## インストール

```bash
go get github.com/mythrnr/paypayopa-sdk-go
```

## 設計

REST API ベースの SDK と異なり、この SDK は
[インテグレーションシナリオ](https://developer.paypay.ne.jp/products)
ごとにクライアントを提供する.
各クライアントは該当する決済フローに関連するメソッドのみを公開する.

## クイックスタート

### 1. 資格情報の生成

[PayPay 開発者ページ](https://developer.paypay.ne.jp)
で生成した API キーを使い `paypayopa.NewCredentials` を呼び出す.

```go
creds := paypayopa.NewCredentials(
    paypayopa.EnvSandbox,
    "YOUR_API_KEY",
    "YOUR_API_KEY_SECRET",
    "YOUR_MERCHANT_ID",
)
```

### 2. クライアントの生成と API 呼び出し

以下の例では Web Payment の QR コードを作成し、削除する.

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

より完全なサンプルアプリケーションは
[paypay-sample-ecommerce-backend-go](https://github.com/mythrnr/paypay-sample-ecommerce-backend-go)
を参照すること.

## 対応インテグレーション

| インテグレーション | ドキュメント | クライアント |
| ------------------ | ------------ | ------------ |
| Web Payment | [Doc](https://developer.paypay.ne.jp/products/docs/webpayment) | `NewWebPayment` |
| Native Payment | [Doc](https://developer.paypay.ne.jp/products/docs/nativepayment) | `NewNativePayment` |
| Dynamic QR | [Doc](https://developer.paypay.ne.jp/products/docs/qrcode) | `NewDynamicQR` |
| App Invoke | [Doc](https://developer.paypay.ne.jp/products/docs/appinvoke) | `NewAppInvoke` |
| Continuous Payment | [Doc](https://developer.paypay.ne.jp/products/docs/continuouspayment) | `NewContinuousPayment` |
| PreAuth & Capture | [Doc](https://developer.paypay.ne.jp/products/docs/preauthcapture) | `NewPreAuthCapture` |
| Request Money | [Doc](https://developer.paypay.ne.jp/products/docs/pendingpayment) | `NewRequestMoney` |

## 環境

以下の定数で接続先環境を指定する.

| 定数 | 環境 |
| ---- | ---- |
| `paypayopa.EnvProduction` | 本番 |
| `paypayopa.EnvStaging` | ステージング |
| `paypayopa.EnvSandbox` | サンドボックス (開発用) |

## エラーハンドリング

全ての API メソッドは `(response, *ResultInfo, error)` を返す.

- `error` はリクエスト自体が失敗した場合 (ネットワークエラー、マーシャリング失敗など)
  に非 nil となる.
- `ResultInfo.Success()` は HTTP ステータスコードが 400 未満の場合に `true` を返す.
- `ResultInfo.Code` と `ResultInfo.Message` に API レベルのエラー詳細が含まれる.

```go
res, info, err := client.CreateQRCode(ctx, payload)
if err != nil {
    // トランスポートまたはマーシャリングのエラー
    log.Fatal(err)
}

if !info.Success() {
    // API がエラーレスポンスを返した場合
    log.Printf("code=%s message=%s", info.Code, info.Message)
}
```

## カスタム HTTP クライアント

各クライアントのコンストラクタには、設定済みの `*http.Client` を受け取る
`WithHTTPClient` バリアントがある.
タイムアウト、プロキシ、トランスポートレベルの設定をカスタマイズする場合に使用する.

```go
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

wp := paypayopa.NewWebPaymentWithHTTPClient(creds, httpClient)
```

## Webhook

SDK は Webhook ペイロード用の型付き構造体を提供する.
`Webhook*` 型については
[パッケージドキュメント](https://pkg.go.dev/github.com/mythrnr/paypayopa-sdk-go)
を参照すること.

## ライセンス

[MIT](./LICENSE)
