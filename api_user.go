package paypayopa

import (
	"context"
	"net/url"
	"time"
)

func createAccountLinkQrCode(
	ctx context.Context,
	client opaClient,
	req *CreateAccountLinkQrCodePayload,
) (*CreateAccountLinkQrCodeResponse, *ResultInfo, error) {
	const timeout = 10 * time.Second

	res := &CreateAccountLinkQrCodeResponse{}
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

func getMaskedUserProfile(
	ctx context.Context,
	client opaClient,
	userAuthorizationID string,
) (*MaskedUserProfileResponse, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &MaskedUserProfileResponse{}
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

func getUserAuthorizationStatus(
	ctx context.Context,
	client opaClient,
	userAuthorizationID string,
) (*GetUserAuthorizationStatusResponse, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &GetUserAuthorizationStatusResponse{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v2/user/authorizations/"+userAuthorizationID,
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func unlinkUser(
	ctx context.Context,
	client opaClient,
	userAuthorizationID string,
) (*ResultInfo, error) {
	const timeout = 15 * time.Second

	return client.DELETE(
		ctxWithTimeout(ctx, timeout),
		"/v2/user/authorizations/"+userAuthorizationID,
	)
}
