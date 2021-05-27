package paypayopa

import (
	"bytes"
	"context"
	"errors"
	"io"
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

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusOK),
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewBufferString(`{
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
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := createPaymentAuthorization(
			ctx, client,
			&CreatePaymentAuthorizationPayload{})

		t.Log(pay, info, err)
		assert.Nil(t, err)

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

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "INVALID_PARAMS",
						"message": "Invalid parameters received",
						"codeId": "00200004"
					}
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := createPaymentAuthorization(
			ctx, client,
			&CreatePaymentAuthorizationPayload{})

		t.Log(pay, info, err)
		assert.Nil(t, err)

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
		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(nil, expected)

		client := newClientWithHTTPClient(
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := createPaymentAuthorization(
			ctx, client,
			&CreatePaymentAuthorizationPayload{})

		t.Log(pay, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}

func Test_capturePaymentAuthorization(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusOK),
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewBufferString(`{
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
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := capturePaymentAuthorization(
			ctx, client,
			&CapturePaymentAuthorizationPayload{})

		t.Log(pay, info, err)
		assert.Nil(t, err)

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

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusAccepted),
				StatusCode: http.StatusAccepted,
				// nolint:lll
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "USER_CONFIRMATION_REQUIRED",
						"message": "User confirmation required as requested amount is above allowed limit",
						"codeId": "08300103"
					}
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := capturePaymentAuthorization(
			ctx, client,
			&CapturePaymentAuthorizationPayload{})

		t.Log(pay, info, err)
		assert.Nil(t, err)

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

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "ALREADY_CAPTURED",
						"message": "Cannot capture already captured acquiring order",
						"codeId": "00200039"
					}
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := capturePaymentAuthorization(
			ctx, client,
			&CapturePaymentAuthorizationPayload{})

		t.Log(pay, info, err)
		assert.Nil(t, err)

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
		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(nil, expected)

		client := newClientWithHTTPClient(
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := capturePaymentAuthorization(
			ctx, client,
			&CapturePaymentAuthorizationPayload{})

		t.Log(pay, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}

func Test_revertPaymentAuthorization(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusOK),
				StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewBufferString(`{
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
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := revertPaymentAuthorization(
			ctx, client,
			&RevertPaymentAuthorizationPayload{})

		t.Log(pay, info, err)
		assert.Nil(t, err)

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

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "ORDER_NOT_CANCELABLE",
						"message": "Order is not cancelable",
						"codeId": "00200042"
					}
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := revertPaymentAuthorization(
			ctx, client,
			&RevertPaymentAuthorizationPayload{})

		t.Log(pay, info, err)
		assert.Nil(t, err)

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
		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(nil, expected)

		client := newClientWithHTTPClient(
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := revertPaymentAuthorization(
			ctx, client,
			&RevertPaymentAuthorizationPayload{})

		t.Log(pay, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}
