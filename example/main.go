package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Jonescy/explorer-api"
	"github.com/Jonescy/explorer-api/modules/accounts"
)

func main() {
	client := explorer.NewClient(
		explorer.WithAPIKey("YourApiKeyToken"),
		explorer.WithBaseURL(explorer.Ethereum),
		// sub one to avoid hitting the limit
		explorer.WithLimitTier(explorer.TierFree-1),
		explorer.WithHTTPClient(http.DefaultClient),
		explorer.WithTimeout(0),
	)
	var wg sync.WaitGroup
	service := accounts.Service{Client: client}
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
