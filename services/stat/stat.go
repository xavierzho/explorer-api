package stat

import (
	"github.com/Jonescy/explorer-api/services"
	"github.com/Jonescy/explorer-api/utils"
)

type Service services.Service

func (*Service) Name() string { return "stats" }

func (s *Service) EthSupply() (supply *utils.BigInt, err error) {
	err = s.Client.Call(s.Name(), "ethsupply", nil, &supply)
	return
}

func (s *Service) Eth2Supply() (supply Supply, err error) {
	err = s.Client.Call(s.Name(), "eth2supply", nil, nil)
	return
}

func (s *Service) EthLastPrice() (prices Prices, err error) {
	err = s.Client.Call(s.Name(), "ethprice", nil, &prices)
	return
}

func (s *Service) ChainSize(startDate, endDate string, isAsc bool) (sizes []ChainSize, err error) {
	sort := "asc"
	if !isAsc {
		sort = "desc"
	}
	err = s.Client.Call(s.Name(), "chainsize", map[string]string{
		"startdate": startDate,
		"enddate":   endDate,
		"sort":      sort,
		"chaintype": "geth",
		"syncmode":  "default",
	}, &sizes)
	return
}

func (s *Service) NodeCount() (nodeCount NodeCount, err error) {
	err = s.Client.Call(s.Name(), "nodecount", nil, &nodeCount)
	return
}
