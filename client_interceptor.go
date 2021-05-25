package paypayopa

import (
	"net/http"
	"net/url"
)

type authInterceptor struct {
	creds *Credential
	next  http.RoundTripper
}

var _ http.RoundTripper = (*authInterceptor)(nil)

func newAuthenticateInterceptor(
	creds *Credential,
	next http.RoundTripper,
) http.RoundTripper {
	if creds == nil {
		panic("*Credential must not be nil")
	}

	return &authInterceptor{creds: creds, next: next}
}

// RoundTrip interrupts the request and sets the authentication information.
//
// RoundTrip はリクエストに割り込んで認証情報を設定する.
func (i *authInterceptor) RoundTrip(req *http.Request) (*http.Response, error) {
	a := newAuthenticate(i.creds.apiKey, i.creds.apiKeySecret)
	if err := a.setRequest(req); err != nil {
		return nil, err
	}

	header, err := a.hmacHeader()
	if err != nil {
		return nil, err
	}

	u, _ := url.Parse(string(i.creds.env) + req.URL.String())
	req.URL = u

	req.Header.Set("Content-Type", a.contentType())
	req.Header.Set(headerNameAuth, header)

	if i.creds.merchantID != "" {
		req.Header.Set(headerNameMerchant, i.creds.merchantID)
	}

	return i.next.RoundTrip(req)
}
