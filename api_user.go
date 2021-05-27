package paypayopa

import (
	"context"
	"encoding/json"
	"net/url"
	"time"
)

type CreateAccountLinkQRCodePayload struct {
	Scopes       []Scope `json:"scopes"`
	Nonce        string  `json:"nonce"`
	RedirectType string  `json:"redirectType"`
	RedirectURL  string  `json:"redirectUrl"`
	ReferenceID  string  `json:"referenceId"`
	PhoneNumber  string  `json:"phoneNumber"`
	DeviceID     string  `json:"deviceId"`
	UserAgent    string  `json:"userAgent"`
}

type CreateAccountLinkQRCodeResponse struct {
	LinkQRCodeURL string `json:"linkQRCodeURL"`
}

func createAccountLinkQRCode(
	ctx context.Context,
	client *opaClient,
	req *CreateAccountLinkQRCodePayload,
) (*CreateAccountLinkQRCodeResponse, *ResultInfo, error) {
	const timeout = 10 * time.Second

	res := &CreateAccountLinkQRCodeResponse{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v1/qr/sessions",
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

type GetUserAuthorizationStatusResponse struct {
	UserAuthorizationID string           `json:"userAuthorizationId"`
	ReferenceIDs        *json.RawMessage `json:"referenceIds"`
	Status              string           `json:"status"`
	Scopes              []string         `json:"scopes"`
	ExpireAt            int64            `json:"expireAt"`
	IssuedAt            int64            `json:"issuedAt"`
}

func getUserAuthorizationStatus(
	ctx context.Context,
	client *opaClient,
	userAuthorizationID string,
) (*GetUserAuthorizationStatusResponse, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &GetUserAuthorizationStatusResponse{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v2/user/authorizations?"+url.Values{
			"userAuthorizationId": []string{userAuthorizationID},
		}.Encode(),
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func unlinkUser(
	ctx context.Context,
	client *opaClient,
	userAuthorizationID string,
) (*ResultInfo, error) {
	const timeout = 15 * time.Second

	return client.DELETE(
		ctxWithTimeout(ctx, timeout),
		"/v2/user/authorizations/"+userAuthorizationID,
	)
}

type GetMaskedUserProfileResponse struct {
	PhoneNumber string `json:"phoneNumber"`
}

func getMaskedUserProfile(
	ctx context.Context,
	client *opaClient,
	userAuthorizationID string,
) (*GetMaskedUserProfileResponse, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &GetMaskedUserProfileResponse{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v2/user/profile/secure?"+url.Values{
			"userAuthorizationId": []string{userAuthorizationID},
		}.Encode(),
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}
