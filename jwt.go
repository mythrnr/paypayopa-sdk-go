package paypayopa

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type AuthorizationResponseToken struct {
	Audience            string
	Issuer              string
	ExpiresAt           int64
	Result              UserAuthorizeResult
	ProfileIdentifier   string
	Nonce               string
	UserAuthorizationID string
	ReferenceID         string
}

var ErrJWTAudienceNotMatch = errors.New("jwt: audience not match")

type rawToken struct {
	Result              UserAuthorizeResult `json:"result"`
	ProfileIdentifier   string              `json:"profileIdentifier"`
	Nonce               string              `json:"nonce"`
	UserAuthorizationID string              `json:"userAuthorizationId"`
	ReferenceID         string              `json:"referenceId"`

	Audience string `json:"aud"`
	jwt.RegisteredClaims
}

func decodeAuthorizationResponseToken(
	creds *Credentials,
	token string,
) (*AuthorizationResponseToken, error) {
	claims := rawToken{}

	if _, err := jwt.ParseWithClaims(
		token, &claims,
		func(_ *jwt.Token) (any, error) {
			return []byte(creds.apiKeySecret), nil
		},
	); err != nil {
		return nil, fmt.Errorf("jwt: invalid: %w", err)
	}

	if creds.merchantID != claims.Audience {
		return nil, ErrJWTAudienceNotMatch
	}

	return &AuthorizationResponseToken{
		Audience:            claims.Audience,
		Issuer:              claims.Issuer,
		ExpiresAt:           claims.ExpiresAt.Unix(),
		Result:              claims.Result,
		ProfileIdentifier:   claims.ProfileIdentifier,
		Nonce:               claims.Nonce,
		UserAuthorizationID: claims.UserAuthorizationID,
		ReferenceID:         claims.ReferenceID,
	}, nil
}
