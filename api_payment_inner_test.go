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

func Test_createPayment(t *testing.T) {
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
		pay, info, err := createPayment(ctx, client, &CreatePaymentPayload{})

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

	t.Run("Invalid request params", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "INVALID_REQUEST_PARAMS",
						"message": "Invalid request params",
						"codeId": "08100006"
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
		pay, info, err := createPayment(ctx, client, &CreatePaymentPayload{})

		t.Log(pay, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INVALID_REQUEST_PARAMS", info.Code)
		assert.Equal(t, "Invalid request params", info.Message)
		assert.Equal(t, "08100006", info.CodeID)
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
		pay, info, err := createPayment(ctx, client, &CreatePaymentPayload{})

		t.Log(pay, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}

func Test_cancelPayment(t *testing.T) {
	t.Parallel()

	t.Run("Accepted", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusAccepted),
				StatusCode: http.StatusAccepted,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "REQUEST_ACCEPTED",
						"message": "Request accepted",
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
		info, err := cancelPayment(ctx, client, "1")

		t.Log(info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "REQUEST_ACCEPTED", info.Code)
		assert.Equal(t, "Request accepted", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusAccepted, info.StatusCode)
		assert.True(t, info.Success())
	})

	t.Run("Order cannot be reversed", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "ORDER_NOT_REVERSIBLE",
						"message": "Order cannot be reversed",
						"codeId": "00200044"
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
		info, err := cancelPayment(ctx, client, "1")

		t.Log(info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "ORDER_NOT_REVERSIBLE", info.Code)
		assert.Equal(t, "Order cannot be reversed", info.Message)
		assert.Equal(t, "00200044", info.CodeID)
		assert.Equal(t, http.StatusBadRequest, info.StatusCode)
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
		info, err := cancelPayment(ctx, client, "1")

		t.Log(info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
	})
}

func Test_getPaymentDetails(t *testing.T) {
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
		pay, info, err := getPaymentDetails(ctx, client, "test-merchant-payment-id")

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

	t.Run("Dynamic QR payment not found", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
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
		pay, info, err := getPaymentDetails(ctx, client, "test-merchant-payment-id")

		t.Log(pay, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "DYNAMIC_QR_PAYMENT_NOT_FOUND", info.Code)
		assert.Equal(t, "Dynamic QR payment not found", info.Message)
		assert.Equal(t, "01652075", info.CodeID)
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
		pay, info, err := getPaymentDetails(ctx, client, "test-merchant-payment-id")

		t.Log(pay, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}

func Test_createContinuousPayment(t *testing.T) {
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
		pay, info, err := createContinuousPayment(
			ctx, client,
			&CreateContinuousPaymentPayload{})

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

	t.Run("Request order not found", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "REQUEST_ORDER_NOT_FOUND",
						"message": "Request order not found",
						"codeId": "08300005"
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
		pay, info, err := createContinuousPayment(
			ctx, client,
			&CreateContinuousPaymentPayload{})

		t.Log(pay, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "REQUEST_ORDER_NOT_FOUND", info.Code)
		assert.Equal(t, "Request order not found", info.Message)
		assert.Equal(t, "08300005", info.CodeID)
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
		pay, info, err := createContinuousPayment(
			ctx, client,
			&CreateContinuousPaymentPayload{})

		t.Log(pay, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, pay)
	})
}

func Test_consultExpectedCashbackInfo(t *testing.T) {
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
						"campaignMessage": "100円相当戻ってくる！"
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
		cashback, info, err := consultExpectedCashbackInfo(
			ctx, client,
			&ConsultExpectedCashbackInfoPayload{})

		t.Log(cashback, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, cashback)
		assert.Equal(t, "100円相当戻ってくる！", cashback.Campaignmessage)
	})

	t.Run("Success (EN)", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Run(func(args mock.Arguments) {
				req, _ := args[0].(*http.Request)
				assert.Equal(t, "EN", req.Header.Get(headerNameLang))
			}).
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
						"campaignMessage": "Get 100 Yen cashback"
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
		cashback, info, err := consultExpectedCashbackInfo(
			ctx, client,
			&ConsultExpectedCashbackInfoPayload{Lang: LangEN})

		t.Log(cashback, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, cashback)
		assert.Equal(t, "Get 100 Yen cashback", cashback.Campaignmessage)
	})

	t.Run("Something went wrong on PayPay service side", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusInternalServerError),
				StatusCode: http.StatusInternalServerError,
				Body: ioutil.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "INTERNAL_SERVER_ERROR",
						"message": "Something went wrong on PayPay service side",
						"codeId": "08101000"
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
		cashback, info, err := consultExpectedCashbackInfo(
			ctx, client,
			&ConsultExpectedCashbackInfoPayload{})

		t.Log(cashback, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INTERNAL_SERVER_ERROR", info.Code)
		assert.Equal(t, "Something went wrong on PayPay service side", info.Message)
		assert.Equal(t, "08101000", info.CodeID)
		assert.Equal(t, http.StatusInternalServerError, info.StatusCode)
		assert.False(t, info.Success())

		assert.Nil(t, cashback)
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
		cashback, info, err := consultExpectedCashbackInfo(
			ctx, client,
			&ConsultExpectedCashbackInfoPayload{})

		t.Log(cashback, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, cashback)
	})
}
