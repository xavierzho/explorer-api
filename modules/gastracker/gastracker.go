package gastracker

import (
	"github.com/xavierzho/explorer-api/modules"
	"github.com/xavierzho/explorer-api/utils"
)

type Service modules.Service

// Name returns the name of this service.
func (s *Service) Name() string { return "gastracker" }

// GasEstimate Get Estimation of Confirmation Time
//
// description: https://docs.etherscan.io/api-endpoints/gas-tracker#get-estimation-of-confirmation-time
//
// Returns the estimated time, in seconds, for a transaction to be confirmed on the blockchain.
func (s *Service) GasEstimate(gasprice string) (gas string, err error) {
	err = s.Client.Call(s, "gasestimate", map[string]string{
		"gasprice": gasprice,
	}, &gas)
	return
}

// GasOracle Get Gas Oracle
//
// description: https://docs.etherscan.io/api-endpoints/gas-tracker#get-gas-oracle
//
// Returns the current Safe, Proposed and Fast gas prices.
func (s *Service) GasOracle() (gasOracle GasOracle, err error) {
	err = s.Client.Call(s, "gasoracle", nil, &gasOracle)
	return
}

// DailyAverageGasLimit Get Daily Average Gas Limit. [PRO]
//
// description:
//
// Returns the historical daily average gas limit of the Ethereum network.
func (s *Service) DailyAverageGasLimit(start, end utils.Time, isDesc bool) (gasLimit []DailyAverageGasLimit, err error) {
	sortStr := "asc"
	if isDesc {
		sortStr = "desc"
	}
	err = s.Client.Call(s, "dailyavggaslimit", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sortStr,
	}, &gasLimit)
	return
}

// DailyGasUsed Get Ethereum Daily Total Gas Used.[PRO]
//
// description:https://docs.etherscan.io/api-endpoints/gas-tracker#get-ethereum-daily-total-gas-used
//
// Returns the total amount of gas used daily for transactions on the Ethereum network.
func (s *Service) DailyGasUsed(start, end utils.Time, isDesc bool) (gasUsed []DailyGasUsed, err error) {
	sortStr := "asc"
	if isDesc {
		sortStr = "desc"
	}
	err = s.Client.Call(s, "dailygasused", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sortStr,
	}, &gasUsed)
	return
}

// DailyGasPrice Get Ethereum Daily Gas Price.[PRO]
//
// description:https://docs.etherscan.io/api-endpoints/gas-tracker#get-ethereum-daily-gas-price
//
// Returns the daily gas price of the Ethereum network.
func (s *Service) DailyGasPrice(start, end utils.Time, isDesc bool) (gasPrice []DailyGasPrice, err error) {
	sortStr := "asc"
	if isDesc {
		sortStr = "desc"
	}
	err = s.Client.Call(s, "dailyavggasprice", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sortStr,
	}, &gasPrice)
	return
}
