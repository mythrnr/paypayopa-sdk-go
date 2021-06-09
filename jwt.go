package paypayopa

import (
	jwt "github.com/dgrijalva/jwt-go/v4"
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

type rawToken struct {
	Result              UserAuthorizeResult `json:"result"`
	ProfileIdentifier   string              `json:"profileIdentifier"`
	Nonce               string              `json:"nonce"`
	UserAuthorizationID string              `json:"userAuthorizationId"`
	ReferenceID         string              `json:"referenceId"`

	jwt.StandardClaims
}

func decodeAuthorizationResponseToken(
	creds *Credentials,
	token string,
) (*AuthorizationResponseToken, error) {
	claims := rawToken{}
	_, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(creds.apiKeySecret), nil
		},
		jwt.WithAudience(creds.merchantID),
	)

	aud := ""
	if 0 < len(claims.Audience) {
		aud = claims.Audience[0]
	}

	return &AuthorizationResponseToken{
		Audience:            aud,
		Issuer:              claims.Issuer,
		ExpiresAt:           claims.ExpiresAt.Unix(),
		Result:              claims.Result,
		ProfileIdentifier:   claims.ProfileIdentifier,
		Nonce:               claims.Nonce,
		UserAuthorizationID: claims.UserAuthorizationID,
		ReferenceID:         claims.ReferenceID,
	}, err
}
