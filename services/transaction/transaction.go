package transaction

import "github.com/Jonescy/explorer-api/services"

type Service services.Service

func (*Service) Name() string { return "transaction" }

// GetExecutionStatus returns the status of a given transaction hash.
func (s *Service) GetExecutionStatus(txhash string) (status ExecStatus, err error) {
	err = s.Client.Call(s.Name(), "getstatus", map[string]string{"txhash": txhash}, &status)
	return
}

// GetReceiptStatus returns the status of a given transaction hash.
func (s *Service) GetReceiptStatus(txhash string) (status ReceiptStatus, err error) {
	err = s.Client.Call(s.Name(), "gettxreceiptstatus", map[string]string{"txhash": txhash}, &status)
	return
}
