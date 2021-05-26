package paypayopa

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// timeout is the maximum timeout setting for the client.
// It is recommended to set it to 30s or more, so double it.
// Normally, no response is expected if you wait longer than this.
//
// timeout はクライアントの最長タイムアウト設定.
// 30s 以上に設定することが推奨されている為, 倍を設定.
// 通常, これ以上待機してもレスポンスは期待できない.
const timeout = 60 * time.Second

type ctxkeyTimeout struct{}

func ctxWithTimeout(ctx context.Context, d time.Duration) context.Context {
	return context.WithValue(ctx, &ctxkeyTimeout{}, d)
}

func getTimeout(ctx context.Context) time.Duration {
	if d, ok := ctx.Value(&ctxkeyTimeout{}).(time.Duration); ok {
		return d
	}

	return timeout
}

// opaClient is the client for handling requests/responses to the PayPay API.
//
// opaClient は PayPay API へのリクエスト/レスポンスのハンドリングを行うクライアント.
type opaClient struct{ http *http.Client }

func newClient(creds *Credential) *opaClient {
	return newClientWithHTTPClient(creds, &http.Client{})
}

func newClientWithHTTPClient(creds *Credential, hc *http.Client) *opaClient {
	if hc == nil {
		panic("*http.Client must not be nil")
	}

	if hc.Timeout == 0 {
		hc.Timeout = timeout
	}

	tr := hc.Transport
	if tr == nil {
		tr = http.DefaultTransport
	}

	hc.Transport = newAuthenticateInterceptor(creds, tr)

	return &opaClient{http: hc}
}

// GET sends a GET request.
// response is stored in res, and returns
// the processing result and error object.
//
// GET は GET リクエストを送信する.
// res にレスポンスを格納し, 処理結果とエラーオブジェクトを返す.
func (c *opaClient) GET(
	ctx context.Context,
	path string,
	res interface{},
) (*ResultInfo, error) {
	req, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, res)
}

// DELETE sends a DELETE request, and returns
// the processing result and error object.
//
// DELETE は DELETE リクエストを送信し, 処理結果とエラーオブジェクトを返す.
func (c *opaClient) DELETE(
	ctx context.Context,
	path string,
) (*ResultInfo, error) {
	req, err := c.Request(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, nil)
}

// POST sends a POST request.
// response is stored in res, and returns
// the processing result and error object.
//
// POST は POST リクエストを送信する.
// res にレスポンスを格納し, 処理結果とエラーオブジェクトを返す.
func (c *opaClient) POST(
	ctx context.Context,
	path string,
	res interface{},
	req interface{},
) (*ResultInfo, error) {
	rq, err := c.Request(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}

	return c.Do(rq, res)
}

// Request creates a *http.Request from arguments.
// Basically, you can simply execute the GET, DELETE, and POST methods,
// but if you want to configure the request further, call this method.
// The configured request is sent by calling the Do method.
//
// Requestは、引数から*http.Requestを作成する.
// 基本的には単純に GET, DELETE, POST のメソッドを実行すればいいが,
// リクエストを更に設定したい場合はこのメソッドを呼ぶ.
// 設定をしたリクエストは Do メソッドを呼んで送信する.
func (c *opaClient) Request(
	ctx context.Context,
	method, path string,
	req interface{},
) (*http.Request, error) {
	var reader io.Reader

	if req != nil {
		b, err := json.Marshal(req)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request: %w", err)
		}

		reader = bytes.NewReader(b)
	}

	rq, err := http.NewRequestWithContext(ctx, method, path, reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return rq, nil
}

// Do sends a request.
// response is stored in res, and returns
// the processing result and error object.
//
// Do はリクエストを送信する.
// res にレスポンスを格納し, 処理結果とエラーオブジェクトを返す.
func (c *opaClient) Do(req *http.Request, res interface{}) (*ResultInfo, error) {
	ctx, cancel := context.WithTimeout(
		req.Context(),
		getTimeout(req.Context()),
	)

	req = req.WithContext(ctx)

	defer func() {
		cancel()

		if req.Body != nil {
			req.Body.Close()
		}
	}()

	rs, err := c.http.Do(req)
	if errors.Is(err, context.DeadlineExceeded) {
		return nil, fmt.Errorf("request timeout: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("an error occurred during request: %w", err)
	}

	defer rs.Body.Close()

	b, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	body := &response{}

	if err := json.Unmarshal(b, body); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	info := body.Result
	info.StatusCode = rs.StatusCode

	if res == nil || len(body.Data) == 0 {
		return info, nil
	}

	if err := json.Unmarshal(body.Data, res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response.data: %w", err)
	}

	return info, nil
}
