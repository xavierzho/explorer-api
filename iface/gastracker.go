package iface

import "github.com/xavierzho/explorer-api/utils"

type GasTracker interface {
	Module
	GasEstimate(gasprice string) (gas string, err error)
	GasOracle() (gasOracle GasOracle, err error)
	DailyAverageGasLimit(start, end utils.Time, isDesc bool) (gasLimit []DailyAverageGasLimit, err error)
	DailyGasUsed(start, end utils.Time, isDesc bool) (gasUsed []DailyGasUsed, err error)
	DailyGasPrice(start, end utils.Time, isDesc bool) (gasPrice []DailyGasPrice, err error)
}

type GasOracle struct {
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	LastBlock       string `json:"LastBlock"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
	FastGasPrice    string `json:"FastGasPrice"`
	GasUsedRatio    string `json:"gasUsedRatio"`
}

type DailyAverageGasLimit struct {
	GasLimit      string `json:"gasLimit"`
	UnixTimeStamp string `json:"unixTimeStamp"`
	UTCDate       string `json:"UTCDate"`
}

type DailyGasUsed struct {
	GasUsed       string `json:"gasUsed"`
	UnixTimeStamp string `json:"unixTimeStamp"`
	UTCDate       string `json:"UTCDate"`
}

type DailyGasPrice struct {
	MinGasPriceWei string `json:"minGasPrice_Wei"`
	UnixTimeStamp  string `json:"unixTimeStamp"`
	UTCDate        string `json:"UTCDate"`
	AvgGasPriceWei string `json:"avgGasPrice_Wei"`
	MaxGasPriceWei string `json:"maxGasPrice_Wei"`
}
