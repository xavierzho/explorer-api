package main

import (
	"fmt"
	"sync"

	"github.com/Jonescy/explorer-api/config"
	"github.com/Jonescy/explorer-api/core"
	"github.com/Jonescy/explorer-api/services/account"
)

func main() {
	client := core.NewClient(
		core.WithAPIKey("YourApiKeyToken"),
		core.WithBaseURL(config.Ethereum),
		// sub one to avoid hitting the limit
		core.WithLimitTier(config.TierFree-1))
	var wg sync.WaitGroup
	service := account.Service{Client: client}
	// using the same client to call the same endpoint concurrently, testing the rate limit
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			balance, err := service.EtherBalance("0xBE0eB53F46cd790Cd13851d5EFf43D12404d33E8")
			if err != nil {
				fmt.Println("has some error", err)
				return
			}

			println("balance", idx, balance.Int().String())
		}(i)
	}
	wg.Wait()
}
