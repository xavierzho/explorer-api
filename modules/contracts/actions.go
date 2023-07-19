package contracts

import "github.com/xavierzho/explorer-api"

type Action interface {
	explorer.Module
	GetABI(address string) (abi string, err error)
	GetSourceCode(address string) (sourceCodes []SourceCode, err error)
	GetContractCreation(address []string) (contractCreations []Creation, err error)
}

type SourceCode struct {
	SourceCode           string `json:"SourceCode"`
	ABI                  string `json:"ABI"`
	SwarmSource          string `json:"SwarmSource"`
	CompilerVersion      string `json:"CompilerVersion"`
	EVMVersion           string `json:"EVMVersion"`
	Implementation       string `json:"Implementation"`
	ContractName         string `json:"ContractName"`
	Runs                 string `json:"Runs"`
	LicenseType          string `json:"LicenseType"`
	Proxy                string `json:"Proxy"`
	ConstructorArguments string `json:"ConstructorArguments"`
	Library              string `json:"Library"`
	OptimizationUsed     string `json:"OptimizationUsed"`
}
type Creation struct {
	ContractCreator string `json:"contractCreator"`
	ContractAddress string `json:"contractAddress"`
	TxHash          string `json:"txHash"`
}
