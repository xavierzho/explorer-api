package main

import (
	"context"
	"fmt"
	"github.com/xavierzho/explorer-api"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/common"

	"github.com/xavierzho/explorer-api/modules/accounts"
)

func main() {
	client := explorer.NewClient(
		explorer.WithAPIKey("Y3BFMKAY6K8CZGYC7Z7QSE2MBMSD7FIM5A"),
		explorer.WithBaseURL(explorer.Ethereum),
		// sub one to avoid hitting the limit
		explorer.WithLimitTier(explorer.TierFree-1),
		explorer.WithHTTPClient(http.DefaultClient),
		explorer.WithTimeout(0),
	)

	client.AfterHook = func(ctx context.Context, body []byte) {
		fmt.Println("after hook")
		// verbose the response body
		fmt.Println("body", string(body))
	}
	var wg sync.WaitGroup
	service := accounts.Service{Client: client}
	// using the same client to call the same endpoint concurrently, testing the rate limit
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			balance, err := service.EtherBalance(common.HexToAddress("0xBE0eB53F46cd790Cd13851d5EFf43D12404d33E8"))
			if err != nil {
				//fmt.Println("has some error", err)
				return
			}

			println("balance", idx, balance.Int().String())
		}(i)
	}
	wg.Wait()
}
