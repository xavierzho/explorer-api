package block

import "github.com/Jonescy/explorer-api"

type Action interface {
	explorer.Endpoint
	// BlockReward Get Block And Uncle Rewards by BlockNo
	BlockReward(blockNo int) (Reward, error)
	// BlockCountDown Get Estimated Block Countdown Time by BlockNo
	BlockCountDown(blockNo int) (CountDown, error)
	// BlockNoByTimestamp Get Block Number by Timestamp
	BlockNoByTimestamp(ts int, closest string) (blockNo string, err error)
}

type Reward struct {
	TimeStamp string `json:"timeStamp"`
	Uncles    []struct {
		UnclePosition string `json:"unclePosition"`
		Miner         string `json:"miner"`
		BlockReward   string `json:"blockreward"`
	} `json:"uncles"`
	BlockNumber          string `json:"blockNumber"`
	UncleInclusionReward string `json:"uncleInclusionReward"`
	BlockReward          string `json:"blockReward"`
	BlockMiner           string `json:"blockMiner"`
}

type CountDown struct {
	CurrentBlock      string `json:"CurrentBlock"`
	RemainingBlock    string `json:"RemainingBlock"`
	EstimateTimeInSec string `json:"EstimateTimeInSec"`
	CountdownBlock    string `json:"CountdownBlock"`
}

type Sizes []Size
type Size struct {
	BlockSizeBytes int    `json:"blockSize_bytes"`
	UnixTimeStamp  string `json:"unixTimeStamp"`
	UTCDate        string `json:"UTCDate"`
}
