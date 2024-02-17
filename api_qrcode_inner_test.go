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

func Test_createQRCode(t *testing.T) {
	t.Parallel()

	t.Run("Created", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusCreated),
				StatusCode: http.StatusCreated,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "SUCCESS",
						"message": "Success",
						"codeId": "08100001"
					},
					"data": {
						"codeId": "generated-qr-code-id",
						"url": "string",
						"deeplink": "string",
						"expiryDate": 0,
						"merchantPaymentId": "string",
						"amount": {
							"amount": 0,
							"currency": "JPY"
						},
						"orderDescription": "string",
						"orderItems": [
							{
								"name": "string",
								"category": "string",
								"quantity": 1,
								"productId": "string",
								"unit_price": {
								"amount": 0,
								"currency": "JPY"
								}
							}
						],
						"metadata": {},
						"codeType": "string",
						"storeInfo": "string",
						"storeId": "string",
						"terminalId": "string",
						"requestedAt": 0,
						"redirectUrl": "string",
						"redirectType": "WEB_LINK",
						"isAuthorization": true,
						"authorizationExpiry": null
					}
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		qrcode, info, err := createQRCode(ctx, client, &CreateQRCodePayload{})

		t.Log(qrcode, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusCreated, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, qrcode)
		assert.Equal(t, "generated-qr-code-id", qrcode.CodeID)
	})

	t.Run("Dynamic QR bad request error", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "DYNAMIC_QR_BAD_REQUEST",
						"message": "Dynamic QR bad request error",
						"codeId": "01650000"
					}
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		qrcode, info, err := createQRCode(ctx, client, &CreateQRCodePayload{})

		t.Log(qrcode, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "DYNAMIC_QR_BAD_REQUEST", info.Code)
		assert.Equal(t, "Dynamic QR bad request error", info.Message)
		assert.Equal(t, "01650000", info.CodeID)
		assert.Equal(t, http.StatusBadRequest, info.StatusCode)
		assert.False(t, info.Success())

		assert.Nil(t, qrcode)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		expected := errors.New("RoundTrip error")
		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(nil, expected)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		qrcode, info, err := createQRCode(ctx, client, &CreateQRCodePayload{})

		t.Log(qrcode, info, err)
		require.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, qrcode)
	})
}

func Test_deleteQRCode(t *testing.T) {
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
					}
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		info, err := deleteQRCode(ctx, client, "test-code-id")

		t.Log(info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())
	})

	t.Run("Dynamic qr code not found", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusNotFound),
				StatusCode: http.StatusNotFound,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "DYNAMIC_QR_NOT_FOUND",
						"message": "Dynamic qr code not found",
						"codeId": "01652072"
					}
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		info, err := deleteQRCode(ctx, client, "test-code-id")

		t.Log(info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "DYNAMIC_QR_NOT_FOUND", info.Code)
		assert.Equal(t, "Dynamic qr code not found", info.Message)
		assert.Equal(t, "01652072", info.CodeID)
		assert.Equal(t, http.StatusNotFound, info.StatusCode)
		assert.False(t, info.Success())
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		expected := errors.New("RoundTrip error")
		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(nil, expected)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		info, err := deleteQRCode(ctx, client, "test-code-id")

		t.Log(info, err)
		require.ErrorIs(t, err, expected)
		assert.Nil(t, info)
	})
}

func Test_getCodePaymentDetails(t *testing.T) {
	t.Parallel()

	t.Run("Created", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusCreated),
				StatusCode: http.StatusCreated,
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
						"merchantPaymentId": "merchant-payment-id",
						"amount": {
							"amount": 0,
							"currency": "JPY"
						},
						"requestedAt": 0,
						"expiresAt": null,
						"canceledAt": null,
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
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := getCodePaymentDetails(ctx, client, "merchant-payment-id")

		t.Log(pay, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusCreated, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, pay)
		assert.Equal(t, "merchant-payment-id", pay.MerchantPaymentID)
	})

	t.Run("Dynamic QR payment not found", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusNotFound),
				StatusCode: http.StatusNotFound,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "DYNAMIC_QR_PAYMENT_NOT_FOUND",
						"message": "Dynamic QR payment not found",
						"codeId": "01652075"
					}
				}`)),
			}, nil)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := getCodePaymentDetails(ctx, client, "not-exist-id")

		t.Log(pay, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "DYNAMIC_QR_PAYMENT_NOT_FOUND", info.Code)
		assert.Equal(t, "Dynamic QR payment not found", info.Message)
		assert.Equal(t, "01652075", info.CodeID)
		assert.Equal(t, http.StatusNotFound, info.StatusCode)
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
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := getCodePaymentDetails(ctx, client, "merchant-payment-id")

		t.Log(pay, info, err)
		require.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}
