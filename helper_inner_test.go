package paypayopa

import (
	"bytes"
	"io"
	"net/http"

	"github.com/mythrnr/paypayopa-sdk-go/internal/mocks"
	"github.com/stretchr/testify/mock"
)

func newTestClient(rt http.RoundTripper) *http.Client {
	hc := new(http.Client)
	hc.Transport = rt

	return hc
}

func newTestRoundTripper(
	statusCode int,
	body string,
	run ...func(mock.Arguments),
) *mocks.RoundTripper {
	return newTestRoundTripperWithBody(
		statusCode,
		io.NopCloser(bytes.NewBufferString(body)),
		run...,
	)
}

func newTestRoundTripperWithBody(
	statusCode int,
	body io.ReadCloser,
	run ...func(mock.Arguments),
) *mocks.RoundTripper {
	res := new(http.Response)
	res.Status = http.StatusText(statusCode)
	res.StatusCode = statusCode
	res.Body = body

	rt := new(mocks.RoundTripper)
	call := rt.On("RoundTrip", mock.Anything)

	for _, r := range run {
		call.Run(r)
	}

	call.Return(res, nil)

	return rt
}
