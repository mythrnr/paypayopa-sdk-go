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
	timeout := 10
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &CreateAccountLinkQrCodeResponse{}
	info, err := client.POST(
		ctx,
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
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &MaskedUserProfileResponse{}
	info, err := client.GET(
		ctx,
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
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &GetUserAuthorizationStatusResponse{}
	info, err := client.GET(
		ctx,
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
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	return client.DELETE(ctx, "/v2/user/authorizations/"+userAuthorizationID)
}
