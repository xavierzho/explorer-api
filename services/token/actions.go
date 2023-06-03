package token

import "github.com/Jonescy/explorer-api"

type Action interface {
	explorer.Endpoint
	ERC20TotalSupply(address string) (totalSupply string, err error)
}
