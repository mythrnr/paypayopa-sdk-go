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

func Test_refundPayment(t *testing.T) {
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
						"status": "string",
						"acceptedAt": 0,
						"merchantRefundId": "test-refund-id",
						"paymentId": "string",
						"amount": {
							"amount": 0,
							"currency": "JPY"
						},
						"requestedAt": 0,
						"reason": "string",
						"assumeMerchant": "string"
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
		refund, info, err := refundPayment(ctx, client, &RefundPaymentPayload{})

		t.Log(refund, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusCreated, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, refund)
		assert.Equal(t, "test-refund-id", refund.MerchantRefundID)
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
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		ctx := context.Background()
		refund, info, err := refundPayment(ctx, client, &RefundPaymentPayload{})

		t.Log(refund, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "INVALID_PARAMS", info.Code)
		assert.Equal(t, "Invalid parameters received", info.Message)
		assert.Equal(t, "00200004", info.CodeID)
		assert.Equal(t, http.StatusBadRequest, info.StatusCode)
		assert.False(t, info.Success())

		assert.Nil(t, refund)
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
		refund, info, err := refundPayment(ctx, client, &RefundPaymentPayload{})

		t.Log(refund, info, err)
		require.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, refund)
	})
}

func Test_getRefundDetails(t *testing.T) {
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
						"status": "string",
						"acceptedAt": 0,
						"merchantRefundId": "test-refund-id",
						"paymentId": "string",
						"amount": {
							"amount": 0,
							"currency": "JPY"
						},
						"requestedAt": 0,
						"reason": "string",
						"assumeMerchant": "string"
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
		refund, info, err := getRefundDetails(ctx, client, "test-refund-id")

		t.Log(refund, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "08100001", info.CodeID)
		assert.Equal(t, http.StatusCreated, info.StatusCode)
		assert.True(t, info.Success())

		require.NotNil(t, refund)
		assert.Equal(t, "test-refund-id", refund.MerchantRefundID)
	})

	t.Run("Transaction failed", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusInternalServerError),
				StatusCode: http.StatusInternalServerError,
				Body: io.NopCloser(bytes.NewBufferString(`{
					"resultInfo": {
						"code": "TRANSACTION_FAILED",
						"message": "Transaction failed",
						"codeId": "00200002"
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
		refund, info, err := getRefundDetails(ctx, client, "test-refund-id")

		t.Log(refund, info, err)
		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "TRANSACTION_FAILED", info.Code)
		assert.Equal(t, "Transaction failed", info.Message)
		assert.Equal(t, "00200002", info.CodeID)
		assert.Equal(t, http.StatusInternalServerError, info.StatusCode)
		assert.False(t, info.Success())

		assert.Nil(t, refund)
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
		refund, info, err := getRefundDetails(ctx, client, "test-refund-id")

		t.Log(refund, info, err)
		require.ErrorIs(t, err, expected)
		assert.Nil(t, info)
		assert.Nil(t, refund)
	})
}
