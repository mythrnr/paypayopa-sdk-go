package paypayopa

import (
	"context"
	"time"
)

type RefundPaymentPayload struct {
	MerchantRefundID string       `json:"merchantRefundId"`
	PaymentID        string       `json:"paymentId"`
	Amount           *MoneyAmount `json:"amount"`
	RequestedAt      int64        `json:"requestedAt"`
	Reason           string       `json:"reason"`
}

type RefundResponse struct {
	Status           string       `json:"status"`
	AcceptedAt       int64        `json:"acceptedAt"`
	MerchantRefundID string       `json:"merchantRefundId"`
	PaymentID        string       `json:"paymentId"`
	Amount           *MoneyAmount `json:"amount"`
	RequestedAt      int64        `json:"requestedAt"`
	Reason           string       `json:"reason"`
	AssumeMerchant   string       `json:"assumeMerchant"`
}

func refundPayment(
	ctx context.Context,
	client *opaClient,
	req *RefundPaymentPayload,
) (*RefundResponse, *ResultInfo, error) {
	const timeout = 30 * time.Second

	res := &RefundResponse{}
	info, err := client.POST(
		ctxWithTimeout(ctx, timeout),
		"/v2/refunds",
		res,
		req,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}

func getRefundDetails(
	ctx context.Context,
	client *opaClient,
	merchantRefundID string,
) (*RefundResponse, *ResultInfo, error) {
	const timeout = 15 * time.Second

	res := &RefundResponse{}
	info, err := client.GET(
		ctxWithTimeout(ctx, timeout),
		"/v2/refunds/"+merchantRefundID,
		res,
	)

	if err != nil || !info.Success() {
		return nil, info, err
	}

	return res, info, nil
}
