package transactions

import "github.com/xavierzho/explorer-api"

type Action interface {
	explorer.Module
	GetExecutionStatus(txhash string) (status ExecStatus, err error)
	GetReceiptStatus(txhash string) (status ReceiptStatus, err error)
}

type ExecStatus struct {
	IsError        string `json:"isError"`
	ErrDescription string `json:"errDescription"`
}
type ReceiptStatus struct {
	Status string `json:"status"`
}
