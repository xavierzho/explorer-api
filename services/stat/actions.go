package stat

import (
	"github.com/Jonescy/explorer-api"
	"github.com/Jonescy/explorer-api/utils"
)

type Action interface {
	explorer.Endpoint
	EthSupply() (*utils.BigInt, error)
	Eth2Supply() (Supply, error)
	EthLastPrice() (Prices, error)
	ChainSize(startDate, endDate string, isAsc bool) (sizes []ChainSize, err error)
	NodeCount() (nodeCount NodeCount, err error)
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

type ChainSize struct {
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
