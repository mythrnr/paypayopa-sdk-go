package paypayopa

import "encoding/json"

type CheckUserWalletBalancePayload struct {
	GetUserWalletBalancePayload
	Amount int
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

type GetUserWalletBalancePayload struct {
	UserAuthorizationID string
	Currency            Currency
	ProductType         string
}

type CheckUserWalletBalance struct {
	HasEnoughBalance bool `json:"hasEnoughBalance"`
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

type UserWalletBalanceResponse struct {
	UserAuthorizationID string       `json:"userAuthorizationId"`
	TotalBalance        *MoneyAmount `json:"totalBalance"`
	Preference          struct {
		UseCashback            bool `json:"useCashback"`
		CashbackAutoInvestment bool `json:"cashbackAutoInvestment"`
	} `json:"preference"`
}
