package stats

import (
	"github.com/xavierzho/explorer-api"
	"github.com/xavierzho/explorer-api/utils"
)

type ProAction interface {
	GetDailyTxFee(start, end utils.Time, isDesc bool) (fees []DailyTxFee, err error)
	GetDailyNewAddress(start, end utils.Time, isDesc bool) (addresses []DailyNewAddress, err error)
	GetDailyNetworkUtilization(start, end utils.Time, isDesc bool) (utilizations []DailyNetworkUtilization, err error)
	GetDailyAvgNetworkHashRate(start, end utils.Time, isDesc bool) (hashRates []DailyAvgNetworkHashRate, err error)
	GetDailyTxCount(start, end utils.Time, isDesc bool) (txCounts []DailyTxCount, err error)
	GetDailyAvgNetworkDifficulty(start, end utils.Time, isDesc bool) (difficulties []DailyAvgNetworkDifficulty, err error)
	GetEtherHistoricalDailyMarketCap(start, end utils.Time, isDesc bool) (marketCaps []EtherHistoricalDailyMarketCap, err error)
	GetEtherHistoricalPrice(start, end utils.Time, isDesc bool) (prices []EtherHistoricalPrice, err error)
	GetHistoricalERC20Supply(contractAddress string, blockNo uint64) (supply *utils.BN, err error)
}
type Action interface {
	explorer.Module
	EthSupply() (*utils.BN, error)
	Eth2Supply() (Supply, error)
	EthLastPrice() (Prices, error)
	GetNodeSize(startDate, endDate string, isAsc bool) (sizes []NodeSize, err error)
	NodeCount() (nodeCount NodeCount, err error)
	GetERC20Supply(contractAddress string) (supply *utils.BN, err error)
}

type Supply struct {
	EthSupply      string `json:"EthSupply"`
	Eth2Staking    string `json:"Eth2Staking"`
	WithdrawnTotal string `json:"WithdrawnTotal"`
	BurntFees      string `json:"BurntFees"`
}

type Prices struct {
	Ethusd          string `json:"ethusd"`
	Ethbtc          string `json:"ethbtc"`
	EthusdTimestamp string `json:"ethusd_timestamp"`
	EthbtcTimestamp string `json:"ethbtc_timestamp"`
}

type NodeSize struct {
	ClientType     string `json:"clientType"`
	ChainTimeStamp string `json:"chainTimeStamp"`
	BlockNumber    string `json:"blockNumber"`
	SyncMode       string `json:"syncMode"`
	ChainSize      string `json:"chainSize"`
}

type NodeCount struct {
	TotalNodeCount string `json:"TotalNodeCount"`
	UTCDate        string `json:"UTCDate"`
}

type DailyTxFee struct {
	UnixTimeStamp     string `json:"unixTimeStamp"`
	UTCDate           string `json:"UTCDate"`
	TransactionFeeEth string `json:"transactionFee_Eth"`
}
type DailyNewAddress struct {
	NewAddressCount int    `json:"newAddressCount"`
	UnixTimeStamp   string `json:"unixTimeStamp"`
	UTCDate         string `json:"UTCDate"`
}

type DailyNetworkUtilization struct {
	UnixTimeStamp      string `json:"unixTimeStamp"`
	UTCDate            string `json:"UTCDate"`
	NetworkUtilization string `json:"networkUtilization"`
}

type DailyAvgNetworkHashRate struct {
	UnixTimeStamp   string `json:"unixTimeStamp"`
	UTCDate         string `json:"UTCDate"`
	NetworkHashRate string `json:"networkHashRate"`
}

type DailyTxCount struct {
	UnixTimeStamp    string `json:"unixTimeStamp"`
	UTCDate          string `json:"UTCDate"`
	TransactionCount int    `json:"transactionCount"`
}

type DailyAvgNetworkDifficulty struct {
	UnixTimeStamp     string `json:"unixTimeStamp"`
	NetworkDifficulty string `json:"networkDifficulty"`
	UTCDate           string `json:"UTCDate"`
}

type EtherHistoricalDailyMarketCap struct {
	MarketCap     string `json:"marketCap"`
	Price         string `json:"price"`
	UnixTimeStamp string `json:"unixTimeStamp"`
	UTCDate       string `json:"UTCDate"`
	Supply        string `json:"supply"`
}

type EtherHistoricalPrice struct {
	UnixTimeStamp string `json:"unixTimeStamp"`
	UTCDate       string `json:"UTCDate"`
	Value         string `json:"value"`
}
