package paypayopa

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	headerNameAuth     = "Authorization"
	headerNameMerchant = "X-ASSUME-MERCHANT"

	contentType         = "application/json;charset=UTF-8;"
	contentTypeEmpty    = "empty"
	headerAuthPrefix    = "hmac OPA-Auth"
	recommendedNonceLen = 8
)

// authenticate is a structure for generating API credentials.
// Create and destroy on each request.
//
// authenticate は API の認証情報を生成するための構造体.
// リクエストごとに作成して破棄する.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#tag/API
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#tag/API
type authenticate struct {
	apiKey       string
	apiKeySecret string
	body         []byte
	epoch        int64
	method       string
	nonce        string
	uri          string

	hashCache string
}

func newAuthenticate(apiKey, apiKeySecret string) *authenticate {
	return &authenticate{
		apiKey:       apiKey,
		apiKeySecret: apiKeySecret,
		epoch:        time.Now().Unix(),
		nonce:        Nonce(recommendedNonceLen),
	}
}

func (a *authenticate) contentType() string {
	if len(a.body) == 0 {
		return contentTypeEmpty
	}

	return contentType
}

func (a *authenticate) setRequest(req *http.Request) error {
	var body []byte

	if req.GetBody != nil {
		b, err := req.GetBody()
		if err != nil {
			return fmt.Errorf("failed to read body: %w", err)
		}

		body, err = ioutil.ReadAll(b)
		if err != nil {
			return fmt.Errorf("failed to read body: %w", err)
		}
	}

	a.body = body
	a.method = req.Method
	a.uri = req.URL.Path

	return nil
}

// hash is the process of Step 1 of the authentication header creation.
// The hash is cached to be used for each request.
//
// hash は認証ヘッダ作成の Step 1 の処理.
// リクエストごとに使い捨てる為, hash はキャッシュする.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#section/HMAC-auth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#section/HMAC-auth
func (a *authenticate) hash() (string, error) {
	if len(a.body) == 0 {
		a.hashCache = contentTypeEmpty
	}

	if a.hashCache != "" {
		return a.hashCache, nil
	}

	hash := md5.New()
	if _, err := hash.Write([]byte(a.contentType())); err != nil {
		return "", fmt.Errorf("failed to write hash: %w", err)
	}

	if _, err := hash.Write(a.body); err != nil {
		return "", fmt.Errorf("failed to write hash: %w", err)
	}

	a.hashCache = base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return a.hashCache, nil
}

// hash is the process of Step 2 of the authentication header creation.
// hash は認証ヘッダ作成の Step 2 の処理.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#section/HMAC-auth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#section/HMAC-auth
func (a *authenticate) macData() ([]byte, error) {
	h, err := a.hash()
	if err != nil {
		return nil, err
	}

	segments := []string{
		a.uri,
		a.method,
		a.nonce,
		strconv.FormatInt(a.epoch, 10),
		a.contentType(),
		h,
	}

	return []byte(strings.Join(segments, "\n")), nil
}

// hash is the process of Step 3 of the authentication header creation.
// hash は認証ヘッダ作成の Step 3 の処理.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#section/HMAC-auth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#section/HMAC-auth
func (a *authenticate) base64hmacString() (string, error) {
	b, err := a.macData()
	if err != nil {
		return "", err
	}

	mac := hmac.New(sha256.New, []byte(a.apiKeySecret))
	if _, err := mac.Write(b); err != nil {
		return "", fmt.Errorf("failed to write hmac: %w", err)
	}

	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

// hash is the process of Step 4 of the authentication header creation.
// hash は認証ヘッダ作成の Step 4 の処理.
//
// API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#section/HMAC-auth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#section/HMAC-auth
func (a *authenticate) hmacHeader() (string, error) {
	h, err := a.hash()
	if err != nil {
		return "", err
	}

	macData, err := a.base64hmacString()
	if err != nil {
		return "", err
	}

	segments := []string{
		headerAuthPrefix,
		a.apiKey,
		macData,
		a.nonce,
		strconv.FormatInt(a.epoch, 10),
		h,
	}

	return strings.Join(segments, ":"), nil
}
