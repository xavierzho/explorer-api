package contracts

import (
	"github.com/xavierzho/explorer-api/modules"
	"strings"
)

type Service modules.Service

func (s *Service) Name() string { return "contract" }

// GetABI Get Contract ABI for Verified Contract Source Codes
//
// description: https://docs.etherscan.io/api-endpoints/contracts#get-contract-abi-for-verified-contract-source-codes
//
// Returns the Contract Application Binary Interface ( ABI ) of a verified smart contract.
func (s *Service) GetABI(address string) (abi string, err error) {
	err = s.Client.Call(s, "getabi", map[string]string{
		"address": address,
	}, &abi)
	return
}

// GetSourceCode Get Contract Source Code for Verified Contract Source Codes
//
// description: https://docs.etherscan.io/api-endpoints/contracts#get-contract-source-code-for-verified-contract-source-codes
//
// Returns the Solidity source code of a verified smart contract.
func (s *Service) GetSourceCode(address string) (sourceCodes []SourceCode, err error) {
	err = s.Client.Call(s, "getsourcecode", map[string]string{
		"address": address,
	}, &sourceCodes)
	return
}

// GetContractCreation Get Contract Creator and Creation Tx Hash
//
// description:https://docs.etherscan.io/api-endpoints/contracts#get-contract-creator-and-creation-tx-hash
//
// Returns a contract's deployer address and transaction hash it was created, up to 5 at a time.
func (s *Service) GetContractCreation(address []string) (contractCreations []Creation, err error) {

	err = s.Client.Call(s, "getcontractcreation", map[string]string{
		"contractaddresses": strings.Join(address, ","),
	}, &contractCreations)
	return
}
