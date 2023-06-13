package gastracker

import (
	"github.com/Jonescy/explorer-api"
)

type Service explorer.Service

// Name returns the name of this service.
func (s *Service) Name() string { return "gastracker" }

// GasEstimate returns the current gas price oracle.
func (s *Service) GasEstimate(gasprice string) (gas string, err error) {
	err = s.Client.Call(s.Name(), "gasestimate", map[string]string{
		"gasprice": gasprice,
	}, nil)
	return
}

// GasOracle returns the current gas price oracle.
func (s *Service) GasOracle() (gasOracle GasOracle, err error) {
	err = s.Client.Call(s.Name(), "gasoracle", nil, &gasOracle)
	return
}
