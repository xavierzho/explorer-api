package blocks

import (
	"github.com/Jonescy/explorer-api/modules"
	"github.com/Jonescy/explorer-api/utils"
	"strconv"
)

type Service modules.Service

func (*Service) Name() string { return "block" }

// BlockReward Get Block And Uncle Rewards by BlockNo
//
// description: https://docs.etherscan.io/api-endpoints/blocks#get-block-and-uncle-rewards-by-blockno
//
// Returns the block reward and 'Uncle' block rewards.
func (s *Service) BlockReward(blockNo int) (reward Reward, err error) {
	err = s.Client.Call(s, "getblockreward", map[string]string{
		"blockno": strconv.Itoa(blockNo),
	}, &reward)

	return
}

// BlockCountDown Get Estimated Block Countdown Time by BlockNo
//
// description: https://docs.etherscan.io/api-endpoints/blocks#get-estimated-block-countdown-time-by-blockno
//
// Returns the estimated time remaining, in seconds, until a certain block is mined.
func (s *Service) BlockCountDown(blockNo int) (countdown CountDown, err error) {

	err = s.Client.Call(s, "getblockcountdown", map[string]string{
		"blockno": strconv.Itoa(blockNo),
	}, &countdown)
	return
}

// BlockNoByTimestamp Get Block Number by Timestamp
//
// description: https://docs.etherscan.io/api-endpoints/blocks#get-block-number-by-timestamp
//
// Returns the block number that was mined at a certain timestamp.
func (s *Service) BlockNoByTimestamp(ts int, closest string) (blockNo string, err error) {
	err = s.Client.Call(s, "getblocknobytime", map[string]string{
		"timestamp": strconv.Itoa(ts),
		"closest":   closest,
	}, &blockNo)

	return
}

// GetDailyAvgBlockSize Get Daily Average Block Size. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/blocks#get-daily-average-block-size
//
// Returns the average block size for a given timestamp.
func (s *Service) GetDailyAvgBlockSize(start, end utils.Time, isDesc bool) (size []Size, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "dailyavgblocksize", map[string]string{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &size)

	return
}

// GetDailyBlockCountRewards Get Daily Block Count and Rewards. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/blocks#get-daily-block-count-and-rewards
//
// Returns the number of blocks mined daily and the amount of block rewards.
func (s *Service) GetDailyBlockCountRewards(start, end utils.Time, isDesc bool) (counter []DailyBlockCounter, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "dailyblkcount", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &counter)

	return
}

// GetDailyBlockRewards Get Daily Block Rewards. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/blocks#get-daily-block-rewards
//
// Returns the amount of block rewards distributed to miners daily.
func (s *Service) GetDailyBlockRewards(start, end utils.Time, isDesc bool) (counter []DailyBlockReward, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "dailyblockrewards", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &counter)

	return
}

// GetAvgMinedSecond Get Daily Average Time for A Block to be Included in the Ethereum Blockchain. [PRO]
//
// description:https://docs.etherscan.io/api-endpoints/blocks#get-daily-average-time-for-a-block-to-be-included-in-the-ethereum-blockchain
//
// Returns the daily average of time needed for a block to be successfully mined.
func (s *Service) GetAvgMinedSecond(start, end utils.Time, isDesc bool) (counter []AverageMinedTime, err error) {
	sort := "asc"
	if isDesc {
		sort = "desc"
	}
	err = s.Client.Call(s, "dailyavgblocktime", utils.M{
		"startdate": start.FormatDate(),
		"enddate":   end.FormatDate(),
		"sort":      sort,
	}, &counter)

	return
}
