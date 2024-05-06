package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/xavierzho/explorer-api"
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

	// after v1.3
	service := client.Accounts()
	// compatible old version
	//service := accounts.Service{Client: client}

	//log := logs.Service{Client: client}
	// and so on...
	//_ = log
	// get account balance
	balance, _ := service.EtherBalance(common.HexToAddress("0xBE0eB53F46cd790Cd13851d5EFf43D12404d33E8"))
	fmt.Println(balance)

}
