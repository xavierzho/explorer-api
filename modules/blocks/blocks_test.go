package blocks

import (
	"github.com/xavierzho/explorer-api"
	"testing"
)

var client = explorer.NewClient("YouApiKeyToken", explorer.Ethereum, nil)

var s = &Service{
	Client: client,
}

func TestGetBlockReward(t *testing.T) {
	res, err := s.BlockReward(2165403)
	if err != nil {
		t.Error(err)
	}
	if res.BlockMiner != "0x13a06d3dfe21e0db5c016c03ea7d2509f7f8d1e3" {
		t.Error("wrong block miner")
	}
	if res.UncleInclusionReward != "312500000000000000" {
		t.Error("wrong uncle inclusion reward")
	}
}

func TestGetCountdown(t *testing.T) {
	res, err := s.BlockCountDown(16701588)
	if err != nil {
		t.Error(err)
	}
	if res.CurrentBlock != "12715477" {
		t.Error("wrong current block")
	}
	if res.CountdownBlock != "16701588" {
		t.Error("wrong countdown block")
	}
	if res.RemainingBlock != "3986111" {
		t.Error("wrong remaining block")
	}
	if res.EstimateTimeInSec != "52616680.2" {
		t.Error("wrong estimate time in sec")
	}
}

func TestGetTimestamp(t *testing.T) {
	res, err := s.BlockNoByTimestamp(1578638524, "before")
	if err != nil {
		t.Error(err)
	}
	if res != "12712551" {
		t.Error("wrong block number")
	}
}
