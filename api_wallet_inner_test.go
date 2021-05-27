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

func Test_checkUserWalletBalance(t *testing.T) {
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
						"hasEnoughBalance": true
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
		wallet, info, err := checkUserWalletBalance(ctx, client, &CheckUserWalletBalancePayload{})

		t.Log(wallet, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, wallet)
		assert.True(t, wallet.HasEnoughBalance)
	})

	t.Run("Invalid request params", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "INVALID_REQUEST_PARAMS",
						"message": "Invalid request params",
						"codeId": "08100006"
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
		wallet, info, err := checkUserWalletBalance(ctx, client, &CheckUserWalletBalancePayload{})

		t.Log(wallet, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INVALID_REQUEST_PARAMS", info.Code)
		assert.Equal(t, "Invalid request params", info.Message)
		assert.Equal(t, "08100006", info.CodeID)
		assert.Equal(t, http.StatusBadRequest, info.StatusCode)
		assert.False(t, info.Success())

		assert.Nil(t, wallet)
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
		wallet, info, err := checkUserWalletBalance(ctx, client, &CheckUserWalletBalancePayload{})

		t.Log(wallet, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, wallet)
	})
}

func Test_getUserWalletBalance(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Run(func(args mock.Arguments) {
				req, _ := args[0].(*http.Request)
				assert.Equal(t,
					string(EnvSandbox)+
						"/v6/wallet/balance?"+
						"currency=JPY"+
						"&productType=REAL_INVESTMENT"+
						"&userAuthorizationId=user-authorization-id",
					req.URL.String(),
				)
			}).
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
						"userAuthorizationId": "user-authorization-id",
						"totalBalance": {
							"amount": 0,
							"currency": "JPY"
						},
						"preference": {
							"useCashback": false,
							"cashbackAutoInvestment": false
						}
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
		wallet, info, err := getUserWalletBalance(
			ctx,
			client,
			&GetUserWalletBalancePayload{
				Currency:            CurrencyJPY,
				ProductType:         ProductTypeRealInvestment,
				UserAuthorizationID: "user-authorization-id",
			})

		t.Log(wallet, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, wallet)
		assert.Equal(t, "user-authorization-id", wallet.UserAuthorizationID)
	})

	t.Run("Something went wrong on PayPay service side", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Run(func(args mock.Arguments) {
				req, _ := args[0].(*http.Request)
				assert.Equal(t,
					string(EnvSandbox)+
						"/v6/wallet/balance?"+
						"currency="+
						"&productType="+
						"&userAuthorizationId=",
					req.URL.String(),
				)
			}).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusInternalServerError),
				StatusCode: http.StatusInternalServerError,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "INTERNAL_SERVER_ERROR",
						"message": "Something went wrong on PayPay service side",
						"codeId": "08101000"
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
		wallet, info, err := getUserWalletBalance(ctx, client, &GetUserWalletBalancePayload{})

		t.Log(wallet, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INTERNAL_SERVER_ERROR", info.Code)
		assert.Equal(t, "Something went wrong on PayPay service side", info.Message)
		assert.Equal(t, "08101000", info.CodeID)
		assert.Equal(t, http.StatusInternalServerError, info.StatusCode)
		assert.False(t, info.Success())

		assert.Nil(t, wallet)
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
		wallet, info, err := getUserWalletBalance(ctx, client, &GetUserWalletBalancePayload{})

		t.Log(wallet, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, wallet)
	})
}

func Test_createTopupQRCode(t *testing.T) {
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
						"codeId": "string",
						"url": "string",
						"status": "string",
						"merchantTopUpId": "string",
						"userAuthorizationId": "test-user-authorization-id",
						"minimumTopUpAmount": {
							"amount": 0,
							"currency": "JPY"
						},
						"metadata": {},
						"expiryDate": 0,
						"codeType": "TOPUP_QR",
						"requestedAt": 0,
						"redirectType": "WEB_LINK",
						"redirectUrl": "string",
						"userAgent": "string"
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
		qrcode, info, err := createTopupQRCode(ctx, client, &CreateTopupQRCodePayload{})

		t.Log(qrcode, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusCreated, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, qrcode)
		assert.Equal(t, "test-user-authorization-id", qrcode.UserAuthorizationID)
	})

	t.Run("Invalid request params", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "INVALID_REQUEST_PARAMS",
						"message": "Invalid request params",
						"codeId": "08100006"
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
		qrcode, info, err := createTopupQRCode(ctx, client, &CreateTopupQRCodePayload{})

		t.Log(qrcode, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INVALID_REQUEST_PARAMS", info.Code)
		assert.Equal(t, "Invalid request params", info.Message)
		assert.Equal(t, "08100006", info.CodeID)
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
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		qrcode, info, err := createTopupQRCode(ctx, client, &CreateTopupQRCodePayload{})

		t.Log(qrcode, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, qrcode)
	})
}

func Test_deleteTopupQRCode(t *testing.T) {
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
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		info, err := deleteTopupQRCode(ctx, client, "1")

		t.Log(info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())
	})

	t.Run("Unauthorized request", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusUnauthorized),
				StatusCode: http.StatusUnauthorized,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "UNAUTHORIZED",
						"message": "Unauthorized request",
						"codeId": "08100016"
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
		info, err := deleteTopupQRCode(ctx, client, "1")

		t.Log(info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "UNAUTHORIZED", info.Code)
		assert.Equal(t, "Unauthorized request", info.Message)
		assert.Equal(t, "08100016", info.CodeID)
		assert.Equal(t, http.StatusUnauthorized, info.StatusCode)
		assert.False(t, info.Success())
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
		info, err := deleteTopupQRCode(ctx, client, "1")

		t.Log(info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
	})
}

func Test_getTopupDetails(t *testing.T) {
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
						"topUpId": "topup-id",
						"merchantTopUpId": "string",
						"userAuthorizationId": "test-user-authorization-id",
						"requestedAt": 0,
						"acceptedAt": 0,
						"expiryDate": 0,
						"status": "CREATED",
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
		qrcode, info, err := getTopupDetails(ctx, client, "topup-id")

		t.Log(qrcode, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, qrcode)
		assert.Equal(t, "topup-id", qrcode.TopupID)
	})

	t.Run("Something went wrong on PayPay service side", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusInternalServerError),
				StatusCode: http.StatusInternalServerError,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "INTERNAL_SERVER_ERROR",
						"message": "Something went wrong on PayPay service side",
						"codeId": "08101000"
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
		qrcode, info, err := getTopupDetails(ctx, client, "topup-id")

		t.Log(qrcode, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INTERNAL_SERVER_ERROR", info.Code)
		assert.Equal(t, "Something went wrong on PayPay service side", info.Message)
		assert.Equal(t, "08101000", info.CodeID)
		assert.Equal(t, http.StatusInternalServerError, info.StatusCode)
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
			NewCredential(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		qrcode, info, err := getTopupDetails(ctx, client, "topup-id")

		t.Log(qrcode, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, qrcode)
	})
}
