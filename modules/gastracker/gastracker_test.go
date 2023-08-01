package gastracker

import (
	"github.com/xavierzho/explorer-api"
	"testing"
)

var client = explorer.NewClient("YouApiKeyToken", explorer.Ethereum, nil)

var s = &Service{
	Client: client,
}

func TestGasEstimate(t *testing.T) {
	res, err := s.GasEstimate("2000000000")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestGasOracle(t *testing.T) {
	res, err := s.GasOracle()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
