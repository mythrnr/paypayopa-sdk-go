package paypayopa

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/mythrnr/paypayopa-sdk-go/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_createPaymentAuthorization(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		rt := newTestRoundTripper(http.StatusOK, `{
					"resultInfo": {
						"code": "SUCCESS",
						"message": "Success",
						"codeId": "08100001"
					},
					"data": {
						"paymentId": "string",
						"status": "string",
						"acceptedAt": 0,
						"refunds": {
							"data": [
								{
									"status": "string",
									"acceptedAt": 0,
									"merchantRefundId": "string",
									"paymentId": "string",
									"amount": {
										"amount": 0,
										"currency": "JPY"
									},
									"requestedAt": 0,
									"reason": "string"
								}
							]
						},
						"captures": {
							"data": [
								{
									"acceptedAt": 0,
									"merchantCaptureId": "string",
									"amount": {
										"amount": 0,
										"currency": "JPY"
									},
									"orderDescription": "string",
									"requestedAt": 0,
									"expiresAt": null,
									"status": "string"
								}
							]
						},
						"revert": {
							"acceptedAt": 0,
							"merchantRevertId": "string",
							"requestedAt": 0,
							"reason": "string"
						},
						"merchantPaymentId": "test-merchant-payment-id",
						"userAuthorizationId": "string",
						"amount": {
							"amount": 0,
							"currency": "JPY"
						},
						"requestedAt": 0,
						"expiresAt": null,
						"storeId": "string",
						"terminalId": "string",
						"orderReceiptNumber": "string",
						"orderDescription": "string",
						"orderItems": [
							{
								"name": "string",
								"category": "string",
								"quantity": 1,
								"productId": "string",
								"unitPrice": {
									"amount": 0,
									"currency": "JPY"
								}
							}
						],
						"metadata": {}
					}
				}`)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			newTestClient(rt),
		)

		ctx := context.Background()
		pay, info, err := createPaymentAuthorization(
			ctx, client,
			new(CreatePaymentAuthorizationPayload))

		t.Log(pay, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, pay)
		assert.Equal(t, "test-merchant-payment-id", pay.MerchantPaymentID)
	})

	t.Run("Invalid parameters received", func(t *testing.T) {
		t.Parallel()

		rt := newTestRoundTripper(http.StatusBadRequest, `{
					"resultInfo": {
						"code": "INVALID_PARAMS",
						"message": "Invalid parameters received",
						"codeId": "00200004"
					}
				}`)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			newTestClient(rt),
		)

		ctx := context.Background()
		pay, info, err := createPaymentAuthorization(
			ctx, client,
			new(CreatePaymentAuthorizationPayload))

		t.Log(pay, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INVALID_PARAMS", info.Code)
		assert.Equal(t, "Invalid parameters received", info.Message)
		assert.Equal(t, "00200004", info.CodeID)
		assert.Equal(t, http.StatusBadRequest, info.StatusCode)
		assert.False(t, info.Success())

		assert.Nil(t, pay)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		expected := errors.New("RoundTrip error")
		rt := new(mocks.RoundTripper)
		rt.On("RoundTrip", mock.Anything).
			Return(nil, expected)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			newTestClient(rt),
		)

		ctx := context.Background()
		pay, info, err := createPaymentAuthorization(
			ctx, client,
			new(CreatePaymentAuthorizationPayload))

		t.Log(pay, info, err)
		require.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}

