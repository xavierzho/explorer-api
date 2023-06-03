package contract

import (
	"github.com/Jonescy/explorer-api/services"
)

type Service services.Service

func (s *Service) Name() string { return "contract" }
func (s *Service) GetABI(address string) (abi string, err error) {
	err = s.Client.Call(s.Name(), "getabi", map[string]string{
		"address": address,
	}, &abi)
	return
}

func (s *Service) GetSourceCode(address string) (sourceCodes []SourceCode, err error) {
	err = s.Client.Call(s.Name(), "getsourcecode", map[string]string{
		"address": address,
	}, &sourceCodes)
	return
}

func (s *Service) GetContractCreation(address string) (contractCreations []Creation, err error) {
	err = s.Client.Call(s.Name(), "getcontractcreation", nil, nil)
	return
}
