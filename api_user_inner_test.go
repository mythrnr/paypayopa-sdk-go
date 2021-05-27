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

func Test_createAccountLinkQRCode(t *testing.T) {
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
						"linkQRCodeURL": "https://localhost/linkqr"
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
		qrcode, info, err := createAccountLinkQRCode(ctx, client, &CreateAccountLinkQRCodePayload{})

		t.Log(qrcode, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusCreated, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, qrcode)
		assert.Equal(t, "https://localhost/linkqr", qrcode.LinkQRCodeURL)
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
		qrcode, info, err := createAccountLinkQRCode(ctx, client, &CreateAccountLinkQRCodePayload{})

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
		qrcode, info, err := createAccountLinkQRCode(ctx, client, &CreateAccountLinkQRCodePayload{})

		t.Log(qrcode, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, qrcode)
	})
}

func Test_getUserAuthorizationStatus(t *testing.T) {
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
						"userAuthorizationId": "test-user-authorization-id",
						"referenceIds": [],
						"status": "ACTIVE",
						"scopes": [
							"continuous_payment"
						],
						"expireAt": 0,
						"issuedAt": 0
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
		status, info, err := getUserAuthorizationStatus(ctx, client, "test-user-authorization-id")

		t.Log(status, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, status)
		assert.Equal(t, "test-user-authorization-id", status.UserAuthorizationID)
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
		qrcode, info, err := getUserAuthorizationStatus(ctx, client, "test-user-authorization-id")

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
		qrcode, info, err := getUserAuthorizationStatus(ctx, client, "test-user-authorization-id")

		t.Log(qrcode, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, qrcode)
	})
}

func Test_unlinkUser(t *testing.T) {
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
		info, err := unlinkUser(ctx, client, "user-authorization-id")

		t.Log(info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())
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
		info, err := unlinkUser(ctx, client, "user-authorization-id")

		t.Log(info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INTERNAL_SERVER_ERROR", info.Code)
		assert.Equal(t, "Something went wrong on PayPay service side", info.Message)
		assert.Equal(t, "08101000", info.CodeID)
		assert.Equal(t, http.StatusInternalServerError, info.StatusCode)
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
		info, err := unlinkUser(ctx, client, "user-authorization-id")

		t.Log(info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
	})
}

func Test_getMaskedUserProfile(t *testing.T) {
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
						"phoneNumber": "*******1234"
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
		profile, info, err := getMaskedUserProfile(ctx, client, "test-user-authorization-id")

		t.Log(profile, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, profile)
		assert.Equal(t, "*******1234", profile.PhoneNumber)
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
		profile, info, err := getMaskedUserProfile(ctx, client, "test-user-authorization-id")

		t.Log(profile, info, err)
		assert.Nil(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INTERNAL_SERVER_ERROR", info.Code)
		assert.Equal(t, "Something went wrong on PayPay service side", info.Message)
		assert.Equal(t, "08101000", info.CodeID)
		assert.Equal(t, http.StatusInternalServerError, info.StatusCode)
		assert.False(t, info.Success())

		assert.Nil(t, profile)
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
		qrcode, info, err := getMaskedUserProfile(ctx, client, "test-user-authorization-id")

		t.Log(qrcode, info, err)
		assert.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, qrcode)
	})
}
