package transactions

import "github.com/xavierzho/explorer-api/iface"

type Service iface.Service

func (*Service) Name() string { return "transaction" }

// GetExecutionStatus Check Contract Execution Status
//
// description: https://docs.etherscan.io/api-endpoints/stats#check-contract-execution-status
//
// Returns the status code of a contract execution.
func (s *Service) GetExecutionStatus(txhash string) (status iface.ExecStatus, err error) {
	err = s.Client.Call(s, "getstatus", map[string]string{"txhash": txhash}, &status)
	return
}

// GetReceiptStatus Check Transaction Receipt Status
//
// description: https://docs.etherscan.io/api-endpoints/stats#check-transaction-receipt-status
//
// Returns the status code of a transaction execution.
func (s *Service) GetReceiptStatus(txhash string) (status iface.ReceiptStatus, err error) {
	err = s.Client.Call(s, "gettxreceiptstatus", map[string]string{"txhash": txhash}, &status)
	return
}
