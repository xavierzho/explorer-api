package gastracker

import (
	"github.com/Jonescy/explorer-api"
	"testing"
)

var client = explorer.NewClient(explorer.WithAPIKey("YourApiKeyToken"),
	explorer.WithLimitTier(4))
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
