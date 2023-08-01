package explorer

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RTLimiter struct {
	transport http.RoundTripper
	retries   int
	ticker    *time.Ticker
}

// NewRTLimiter create a new rate limiter with request per second(rps) and retry times
func NewRTLimiter(rps int, transport http.RoundTripper, retry int) *RTLimiter {
	interval := time.Second / time.Duration(rps)
	return &RTLimiter{transport: transport, retries: retry, ticker: time.NewTicker(interval)}
}

// RoundTrip implements the RoundTripper interface.
// Exponential backoff is used to retry requests that fail due to rate limiting.
func (r *RTLimiter) RoundTrip(req *http.Request) (*http.Response, error) {
	retryInternal := time.Second
	for i := 0; i < r.retries; i++ {
		<-r.ticker.C
		resp, err := r.transport.RoundTrip(req)
		if err != nil {
			continue
		}
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, resp.Body)

		if bytes.Contains(buf.Bytes(), []byte("rate limit")) {
			time.Sleep(retryInternal)
			// exponential backoff strategy
			retryInternal *= 2
			if retryInternal > time.Minute {
				retryInternal = time.Second
			}
			continue
		}
		resp.Body = io.NopCloser(&buf)
		return resp, nil
	}
	return nil, fmt.Errorf("retries exceeded %d", r.retries)
}
