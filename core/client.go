package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Jonescy/explorer-api/config"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

type BeforeHook func(ctx context.Context, url string) error
type AfterHook func(ctx context.Context, url string, err error)

type Client struct {
	conn       *http.Client
	key        string
	baseUrl    string
	limiter    *rate.Limiter
	beforeHook BeforeHook
	afterHook  AfterHook
}

type ClientOption func(client *Client)

func WithLimitTier(limit config.Tier) ClientOption {
	return func(client *Client) {
		client.limiter = rate.NewLimiter(rate.Every(time.Second), int(limit))
	}
}
func WithHTTPClient(conn *http.Client) ClientOption {
	return func(client *Client) {
		client.conn = conn
	}
}
func WithAPIKey(key string) ClientOption {
	return func(client *Client) {
		client.key = key
	}
}
func WithBaseURL(url config.Network) ClientOption {
	return func(client *Client) {
		client.baseUrl = string(url)
	}
}
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		conn: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	if err := c.validate(); err != nil {
		return nil
	}
	return c
}

func (c *Client) Call(module, action string, param map[string]string, outcome any) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return c.call(ctx, module, action, param, outcome)
}
func (c *Client) CallWithContext(ctx context.Context, module, action string, param map[string]string, outcome any) error {
	return c.call(ctx, module, action, param, outcome)
}
func (c *Client) call(ctx context.Context, module, action string, param map[string]string, outcome any) error {
	link := c.buildURL(module, action, param)
	if c.beforeHook != nil {
		err := c.beforeHook(ctx, link)
		if err != nil {
			return err
		}
	}
	// recover if there shall be an panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("[ouch! panic recovered] please report this with what you did and what you expected, panic detail: %v", r)
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return err
	}
	err = c.limiter.Wait(ctx)
	if err != nil {
		return err
	}
	resp, err := c.conn.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var content bytes.Buffer
	if _, err = io.Copy(&content, resp.Body); err != nil {
		//err = utils.wrapErr(err, "reading response")
		return err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status %v %s, response body: %s", resp.StatusCode, resp.Status, content.String())
		return err
	}

	var envelope Envelope
	err = json.Unmarshal(content.Bytes(), &envelope)
	if err != nil {
		//err = utils.wrapErr(err, "json unmarshal envelope")
		return err
	}
	if envelope.Status != 1 {
		err = fmt.Errorf("etherscan server: %s", envelope.Message)
		return err
	}

	// workaround for missing tokenDecimal for some tokentx calls
	if action == "tokentx" {
		err = json.Unmarshal(bytes.Replace(envelope.Result, []byte(`"tokenDecimal":""`), []byte(`"tokenDecimal":"0"`), -1), outcome)
	} else {
		err = json.Unmarshal(envelope.Result, outcome)
	}
	if err != nil {
		//err = utils.wrapErr(err, "json unmarshal outcome")
		return err
	}

	if c.afterHook != nil {
		c.afterHook(ctx, link, err)
	}
	return nil
}
func (c *Client) validate() error {
	if c.key == "" {
		return errors.New("api key is required for Client")
	}
	if c.limiter == nil {
		return errors.New("rate limiter is required for Client")
	}
	if c.limiter.Burst() < 1 {
		return errors.New("burst rate limit is required for Client")
	}
	return nil
}
func (c *Client) buildURL(module, action string, param map[string]string) (URL string) {
	q := make(url.Values)
	q.Add("module", module)
	q.Add("action", action)
	q.Add("apikey", c.key)
	for k, v := range param {
		q.Add(k, v)
	}
	return fmt.Sprintf("https://%s/api?%s", c.baseUrl, q.Encode())
}

func (c *Client) Post(url string, body io.Reader) (resp *http.Response, err error) {
	return c.conn.Post(url, "application/json", body)
}
