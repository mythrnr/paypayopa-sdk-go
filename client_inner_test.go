package paypayopa

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/mythrnr/paypayopa-sdk-go/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_newClient(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		client := newClient(NewCredentials(
			EnvSandbox,
			"API_KEY",
			"API_KEY_SECRET",
			"MERCHANT_ID",
		))

		t.Log(client)
	})

	assert.Panics(t, func() {
		client := newClient(nil)

		t.Log(client)
	})
}

func Test_newClientWithHTTPClient(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{},
		)

		assert.Equal(t, timeout, client.http.Timeout)
	})

	assert.PanicsWithValue(t, "*http.Client must not be nil", func() {
		newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			nil,
		)
	})

	assert.PanicsWithValue(t, "*Credentials must not be nil", func() {
		newClientWithHTTPClient(nil, &http.Client{})
	})
}

func Test_opaClient_Request(t *testing.T) {
	t.Parallel()

	t.Run("Invalid request body", func(t *testing.T) {
		t.Parallel()

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{},
		)

		expected := errors.New("marshal error")
		marshaler := &mocks.Marshaler{}
		marshaler.On("MarshalJSON").Return(nil, expected)

		req, err := client.Request(
			context.Background(),
			http.MethodPost,
			"/test",
			marshaler,
		)

		assert.Nil(t, req)
		assert.ErrorIs(t, err, expected)
	})

	t.Run("Failed to create request", func(t *testing.T) {
		t.Parallel()

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{},
		)

		req, err := client.Request(
			context.Background(),
			http.MethodPost,
			"%%%%-invalid-url-%%%%",
			map[string]interface{}{},
		)

		actual := &url.Error{}

		assert.Nil(t, req)
		assert.ErrorAs(t, err, &actual)
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{},
		)

		req, err := client.Request(
			context.Background(),
			http.MethodPost,
			"/test",
			map[string]interface{}{},
		)

		require.NoError(t, err)
		assert.NotNil(t, req)
	})
}

func Test_opaClient_Do(t *testing.T) {
	t.Parallel()

	t.Run("Timeout", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(nil, context.DeadlineExceeded)

		client := newClientWithHTTPClient(
			NewCredentials(
				EnvSandbox,
				"API_KEY",
				"API_KEY_SECRET",
				"MERCHANT_ID",
			),
			&http.Client{Transport: rt},
		)

		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/test",
			nil,
		)

		res := map[string]interface{}{}
		info, err := client.Do(req, res)

		assert.Nil(t, info)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})

	t.Run("Unknown error on RoundTrip", func(t *testing.T) {
		t.Parallel()

		expected := errors.New("unknown error")
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

		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/test",
			nil,
		)

		res := map[string]interface{}{}
		info, err := client.Do(req, res)

		assert.Nil(t, info)
		assert.ErrorIs(t, err, expected)
	})

	t.Run("Failed to read response body", func(t *testing.T) {
		t.Parallel()

		expected := errors.New("read error")
		reader := &mocks.ReadCloser{}
		reader.
			On("Read", mock.Anything).
			Return(0, expected).
			On("Close").
			Return(nil)

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusAccepted),
				StatusCode: http.StatusAccepted,
				Body:       reader,
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

		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/test",
			nil,
		)

		res := map[string]interface{}{}
		info, err := client.Do(req, res)

		assert.Nil(t, info)
		assert.ErrorIs(t, err, expected)
	})

	t.Run("Failed to unmarshal body", func(t *testing.T) {
		t.Parallel()

		rt := &mocks.RoundTripper{}
		rt.On("RoundTrip", mock.Anything).
			Return(&http.Response{
				Status:     http.StatusText(http.StatusAccepted),
				StatusCode: http.StatusAccepted,
				Body:       io.NopCloser(bytes.NewBufferString(`invalid-json`)),
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

		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/test",
			nil,
		)

		res := map[string]interface{}{}
		info, err := client.Do(req, res)
		actual := &json.SyntaxError{}

		assert.Nil(t, info)
		assert.ErrorAs(t, err, &actual)
	})

	t.Run("response.data is empty", func(t *testing.T) {
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
						"codeId": "success-code-id"
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

		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodDelete,
			"/test",
			nil,
		)

		res := map[string]interface{}{}
		info, err := client.Do(req, res)

		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "success-code-id", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())
	})

	t.Run("target to bind is nil", func(t *testing.T) {
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
						"codeId": "success-code-id"
					},
					"data": {
						"key": "value"
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

		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/test",
			nil,
		)

		info, err := client.Do(req, nil)

		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "success-code-id", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())
	})

	t.Run("Failed to unmarshal response.data", func(t *testing.T) {
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
						"codeId": "success-code-id"
					},
					"data": {
						"key": "value"
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

		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/test",
			nil,
		)

		res := struct {
			Key int `json:"key"`
		}{}

		info, err := client.Do(req, &res)
		actual := &json.UnmarshalTypeError{}

		assert.Nil(t, info)
		assert.ErrorAs(t, err, &actual)
	})

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
						"codeId": "success-code-id"
					},
					"data": {
						"key": "value"
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

		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/test",
			io.NopCloser(bytes.NewBufferString(`{ "test": "value" }`)),
		)

		res := struct {
			Key string `json:"key"`
		}{}

		info, err := client.Do(req, &res)

		require.NoError(t, err)

		require.NotNil(t, info)
		assert.Equal(t, "SUCCESS", info.Code)
		assert.Equal(t, "Success", info.Message)
		assert.Equal(t, "success-code-id", info.CodeID)
		assert.Equal(t, http.StatusOK, info.StatusCode)
		assert.True(t, info.Success())

		assert.Equal(t, "value", res.Key)
	})
}

func Test_ctx_timeout(t *testing.T) {
	t.Parallel()

	t.Run("timeout is set", func(t *testing.T) {
		t.Parallel()

		expected := time.Second

		ctx := ctxWithTimeout(context.Background(), expected)
		actual := getTimeout(ctx)

		assert.Equal(t, expected, actual)
	})

	t.Run("timeout is not set", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		actual := getTimeout(ctx)

		assert.Equal(t, timeout, actual)
	})
}
