package stats

import (
	"github.com/xavierzho/explorer-api"
	"testing"
)

var client = explorer.NewClient("YouApiKeyToken", explorer.Ethereum, nil)

var s = &Service{
	Client: client,
}

func TestEthSupply(t *testing.T) {
	res, err := s.EthSupply()
	if err != nil {
		t.Error(err)
	}
	t.Log(res.Int().String())
}

func TestEth2Supply(t *testing.T) {
	res, err := s.Eth2Supply()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestEthLastPrice(t *testing.T) {
	res, err := s.EthLastPrice()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestEthNodeSize(t *testing.T) {
	sizes, err := s.GetNodeSize("2022-02-01", "2022-02-28", true)
	if err != nil {
		t.Error(err)
	}
	t.Log(sizes)
}

func TestNodeCount(t *testing.T) {
	res, err := s.NodeCount()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
