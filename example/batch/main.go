package main

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/xavierzho/explorer-api"
)

func main() {
	client := explorer.NewClient(
		"Y3BFMKAY6K8CZGYC7Z7QSE2MBMSD7FIM5A",
		explorer.Ethereum,
		explorer.ClientWithRTLimiter(5, 3),
	)

	//client.AfterHook = func(ctx context.Context, body []byte) {
	//	fmt.Println("after hook")
	//	// verbose the response body
	//	fmt.Println("body", string(body))
	//}
	var wg sync.WaitGroup
	service := client.Accounts()
	success := 0
	mux := sync.Mutex{}
	// using the same client to call the same endpoint concurrently, testing the rate limit
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			balance, err := service.EtherBalance(common.HexToAddress("0xBE0eB53F46cd790Cd13851d5EFf43D12404d33E8"))
			if err != nil {
				fmt.Println("has some error", err)
				return
			}
			mux.Lock()
			success++
			mux.Unlock()
			println("balance", idx, balance.Int().String())
		}(i)
	}
	wg.Wait()
	fmt.Println("success", success)
}