func Test_capturePaymentAuthorization(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		rt := newTestRoundTripper(http.StatusOK, `{
					"resultInfo": {
						"code": "SUCCESS",
						"message": "Success",
						"codeId": "08100001"
					},
					"data": {
						"paymentId": "string",
						"status": "string",
						"acceptedAt": 0,
						"refunds": {
							"data": [
								{
									"status": "string",
									"acceptedAt": 0,
									"merchantRefundId": "string",
									"paymentId": "string",
									"amount": {
										"amount": 0,
										"currency": "JPY"
									},
									"requestedAt": 0,
									"reason": "string"
								}
							]
						},
						"captures": {
							"data": [
							{
								"acceptedAt": 0,
								"merchantCaptureId": "string",
								"amount": {
								"amount": 0,
								"currency": "JPY"
								},
								"orderDescription": "string",
								"requestedAt": 0,
								"status": "string"
							}
							]
						},
						"merchantPaymentId": "test-merchant-payment-id",
						"userAuthorizationId": "string",
						"amount": {
							"amount": 0,
							"currency": "JPY"
						},
						"requestedAt": 0,
						"expiresAt": null,
						"storeId": "string",
						"terminalId": "string",
						"orderReceiptNumber": "string",
						"orderDescription": "string",
						"orderItems": [
							{
								"name": "string",
								"category": "string",
								"quantity": 1,
								"productId": "string",
								"unitPrice": {
									"amount": 0,
									"currency": "JPY"
								}
							}
						],
						"metadata": {},
						"assumeMerchant": "string"
					}
				}`)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			newTestClient(rt),
		)

		ctx := context.Background()
		pay, info, err := capturePaymentAuthorization(
			ctx, client,
			new(CapturePaymentAuthorizationPayload))

		t.Log(pay, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, pay)
		assert.Equal(t, "test-merchant-payment-id", pay.MerchantPaymentID)
	})

	t.Run("User confirmation required", func(t *testing.T) {
		t.Parallel()

		rt := newTestRoundTripper(
			http.StatusAccepted,
			//nolint:lll
			`{
				"resultInfo": {
					"code": "USER_CONFIRMATION_REQUIRED",
					"message": "User confirmation required as requested amount is above allowed limit",
					"codeId": "08300103"
				}
			}`,
		)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			newTestClient(rt),
		)

		ctx := context.Background()
		pay, info, err := capturePaymentAuthorization(
			ctx, client,
			new(CapturePaymentAuthorizationPayload))

		t.Log(pay, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "USER_CONFIRMATION_REQUIRED", info.Code)
		assert.Equal(t, "User confirmation required as requested amount is "+
			"above allowed limit", info.Message)
		assert.Equal(t, "08300103", info.CodeID)
		assert.Equal(t, http.StatusAccepted, info.StatusCode)
		assert.True(t, info.Success())

		assert.Nil(t, pay)
	})

	t.Run("Cannot capture already captured acquiring order", func(t *testing.T) {
		t.Parallel()

		rt := newTestRoundTripper(http.StatusBadRequest, `{
					"resultInfo": {
						"code": "ALREADY_CAPTURED",
						"message": "Cannot capture already captured acquiring order",
						"codeId": "00200039"
					}
				}`)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			newTestClient(rt),
		)

		ctx := context.Background()
		pay, info, err := capturePaymentAuthorization(
			ctx, client,
			new(CapturePaymentAuthorizationPayload))

		t.Log(pay, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "ALREADY_CAPTURED", info.Code)
		assert.Equal(t, "Cannot capture already captured acquiring order", info.Message)
		assert.Equal(t, "00200039", info.CodeID)
		assert.Equal(t, http.StatusBadRequest, info.StatusCode)
		assert.False(t, info.Success())

		assert.Nil(t, pay)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		expected := errors.New("RoundTrip error")
		rt := new(mocks.RoundTripper)
		rt.On("RoundTrip", mock.Anything).
			Return(nil, expected)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			newTestClient(rt),
		)

		ctx := context.Background()
		pay, info, err := capturePaymentAuthorization(
			ctx, client,
			new(CapturePaymentAuthorizationPayload))

		t.Log(pay, info, err)
		require.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}

func Test_revertPaymentAuthorization(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		rt := newTestRoundTripper(http.StatusOK, `{
					"resultInfo": {
						"code": "SUCCESS",
						"message": "Success",
						"codeId": "08100001"
					},
					"data": {
						"status": "string",
						"acceptedAt": 0,
						"paymentId": "test-payment-id",
						"requestedAt": 0,
						"reason": "string"
					}
				}`)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			newTestClient(rt),
		)

		ctx := context.Background()
		pay, info, err := revertPaymentAuthorization(
			ctx, client,
			new(RevertPaymentAuthorizationPayload))

		t.Log(pay, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, pay)
		assert.Equal(t, "test-payment-id", pay.PaymentID)
	})

	t.Run("Order is not cancelable", func(t *testing.T) {
		t.Parallel()

		rt := newTestRoundTripper(http.StatusBadRequest, `{
					"resultInfo": {
						"code": "ORDER_NOT_CANCELABLE",
						"message": "Order is not cancelable",
						"codeId": "00200042"
					}
				}`)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			newTestClient(rt),
		)

		ctx := context.Background()
		pay, info, err := revertPaymentAuthorization(
			ctx, client,
			new(RevertPaymentAuthorizationPayload))

		t.Log(pay, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "ORDER_NOT_CANCELABLE", info.Code)
		assert.Equal(t, "Order is not cancelable", info.Message)
		assert.Equal(t, "00200042", info.CodeID)
		assert.Equal(t, http.StatusBadRequest, info.StatusCode)
		assert.False(t, info.Success())

		assert.Nil(t, pay)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		expected := errors.New("RoundTrip error")
		rt := new(mocks.RoundTripper)
		rt.On("RoundTrip", mock.Anything).
			Return(nil, expected)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			newTestClient(rt),
		)

		ctx := context.Background()
		pay, info, err := revertPaymentAuthorization(
			ctx, client,
			new(RevertPaymentAuthorizationPayload))

		t.Log(pay, info, err)
		require.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}
