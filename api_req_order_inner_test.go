package paypayopa

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/mythrnr/paypayopa-sdk-go/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_createPendingPayment(t *testing.T) {
	t.Parallel()

	t.Run("Created", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusCreated),
				StatusCode: http.StatusCreated,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "SUCCESS",
						"message": "Success",
						"codeId": "08100001"
					},
					"data": {
						"merchantPaymentId": "test-merchant-payment-id",
						"userAuthorizationId": "string",
						"amount": {
							"amount": 0,
							"currency": "JPY"
						},
						"requestedAt": 0,
						"expiryDate": null,
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
		qrcode, info, err := createPendingPayment(ctx, client, &CreatePendingPaymentPayload{})

		t.Log(qrcode, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusCreated, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, qrcode)
		assert.Equal(t, "test-merchant-payment-id", qrcode.MerchantPaymentID)
	})

	t.Run("Request order with same payment ID exists", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "DUPLICATE_REQUEST_ORDER",
						"message": "Request order with same payment ID exists",
						"codeId": ""
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
		qrcode, info, err := createPendingPayment(ctx, client, &CreatePendingPaymentPayload{})

		t.Log(qrcode, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "DUPLICATE_REQUEST_ORDER", info.Code)
		assert.Equal(t, "Request order with same payment ID exists", info.Message)
		assert.Equal(t, "", info.CodeID)
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
		qrcode, info, err := createPendingPayment(ctx, client, &CreatePendingPaymentPayload{})

		t.Log(qrcode, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, qrcode)
	})
}

func Test_cancelPendingOrder(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusAccepted),
				StatusCode: http.StatusAccepted,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
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
		info, err := cancelPendingOrder(ctx, client, "1")

		t.Log(info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusAccepted, info.StatusCode)
		assert.True(t, info.Success())
	})

	t.Run("Request order not in valid state", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusNotFound),
				StatusCode: http.StatusNotFound,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "REQUEST_ORDER_NOT_FOUND",
						"message": "Request order not in valid state",
						"codeId": ""
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
		info, err := cancelPendingOrder(ctx, client, "1")

		t.Log(info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "REQUEST_ORDER_NOT_FOUND", info.Code)
		assert.Equal(t, "Request order not in valid state", info.Message)
		assert.Equal(t, "", info.CodeID)
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
		info, err := cancelPendingOrder(ctx, client, "1")

		t.Log(info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
	})
}

func Test_getRequestedPaymentDetails(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusOK),
				StatusCode: http.StatusOK,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
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
						"merchantPaymentId": "test-merchant-payment-id",
						"userAuthorizationId": "string",
						"amount": {
							"amount": 0,
							"currency": "JPY"
						},
						"requestedAt": 0,
						"expiryDate": null,
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
		pay, info, err := getRequestedPaymentDetails(ctx, client, "test-merchant-payment-id")

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

	t.Run("The set parameter is invalid", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "INVALID_PARAMS",
						"message": "The set parameter is invalid",
						"codeId": ""
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
		pay, info, err := getRequestedPaymentDetails(ctx, client, "test-merchant-payment-id")

		t.Log(pay, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INVALID_PARAMS", info.Code)
		assert.Equal(t, "The set parameter is invalid", info.Message)
		assert.Equal(t, "", info.CodeID)
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
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		pay, info, err := getRequestedPaymentDetails(ctx, client, "test-merchant-payment-id")

		t.Log(pay, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}
