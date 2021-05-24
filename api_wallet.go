package paypayopa

import (
	"context"
	"net/url"
	"strconv"
	"time"
)

func checkUserWalletBalance(
	ctx context.Context,
	client opaClient,
	req *CheckUserWalletBalancePayload,
) (*CheckUserWalletBalance, *ResultInfo, error) {
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &CheckUserWalletBalance{}
	info, err := client.GET(
		ctx,
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
	client opaClient,
	req *CreateTopupQRCodePayload,
) (*TopupQRCodeResponse, *ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &TopupQRCodeResponse{}
	info, err := client.POST(
		ctx,
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
	client opaClient,
	codeID string,
) (*ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	return client.DELETE(ctx, "/v1/code/topup/"+codeID)
}

func getTopupDetails(
	ctx context.Context,
	client opaClient,
	merchantTopUpID string,
) (*TopupQRCodeDetailsResponse, *ResultInfo, error) {
	timeout := 30
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &TopupQRCodeDetailsResponse{}
	info, err := client.GET(
		ctx,
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
	client opaClient,
	req *GetUserWalletBalancePayload,
) (*UserWalletBalanceResponse, *ResultInfo, error) {
	timeout := 15
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	defer cancel()

	res := &UserWalletBalanceResponse{}
	info, err := client.GET(
		ctx,
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
