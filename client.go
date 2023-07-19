package explorer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/xavierzho/explorer-api/utils"

	"golang.org/x/time/rate"
)

// defaultClient default explorer client
var defaultClient = &Client{
	conn:    http.DefaultClient,
	baseUrl: Ethereum.String(),
	limiter: rate.NewLimiter(rate.Every(time.Second), 1),
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
	limiter    *rate.Limiter
	BeforeHook BeforeHook
	AfterHook  AfterHook
}

// ClientOption client option
type ClientOption func(client *Client)

// WithLimitTier is used to set the rate limit tier
func WithLimitTier(limit Tier) ClientOption {
	return func(client *Client) {
		client.limiter = rate.NewLimiter(rate.Every(time.Second), int(limit))
	}
}

// WithHTTPClient is used to set the http client
func WithHTTPClient(conn *http.Client) ClientOption {
	return func(client *Client) {
		client.conn = conn
	}
}

// WithTimeout is used to set the before hook
func WithTimeout(timeout time.Duration) ClientOption {
	return func(client *Client) {
		client.conn.Timeout = timeout
	}
}

// WithAPIKey is used to set the api APIKey
func WithAPIKey(key string) ClientOption {
	return func(client *Client) {
		client.APIKey = key
	}
}

// WithBaseURL is used to set the base url
func WithBaseURL(url Network) ClientOption {
	return func(client *Client) {
		client.baseUrl = url.String()
	}
}

// NewClient new explorer client
func NewClient(opts ...ClientOption) *Client {
	c := *defaultClient
	for _, opt := range opts {
		opt(&c)
	}
	if err := c.validate(); err != nil {
		return nil
	}
	return &c
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
	// do request with rate limit
	err = c.limiter.Wait(ctx)
	if err != nil {
		return err
	}
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
func (c *Client) validate() error {
	if c.APIKey == "" {
		return errors.New("api APIKey is required for Client")
	}
	if c.limiter == nil {
		return errors.New("rate limiter is required for Client")
	}
	if c.limiter.Burst() < 1 {
		return errors.New("burst rate limit is required for Client")
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
