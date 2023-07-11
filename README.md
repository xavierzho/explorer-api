# Etherscan explorer API

[![GoDoc](https://godoc.org/github.com/Jonescy/explorer-api?status.svg)](https://godoc.org/github.com/Jonescy/explorer-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/Jonescy/explorer-api)](https://goreportcard.com/report/github.com/Jonescy/explorer-api)
![License](https://img.shields.io/github/license/Jonescy/explorer-api.svg)

Golang client for the Etherscan.io API(and its families like BscScan). with nearly full implementation(accounts, transactions, tokens, contracts, blocks, stats), full network support(Mainnet, Ropsten, Kovan, Rinkby, Goerli, Tobalaba), and depending on `go-ethereum/common`.


# Usage

```bash
go get github.com/Jonescy/explorer-api
```

Create an API instance and off you go.

```go
package main

import (
	"github.com/Jonescy/explorer-api/modules/logs"
	"net/http"
	"context"

	"github.com/Jonescy/explorer-api"
	"github.com/Jonescy/explorer-api/modules/accounts"
)

func main() {
	// create a api client
	client := explorer.NewClient(
		// added your api key here
		explorer.WithAPIKey("YourApiKeyToken"),
		// using the different network
		explorer.WithBaseURL(explorer.Ethereum),
		// sub one to avoid hitting the limit
		explorer.WithLimitTier(explorer.TierFree-1),
		// custom http client
		explorer.WithHTTPClient(http.DefaultClient),
		// custom http request timeout
		explorer.WithTimeout(0),
	)
	// (optional) add hooks. useful for logging, metrics, etc.
	client.BeforeHook = func(ctx context.Context, url string) error {
		// ...
		return nil
	}
	client.AfterHook = func(ctx context.Context, url string, err error) {
		// ...
	}
	// arbitrary module dependency injection client
	service := accounts.Service{Client: client}
	
	log := logs.Service{Client: client}
	// and so on...
	// get account balance
	balance, _ := service.EtherBalance("0xBE0eB53F46cd790Cd13851d5EFf43D12404d33E8")

}

```
other usage example on `examples/` folder.

You may find full method list at [GoDoc](https://pkg.go.dev/github.com/Jonescy/explorer-api).

# Etherscan API Key

You may apply for an API key on [etherscan](https://etherscan.io/apis).

> The Etherscan Ethereum Developer APIs are provided as a community service and without warranty, so please just use what you need and no more. They support both GET/POST requests and a rate limit of 5 requests/sec (exceed and you will be blocked). 

