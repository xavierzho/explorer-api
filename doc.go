// Package explorer provides Go bindings to the Etherscan.io API.
//
// This work is a nearly Full implementation
// (accounts, transactions, tokens, contracts, blocks, stats),
// with full network support
// (Mainnet, Ropsten, Kovan, Rinkby, Tobalaba, BSC, BSC Testnet, Polygon, Fantom, Arbitrum)
// and only depending on standard library.
//
// Example can be found at https://github.com/xavierzho/etherscan-api
package explorer

import "fmt"

const Version = "v1.2.0"

func init() {
	fmt.Println("explorer-api version:", Version)
}
