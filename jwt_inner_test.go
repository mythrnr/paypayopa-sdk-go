package paypayopa

import (
	"testing"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// nolint:lll
const testJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNRVJDSEFOVF9JRCIsImlzcyI6InBheXBheS5uZS5qcCIsImV4cCI6MjM1NjcsInJlc3VsdCI6InN1Y2NlZWRlZCIsInByb2ZpbGVJZGVudGlmaWVyIjoiKioqKioqKjU2NzgiLCJub25jZSI6InRoZS1zYW1lLW5vbmNlLWluLXRoZS1yZXF1ZXN0IiwidXNlckF1dGhvcml6YXRpb25JZCI6IlBheVBheSB1c2VyIHJlZmVyZW5jZSBpZCIsInJlZmVyZW5jZUlkIjoibWVyY2hhbnQgdXNlciByZWZlcmVuY2UgaWQifQ.AiSmItA9e-JoDSdiBqCrFjDWJO2X31DZdHCl05KaGZc"

func Test_decodeAuthorizationResponseToken(t *testing.T) {
	t.Parallel()

	creds := NewCredential(
		EnvSandbox,
		"API_KEY",
		"API_KEY_SECRET",
		"MERCHANT_ID",
	)

	token, err := decodeAuthorizationResponseToken(creds, testJWT)

	jwtErr := &jwt.TokenExpiredError{}
	assert.ErrorAs(t, err, &jwtErr)
	require.NotNil(t, token)

	assert.Equal(t, "MERCHANT_ID", token.Audience)
	assert.Equal(t, "paypay.ne.jp", token.Issuer)
	assert.Equal(t, int64(23567), token.ExpiresAt)
	assert.Equal(t, UserAuthorizeResultSucceeded, token.Result)
	assert.Equal(t, "*******5678", token.ProfileIdentifier)
	assert.Equal(t, "the-same-nonce-in-the-request", token.Nonce)
	assert.Equal(t, "PayPay user reference id", token.UserAuthorizationID)
	assert.Equal(t, "merchant user reference id", token.ReferenceID)
}
