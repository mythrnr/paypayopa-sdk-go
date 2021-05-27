package paypayopa

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_authenticate(t *testing.T) {
	t.Parallel()

	a := newAuthenticate("API_KEY", "API_KEY_SECRET")
	assert.Equal(t, a.apiKey, "API_KEY")
	assert.Equal(t, a.apiKeySecret, "API_KEY_SECRET")
	assert.Len(t, a.nonce, recommendedNonceLen)
	assert.NotZero(t, a.epoch)

	assert.Empty(t, a.body)
	assert.Empty(t, a.method)
	assert.Empty(t, a.uri)
	assert.Equal(t, contentTypeEmpty, a.contentType())

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/test",
		bytes.NewBufferString(`{ "test": "value" }`),
	)

	require.Nil(t, err)
	require.Nil(t, a.setRequest(req))

	assert.Equal(t, []byte(`{ "test": "value" }`), a.body)
	assert.Equal(t, http.MethodPost, a.method)
	assert.Equal(t, "/test", a.uri)
	assert.Equal(t, contentType, a.contentType())
}
