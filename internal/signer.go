package internal

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	contentType         = "application/json;charset=UTF-8;"
	contentTypeEmpty    = "empty"
	recommendedNonceLen = 8
)

// Signer is a structure for generating API credentials.
// Create and destroy on each request.
//
// Signer は API の認証情報を生成するための構造体.
// リクエストごとに作成して破棄する.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#tag/API
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#tag/API
type Signer struct {
	apiKey       string
	apiKeySecret string
	body         []byte
	epoch        int64
	method       string
	nonce        string
	uri          string

	hashCache string
}

func NewSigner(
	apiKey, apiKeySecret string,
	req *http.Request,
) (*Signer, error) {
	a := &Signer{
		apiKey:       apiKey,
		apiKeySecret: apiKeySecret,
		epoch:        time.Now().Unix(),
		method:       req.Method,
		nonce:        nonce(recommendedNonceLen),
		uri:          req.URL.Path,
	}

	if req.GetBody == nil {
		return a, nil
	}

	b, err := req.GetBody()
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	if a.body, err = io.ReadAll(b); err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	return a, nil
}

func (a *Signer) ContentType() string {
	if len(a.body) == 0 {
		return contentTypeEmpty
	}

	return contentType
}

// hash is the process of Step 1 of the authentication header creation.
// The hash is cached to be used for each request.
//
// hash は認証ヘッダ作成の Step 1 の処理.
// リクエストごとに使い捨てる為, hash はキャッシュする.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#section/HMAC-auth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#section/HMAC-auth
func (a *Signer) hash() string {
	if len(a.body) == 0 {
		a.hashCache = contentTypeEmpty
	}

	if a.hashCache != "" {
		return a.hashCache
	}

	hash := md5.New()
	hash.Write([]byte(a.ContentType()))
	hash.Write(a.body)

	a.hashCache = base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return a.hashCache
}

// macData is the process of Step 2 of the authentication header creation.
//
// macData は認証ヘッダ作成の Step 2 の処理.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#section/HMAC-auth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#section/HMAC-auth
func (a *Signer) macData() []byte {
	segments := []string{
		a.uri,
		a.method,
		a.nonce,
		strconv.FormatInt(a.epoch, 10),
		a.ContentType(),
		a.hash(),
	}

	return []byte(strings.Join(segments, "\n"))
}

// base64hmacString is the process of Step 3 of the authentication header creation.
//
// base64hmacString は認証ヘッダ作成の Step 3 の処理.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#section/HMAC-auth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#section/HMAC-auth
func (a *Signer) base64hmacString() string {
	mac := hmac.New(sha256.New, []byte(a.apiKeySecret))
	mac.Write(a.macData())

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// Sign is the process of Step 4 of the authentication header creation.
//
// Sign は認証ヘッダ作成の Step 4 の処理.
//
// # API Docs
//
// EN: https://www.paypay.ne.jp/opa/doc/v1.0/webcashier#section/HMAC-auth
//
// JP: https://www.paypay.ne.jp/opa/doc/jp/v1.0/webcashier#section/HMAC-auth
func (a *Signer) Sign() string {
	segments := []string{
		"hmac OPA-Auth",
		a.apiKey,
		a.base64hmacString(),
		a.nonce,
		strconv.FormatInt(a.epoch, 10),
		a.hash(),
	}

	return strings.Join(segments, ":")
}
