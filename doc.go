/*
Package paypayopa provides SDK to use PayPay Open Payment API.

Package paypayopa は, PayPay Open Payment API を使用するための SDK を提供する.

See: https://developer.paypay.ne.jp

Example to create and delete QR code using Web Payment

- For more example, see https://github.com/mythrnr/paypay-sample-ecommerce-backend-go

	package main

	import (
		"context"
		"encoding/json"
		"log"
		"os"

		"github.com/google/uuid"
		"github.com/mythrnr/paypayopa-sdk-go"
	)

	func main() {
		creds := paypayopa.NewCredential(
			paypayopa.EnvSandbox,
			os.Getenv("PAYPAYOPA_API_KEY"),
			os.Getenv("PAYPAYOPA_API_KEY_SECRET"),
			os.Getenv("PAYPAYOPA_MERCHANT_ID"),
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
			RedirectURL:  "https://localhost",
			RedirectType: paypayopa.RedirectTypeWebLink,
		})

		if err != nil {
			log.Fatalf("%+v", err)
		}

		b, _ := json.MarshalIndent(info, "", "  ")
		log.Println(string(b))

		if !info.Success() {
			log.Fatalf("%+v", info)
		}

		b, _ = json.MarshalIndent(res, "", "  ")
		log.Println(string(b))

		info, err = wp.DeleteQRCode(ctx, res.CodeID)
		if err != nil {
			log.Fatalf("%+v", err)
		}

		b, _ = json.MarshalIndent(info, "", "  ")
		log.Println(string(b))
	}
*/
package paypayopa
