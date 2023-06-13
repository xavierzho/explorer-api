package explorer

import "encoding/json"

// Envelope is the carrier of nearly every response
type Envelope struct {
	JSONRPC string `json:"jsonrpc,omitempty"`
	ID      int    `json:"id,omitempty"`
	// 1 for good, 0 for error
	Status int `json:"status,string,omitempty"`
	// OK for good, other words when Status equals 0
	Message string `json:"message,omitempty"`
	// where response lies
	Result json.RawMessage `json:"result"`
}
