package explorer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/xavierzho/explorer-api/utils"
)

// DefaultClient default explorer client
func DefaultClient() *Client {
	return &Client{
		conn: &http.Client{
			Transport: NewRTLimiter(1, http.DefaultTransport, 3),
		},
		baseUrl: Ethereum.String(),
	}
}

// ClientWithRTLimiter http.Client with rate limiter with request per second(rps) and retry times
func ClientWithRTLimiter(rps, retry int) *http.Client {
	return &http.Client{
		Transport: NewRTLimiter(rps, http.DefaultTransport, retry),
	}
}

// BeforeHook hook for calling before every request
type BeforeHook func(ctx context.Context, url string) error

// AfterHook hook for calling after every request
type AfterHook func(ctx context.Context, body []byte)

// Client explorer request client
type Client struct {
	conn       *http.Client
	APIKey     string
	baseUrl    string
	BeforeHook BeforeHook
	AfterHook  AfterHook
}

// NewClient new explorer client
//
// if you want to disable rate limit, use http.DefaultClient instead for example:
//
//	NewClient("<you api key>", url, http.DefaultClient)
//
// if you want default rate limit, use nil instead for example:
//
//	NewClient("<you api key>", url, nil)
//
// also you can customize http.RoundTripper， refer RTLimiter:
//
//	rate limit is 1 request per second by default
func NewClient(APIKey string, url Network, hc *http.Client) *Client {
	c := DefaultClient()
	if APIKey == "" {
		fmt.Println("must provide APIKey")
		return nil
	}
	c.APIKey = APIKey
	if url == "" {
		fmt.Println("must select url")
		return nil
	}
	c.baseUrl = url.String()
	if hc != nil {
		c.conn = hc
	}
	return c
}

// Call http call
func (c *Client) Call(module Module, action string, param utils.M, outcome any) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.call(ctx, module.Name(), action, param, outcome)
}

// CallWithContext http call with context
func (c *Client) CallWithContext(ctx context.Context, module, action string, param utils.M, outcome any) error {
	return c.call(ctx, module, action, param, outcome)
}

func (c *Client) call(ctx context.Context, module, action string, param utils.M, outcome any) error {
	var (
		content bytes.Buffer
		err     error
	)
	// build request url
	link := c.buildURL(module, action, param)
	if c.BeforeHook != nil {
		err := c.BeforeHook(ctx, link)
		if err != nil {
			return err
		}
	}

	// recover if there shall be a panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[ouch! panic recovered] please report this with what you did and what you expected, panic detail: %v\n", r)
		}
	}()
	// build request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return err
	}
	// Deprecated: using http round tripper instead
	// do request with rate limit
	//err = c.limiter.Wait(ctx)
	//if err != nil {
	//	fmt.Println("rate limiter wait error:", err)
	//	return err
	//}
	resp, err := c.conn.Do(req)
	if err != nil {
		return err
	}
	// safety close response body
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Printf("error closing response body: %v\n", err)
		}
	}(resp.Body)
	// read response body
	if _, err = io.Copy(&content, resp.Body); err != nil {
		return err
	}
	// check status code
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status %v %s, response body: %s", resp.StatusCode, resp.Status, content.String())
		return err
	}
	// unmarshal response body
	var envelope Envelope
	err = json.Unmarshal(content.Bytes(), &envelope)
	if err != nil {
		return err
	}
	if envelope.Status != 1 && envelope.ID == 0 && envelope.JSONRPC != "2.0" {
		err = fmt.Errorf("etherscan server: %s", envelope.Message)
		return err
	}
	if envelope.Result == nil {
		return fmt.Errorf("rpc error, %s", envelope.Error.Message)
	}

	// unmarshal result
	err = json.Unmarshal(envelope.Result, outcome)
	if err != nil {
		return err
	}
	if c.AfterHook != nil {
		c.AfterHook(ctx, content.Bytes())
	}
	return nil
}

func (c *Client) buildURL(module, action string, param utils.M) (URL string) {
	q := make(url.Values)
	q.Add("module", module)
	q.Add("action", action)
	q.Add("apikey", c.APIKey)
	for k, v := range param {
		q.Add(k, v)
	}
	return fmt.Sprintf("https://%s/api?%s", c.baseUrl, q.Encode())
}

// Post method post for client
func (c *Client) Post(url string, body io.Reader) (resp *http.Response, err error) {
	return c.conn.Post(url, "application/json", body)
}
