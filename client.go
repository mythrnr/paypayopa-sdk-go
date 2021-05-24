package paypayopa

import (
	"bytes"
	"context"
	"encoding/json"
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

// opaClient is the interface definition for
// handling requests/responses to the PayPay API.
//
// opaClient は PayPay API へのリクエスト/レスポンスの
// ハンドリングを行うインターフェース定義.
type opaClient interface {
	// GET sends a GET request.
	// response is stored in res, and returns
	// the processing result and error object.
	//
	// GET は GET リクエストを送信する.
	// res にレスポンスを格納し, 処理結果とエラーオブジェクトを返す.
	GET(ctx context.Context, path string, res interface{}) (*ResultInfo, error)

	// DELETE sends a DELETE request, and returns
	// the processing result and error object.
	//
	// DELETE は DELETE リクエストを送信し, 処理結果とエラーオブジェクトを返す.
	DELETE(ctx context.Context, path string) (*ResultInfo, error)

	// POST sends a POST request.
	// response is stored in res, and returns
	// the processing result and error object.
	//
	// POST は POST リクエストを送信する.
	// res にレスポンスを格納し, 処理結果とエラーオブジェクトを返す.
	POST(ctx context.Context, path string, res interface{}, req interface{}) (*ResultInfo, error)

	// Request creates a *http.Request from arguments.
	// Basically, you can simply execute the GET, DELETE, and POST methods,
	// but if you want to configure the request further, call this method.
	// The configured request is sent by calling the Do method.
	//
	// Requestは、引数から*http.Requestを作成する.
	// 基本的には単純に GET, DELETE, POST のメソッドを実行すればいいが,
	// リクエストを更に設定したい場合はこのメソッドを呼ぶ.
	// 設定をしたリクエストは Do メソッドを呼んで送信する.
	Request(ctx context.Context, method, path string, req interface{}) (*http.Request, error)

	// Do sends a request.
	// response is stored in res, and returns
	// the processing result and error object.
	//
	// Do はリクエストを送信する.
	// res にレスポンスを格納し, 処理結果とエラーオブジェクトを返す.
	Do(req *http.Request, res interface{}) (*ResultInfo, error)
}

type client struct{ http *http.Client }

var _ opaClient = (*client)(nil)

func newClient(creds *Credential) opaClient {
	return newClientWithHTTPClient(creds, &http.Client{})
}

func newClientWithHTTPClient(creds *Credential, hc *http.Client) opaClient {
	if hc.Timeout == 0 {
		hc.Timeout = timeout
	}

	tr := hc.Transport
	if tr == nil {
		tr = http.DefaultTransport
	}

	hc.Transport = newAuthenticateInterceptor(creds, tr)

	return &client{http: hc}
}

func (c *client) GET(
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

func (c *client) DELETE(
	ctx context.Context,
	path string,
) (*ResultInfo, error) {
	req, err := c.Request(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, nil)
}

func (c *client) POST(
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

func (c *client) Request(
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

func (c *client) Do(req *http.Request, res interface{}) (*ResultInfo, error) {
	defer func() {
		if req.Body != nil {
			req.Body.Close()
		}
	}()

	rs, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
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
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return info, nil
}
