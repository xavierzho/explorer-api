package gastracker

import "github.com/Jonescy/explorer-api"

type Action interface {
	explorer.Endpoint
	GasEstimate(gasprice string) (gas string, err error)
	GasOracle() (gasOracle GasOracle, err error)
}

type GasOracle struct {
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	LastBlock       string `json:"LastBlock"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
	FastGasPrice    string `json:"FastGasPrice"`
	GasUsedRatio    string `json:"gasUsedRatio"`
}
