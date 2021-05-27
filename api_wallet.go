package paypayopa

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

type CheckUserWalletBalancePayload struct {
	GetUserWalletBalancePayload
	Amount int
}

type GetUserWalletBalancePayload struct {
	UserAuthorizationID string
	Currency            Currency
	ProductType         ProductType
}

type CheckUserWalletBalance struct {
	HasEnoughBalance bool `json:"hasEnoughBalance"`
}

type UserWalletBalanceResponse struct {
	UserAuthorizationID string       `json:"userAuthorizationId"`
	TotalBalance        *MoneyAmount `json:"totalBalance"`
	Preference          struct {
		UseCashback            bool `json:"useCashback"`
		CashbackAutoInvestment bool `json:"cashbackAutoInvestment"`
	} `json:"preference"`
}

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
			"productType": []string{string(req.ProductType)},
		}.Encode(),
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
			"productType":         []string{string(req.ProductType)},
		}.Encode(),
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

type CreateTopupQRCodePayload struct {
	MerchantTopupID     string           `json:"merchantTopUpId"`
	UserAuthorizationID string           `json:"userAuthorizationId"`
	MinimumTopupAmount  *MoneyAmount     `json:"minimumTopUpAmount"`
	Metadata            *json.RawMessage `json:"metadata"`
	CodeType            CodeType         `json:"codeType"`
	RequestedAt         int64            `json:"requestedAt"`
	RedirectType        string           `json:"redirectType"`
	RedirectURL         string           `json:"redirectUrl"`
	UserAgent           string           `json:"userAgent"`
}

type TopupQRCodeDetailsResponse struct {
	TopupID             string           `json:"topUpId"`
	MerchantTopupID     string           `json:"merchantTopUpId"`
	UserAuthorizationID string           `json:"userAuthorizationId"`
	RequestedAt         int64            `json:"requestedAt"`
	AcceptedAt          int64            `json:"acceptedAt"`
	ExpiryDate          int64            `json:"expiryDate"`
	Status              string           `json:"status"`
	Metadata            *json.RawMessage `json:"metadata"`
}

type TopupQRCodeResponse struct {
	CodeID              string           `json:"codeId"`
	URL                 string           `json:"url"`
	Status              string           `json:"status"`
	MerchantTopupID     string           `json:"merchantTopUpId"`
	UserAuthorizationID string           `json:"userAuthorizationId"`
	MinimumTopupAmount  *MoneyAmount     `json:"minimumTopUpAmount"`
	Metadata            *json.RawMessage `json:"metadata"`
	ExpiryDate          int64            `json:"expiryDate"`
	CodeType            CodeType         `json:"codeType"`
	RequestedAt         int64            `json:"requestedAt"`
	RedirectType        RedirectType     `json:"redirectType"`
	RedirectURL         string           `json:"redirectUrl"`
	UserAagent          string           `json:"userAgent"`
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
