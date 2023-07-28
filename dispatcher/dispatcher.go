package dispatcher

import (
	"github.com/xavierzho/explorer-api"
	"sync"
)

type Dispatcher struct {
	clients  []*explorer.Client
	len      int
	roulette int
	mux      sync.RWMutex
}

// New create a dispatcher with clients
// need keys calc formula = qps / Tier
func New(keys []string) *Dispatcher {
	if len(keys) < 1 {
		return nil
	}
	var dispatcher = new(Dispatcher)
	dispatcher.clients = make([]*explorer.Client, 0, len(keys))
	for _, key := range keys {
		cli := explorer.NewClient(
			explorer.WithAPIKey(key),
			explorer.WithLimitTier(explorer.TierFree),
			explorer.WithBaseURL(explorer.Ethereum),
		)
		dispatcher.Append(cli)
	}
	return dispatcher
}

func (d *Dispatcher) Append(client *explorer.Client) {
	d.mux.Lock()
	defer d.mux.Unlock()
	d.clients = append(d.clients, client)
	d.len++
}

// Next return a client
func (d *Dispatcher) Next() *explorer.Client {
	d.mux.RLocker()
	defer d.mux.RUnlock()
	if d.len == 0 {
		return nil
	}
	if d.len == 1 {
		return d.clients[0]
	}
	client := d.clients[d.roulette%d.len]
	d.roulette++
	if d.roulette >= d.len {
		d.roulette = 0
	}
	return client
}
