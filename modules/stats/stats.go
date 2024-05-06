package stats

import (
	"strconv"

	"github.com/xavierzho/explorer-api/iface"

	"github.com/xavierzho/explorer-api/utils"
)

type Service iface.Service

func (*Service) Name() string { return "stats" }

// EthSupply Get Total supply of Ether
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-total-supply-of-ether
//
// Returns the current amount of Ether in circulation excluding ETH2 Staking rewards and EIP1559 burnt fees.
func (s *Service) EthSupply() (supply *utils.BN, err error) {
	err = s.Client.Call(s, "ethsupply", nil, &supply)
	return
}

// Eth2Supply Get Total Supply of Ether 2
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-total-supply-of-ether-2
//
// Returns the current amount of Ether in circulation, ETH2 Staking rewards, EIP1559 burnt fees, and total withdrawn ETH from the beacon chain.
func (s *Service) Eth2Supply() (supply iface.Supply, err error) {
	err = s.Client.Call(s, "ethsupply2", nil, &supply)
	return
}

// EthLastPrice Get Ether Last Price
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-ether-last-price
//
// Returns the latest price of 1 ETH.
func (s *Service) EthLastPrice() (prices iface.Prices, err error) {
	err = s.Client.Call(s, "ethprice", nil, &prices)
	return
}

// GetNodeSize Get Ethereum Nodes Size
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-ethereum-nodes-size
//
// Returns the size of the Ethereum blockchain, in bytes, over a date range.
func (s *Service) GetNodeSize(startDate, endDate string, isAsc bool) (sizes []iface.NodeSize, err error) {
	sort := "asc"
	if !isAsc {
		sort = "desc"
	}
	err = s.Client.Call(s, "chainsize", utils.M{
		"startdate":  startDate,
		"enddate":    endDate,
		"sort":       sort,
		"clienttype": "geth",
		"syncmode":   "default",
	}, &sizes)
	return
}

// NodeCount Get Total Nodes Count
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-total-nodes-count
//
// Returns the total number of discoverable Ethereum nodes.
func (s *Service) NodeCount() (nodeCount iface.NodeCount, err error) {
	err = s.Client.Call(s, "nodecount", nil, &nodeCount)
	return
}

// GetDailyTxFee Get Daily Network Transaction Fee. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-daily-network-transaction-fee
//
// Returns the amount of transaction fees paid to miners per day.
func (s *Service) GetDailyTxFee(start, end utils.Time, isDesc bool) (fees []iface.DailyTxFee, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "dailytxnfee", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &fees)
	return

}

// GetDailyNewAddress Get Daily New Address Count. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-daily-new-address-count
//
// Returns the number of new addresses created per day.
func (s *Service) GetDailyNewAddress(start, end utils.Time, isDesc bool) (addresses []iface.DailyNewAddress, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "dailynewaddress", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &addresses)
	return
}

// GetDailyNetworkUtilization Get Daily Network Utilization.[PRO]
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-daily-network-utilization
//
// Returns the daily average gas used over gas limit, in percentage.
func (s *Service) GetDailyNetworkUtilization(start, end utils.Time, isDesc bool) (utilizations []iface.DailyNetworkUtilization, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "dailynetworkutilization", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &utilizations)
	return
}

// GetDailyAvgNetworkHashRate Get Daily Average Network Hash Rate. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-daily-average-network-hash-rate
//
// Returns the historical measure of processing power of the Ethereum network.
// Tips: The networkHashRate is represented in GigaHashes ( GH/s ).
func (s *Service) GetDailyAvgNetworkHashRate(start, end utils.Time, isDesc bool) (hashRates []iface.DailyAvgNetworkHashRate, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "dailyavgblocksize", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &hashRates)
	return
}

// GetDailyTxCount Get Daily Transaction Count. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-daily-transaction-count
//
// Returns the number of transactions per day.
func (s *Service) GetDailyTxCount(start, end utils.Time, isDesc bool) (txCounts []iface.DailyTxCount, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "dailytx", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &txCounts)
	return
}

// GetDailyAvgNetworkDifficulty Get Daily Average Network Difficulty. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-daily-average-network-difficulty
//
// Returns the daily average network difficulty.
func (s *Service) GetDailyAvgNetworkDifficulty(start, end utils.Time, isDesc bool) (difficulties []iface.DailyAvgNetworkDifficulty, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "dailyavgnetdifficulty", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &difficulties)
	return
}

// GetEtherHistoricalDailyMarketCap Get Ether Historical Daily Market Cap. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-ether-historical-daily-market-cap
//
// Returns the historical Ether daily market capitalization.
// Tip : The marketCap is represented in million US Dollars ( USD ).
func (s *Service) GetEtherHistoricalDailyMarketCap(start, end utils.Time, isDesc bool) (marketCaps []iface.EtherHistoricalDailyMarketCap, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "ethdailymarketcap", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &marketCaps)
	return
}

// GetEtherHistoricalPrice Get Ether Historical Price. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/stats-1#get-ether-historical-price
//
// Returns the historical price of 1 ETH.
// Tip : The price is represented in USD.
func (s *Service) GetEtherHistoricalPrice(start, end utils.Time, isDesc bool) (prices []iface.EtherHistoricalPrice, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "ethdailyprice", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &prices)
	return
}

// GetERC20Supply Get ERC20-Token TotalSupply by ContractAddress
//
// description: https://docs.etherscan.io/api-endpoints/tokens#get-total-supply-by-contractaddress
//
// Returns the current amount of an ERC-20 token in circulation.
func (s *Service) GetERC20Supply(contractAddress string) (supply *utils.BN, err error) {
	err = s.Client.Call(s, "tokensupply", utils.M{
		"contractaddress": contractAddress,
	}, &supply)
	return
}

// GetHistoricalERC20Supply Get Historical ERC20-Token TotalSupply by ContractAddress & BlockNo.[PRO]
//
// description: https://docs.etherscan.io/api-endpoints/tokens#get-historical-erc20-token-totalsupply-by-contractaddress-and-blockno
//
// Returns the amount of an ERC-20 token in circulation at a certain block height.
func (s *Service) GetHistoricalERC20Supply(contractAddress string, blockNo uint64) (supply *utils.BN, err error) {
	err = s.Client.Call(s, "tokennhrsupply", utils.M{
		"contractaddress": contractAddress,
		"blockno":         strconv.FormatUint(blockNo, 10),
	}, &supply)
	return
}
