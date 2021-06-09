package internal

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Signer(t *testing.T) {
	t.Parallel()

	t.Run("Empty request body", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/test",
			nil,
		)

		require.Nil(t, err)

		a, err := NewSigner("API_KEY", "API_KEY_SECRET", req)

		require.Nil(t, err)

		assert.Equal(t, a.apiKey, "API_KEY")
		assert.Equal(t, a.apiKeySecret, "API_KEY_SECRET")
		assert.Empty(t, a.body)
		assert.NotZero(t, a.epoch)
		assert.Equal(t, http.MethodPost, a.method)
		assert.Len(t, a.nonce, recommendedNonceLen)
		assert.Equal(t, "/test", a.uri)
		assert.Equal(t, contentTypeEmpty, a.ContentType())
	})

	t.Run("Has request body", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodPost,
			"/test",
			bytes.NewBufferString(`{ "test": "value" }`),
		)

		require.Nil(t, err)

		a, err := NewSigner("API_KEY", "API_KEY_SECRET", req)

		require.Nil(t, err)

		assert.Equal(t, a.apiKey, "API_KEY")
		assert.Equal(t, a.apiKeySecret, "API_KEY_SECRET")
		assert.Equal(t, []byte(`{ "test": "value" }`), a.body)
		assert.NotZero(t, a.epoch)
		assert.Equal(t, http.MethodPost, a.method)
		assert.Equal(t, "/test", a.uri)
		assert.Len(t, a.nonce, recommendedNonceLen)
		assert.Equal(t, contentType, a.ContentType())
	})
}
