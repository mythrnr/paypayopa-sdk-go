package paypayopa

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_decodeAuthorizationResponseToken(t *testing.T) {
	t.Parallel()

	t.Run("Invalid JWT", func(t *testing.T) {
		t.Parallel()

		creds := NewCredentials(
			EnvSandbox,
			"API_KEY",
			"API_KEY_SECRET",
			"MERCHANT_ID",
		)

		testJWT := "invalid-jwt"
		token, err := decodeAuthorizationResponseToken(creds, testJWT)

		assert.Nil(t, token)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "jwt: invalid")
	})

	t.Run("JWT is expired", func(t *testing.T) {
		t.Parallel()

		creds := NewCredentials(
			EnvSandbox,
			"API_KEY",
			"API_KEY_SECRET",
			"MERCHANT_ID",
		)

		//nolint:lll
		// cspell:disable-next-line
		testJWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNRVJDSEFOVF9JRCIsImlzcyI6InBheXBheS5uZS5qcCIsImV4cCI6MjM1NjcsInJlc3VsdCI6InN1Y2NlZWRlZCIsInByb2ZpbGVJZGVudGlmaWVyIjoiKioqKioqKjU2NzgiLCJub25jZSI6InRoZS1zYW1lLW5vbmNlLWluLXRoZS1yZXF1ZXN0IiwidXNlckF1dGhvcml6YXRpb25JZCI6IlBheVBheSB1c2VyIHJlZmVyZW5jZSBpZCIsInJlZmVyZW5jZUlkIjoibWVyY2hhbnQgdXNlciByZWZlcmVuY2UgaWQifQ.AiSmItA9e-JoDSdiBqCrFjDWJO2X31DZdHCl05KaGZc"
		token, err := decodeAuthorizationResponseToken(creds, testJWT)

		assert.Nil(t, token)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "token is expired")
	})

	t.Run("JWT has invalid aud", func(t *testing.T) {
		t.Parallel()

		creds := NewCredentials(
			EnvSandbox,
			"API_KEY",
			"API_KEY_SECRET",
			"MERCHANT_ID",
		)

		//nolint:lll
		// cspell:disable-next-line
		testJWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJJTlZBTElEX01FUkNIQU5UX0lEIiwiaXNzIjoicGF5cGF5Lm5lLmpwIiwiZXhwIjo0MTAyNDEyNDAwLCJyZXN1bHQiOiJzdWNjZWVkZWQiLCJwcm9maWxlSWRlbnRpZmllciI6IioqKioqKio1Njc4Iiwibm9uY2UiOiJ0aGUtc2FtZS1ub25jZS1pbi10aGUtcmVxdWVzdCIsInVzZXJBdXRob3JpemF0aW9uSWQiOiJQYXlQYXkgdXNlciByZWZlcmVuY2UgaWQiLCJyZWZlcmVuY2VJZCI6Im1lcmNoYW50IHVzZXIgcmVmZXJlbmNlIGlkIn0.ZSXzBhApBZoQ9811z4M1mfQVR40JxfM_uiFZvh_JLKw"
		token, err := decodeAuthorizationResponseToken(creds, testJWT)

		assert.Nil(t, token)
		assert.ErrorIs(t, err, ErrJWTAudienceNotMatch)
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		creds := NewCredentials(
			EnvSandbox,
			"API_KEY",
			"API_KEY_SECRET",
			"MERCHANT_ID",
		)

		//nolint:lll
		// cspell:disable-next-line
		testJWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNRVJDSEFOVF9JRCIsImlzcyI6InBheXBheS5uZS5qcCIsImV4cCI6NDEwMjQxMjQwMCwicmVzdWx0Ijoic3VjY2VlZGVkIiwicHJvZmlsZUlkZW50aWZpZXIiOiIqKioqKioqNTY3OCIsIm5vbmNlIjoidGhlLXNhbWUtbm9uY2UtaW4tdGhlLXJlcXVlc3QiLCJ1c2VyQXV0aG9yaXphdGlvbklkIjoiUGF5UGF5IHVzZXIgcmVmZXJlbmNlIGlkIiwicmVmZXJlbmNlSWQiOiJtZXJjaGFudCB1c2VyIHJlZmVyZW5jZSBpZCJ9.GSyN-PxiVaq4PzXxXruz75IhnmM7q34JTsGeRrWVabw"
		token, err := decodeAuthorizationResponseToken(creds, testJWT)

		require.NoError(t, err)
		require.NotNil(t, token)

		assert.Equal(t, "MERCHANT_ID", token.Audience)
		assert.Equal(t, "paypay.ne.jp", token.Issuer)
		assert.Equal(t, int64(4102412400), token.ExpiresAt)
		assert.Equal(t, UserAuthorizeResultSucceeded, token.Result)
		assert.Equal(t, "*******5678", token.ProfileIdentifier)
		assert.Equal(t, "the-same-nonce-in-the-request", token.Nonce)
		assert.Equal(t, "PayPay user reference id", token.UserAuthorizationID)
		assert.Equal(t, "merchant user reference id", token.ReferenceID)
	})
}
