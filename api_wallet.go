package paypayopa

import (
	"context"
	"net/url"
	"strconv"
	"time"
)

func checkUserWalletBalance(
	ctx context.Context,
	client *opaClient,
	req *CheckUserWalletBalancePayload,
) (*CheckUserWalletBalance, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &CheckUserWalletBalance{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v2/wallet/check_balance?"+url.Values{
			"userAuthorizationId": []string{req.UserAuthorizationID},
			"amount": []string{
				strconv.FormatInt(int64(req.Amount), 10),
			},
			"currency":    []string{string(req.Currency)},
			"productType": []string{req.ProductType},
		}.Encode(),
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func createTopupQRCode(
	ctx context.Context,
	client *opaClient,
	req *CreateTopupQRCodePayload,
) (*TopupQRCodeResponse, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &TopupQRCodeResponse{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v1/code/topup",
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func deleteTopupQRCode(
	ctx context.Context,
	client *opaClient,
	codeID string,
) (*ResultInfo, error) {
	const timeout = 30 * time.Second

	return client.DELETE(
		ctxWithTimeout(ctx, timeout),
		"/v1/code/topup/"+codeID,
	)
}

func getTopupDetails(
	ctx context.Context,
	client *opaClient,
	merchantTopUpID string,
) (*TopupQRCodeDetailsResponse, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &TopupQRCodeDetailsResponse{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v1/code/topup/"+merchantTopUpID,
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func getUserWalletBalance(
	ctx context.Context,
	client *opaClient,
	req *GetUserWalletBalancePayload,
) (*UserWalletBalanceResponse, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &UserWalletBalanceResponse{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v6/wallet/balance?"+url.Values{
			"userAuthorizationId": []string{req.UserAuthorizationID},
			"currency":            []string{string(req.Currency)},
			"productType":         []string{req.ProductType},
		}.Encode(),
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}
