package dispatcher

import (
	"github.com/xavierzho/explorer-api"
	"sync"
	"testing"
)

func TestDispatcher_Next(t *testing.T) {
	var dispatcher = New([]string{"11", "22", "33", "44", "55"}, explorer.Ethereum, explorer.ClientWithRTLimiter(5, 3))
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			t.Logf("%+v->%+v\n", dispatcher.Next().APIKey, dispatcher.roulette)
			wg.Done()
		}()
	}
	wg.Wait()
}
