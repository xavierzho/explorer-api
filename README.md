# Etherscan explorer API

[![GoDoc](https://godoc.org/github.com/xavierzho/explorer-api?status.svg)](https://godoc.org/github.com/xavierzho/explorer-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/xavierzho/explorer-api)](https://goreportcard.com/report/github.com/xavierzho/explorer-api)
![License](https://img.shields.io/github/license/xavierzho/explorer-api.svg)

Golang client for the Etherscan.io API(and its families like BscScan).

# Features
1. Full implementation of Etherscan API, such as (accounts, transactions, tokens, contracts, blocks, stats)
2. network support (Mainnet, Ropsten, Kovan, Rinkby, Goerli, BNBChain, Poloygan, etc.)
3. depending on `go-ethereum/common`, so you can use `common.Address` directly.
4. support rate limit, you can set the rate limit by yourself.
5. support hooks, you can add hooks to do something before or after the request.
6. combining current limiters to support concurrency.


# Usage

```bash
go get github.com/xavierzho/explorer-api
```

Create an API instance and off you go.

```go
package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/xavierzho/explorer-api/modules/logs"

	"github.com/xavierzho/explorer-api"
	"github.com/xavierzho/explorer-api/modules/accounts"
)

func main() {
	// create a api client
	client := explorer.NewClient(
		"<you api key here>",
		explorer.Ethereum,
		explorer.ClientWithRTLimiter(5, 3),
	)
	// (optional) add hooks. useful for logging, metrics, etc.
	client.BeforeHook = func(ctx context.Context, url string) error {
		// ...
		return nil
	}
	client.AfterHook = func(ctx context.Context, body []byte) {
		// ...
	}
	// arbitrary module dependency injection client
	service := accounts.Service{Client: client}

	log := logs.Service{Client: client}
	// and so on...
	_ = log
	// get account balance
	balance, _ := service.EtherBalance(common.HexToAddress("0xBE0eB53F46cd790Cd13851d5EFf43D12404d33E8"))
	fmt.Println(balance)

}

```
other usage example on `examples/` folder.

You may find full method list at [GoDoc](https://pkg.go.dev/github.com/xavierzho/explorer-api).

# Etherscan API Key

You may apply for an API key on [etherscan](https://etherscan.io/apis).

> The Etherscan Ethereum Developer APIs are provided as a community service and without warranty, so please just use what you need and no more. They support both GET/POST requests and a rate limit of 5 requests/sec (exceed and you will be blocked). 

