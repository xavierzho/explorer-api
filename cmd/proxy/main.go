package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"github.com/xavierzho/explorer-api"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var lock sync.Locker = &sync.Mutex{}

type Key struct {
	list []string
	cur  int
}
type Keys map[explorer.Network]*Key

func (key *Key) Next() string {
	lock.Lock()
	defer lock.Unlock()
	key.cur++
	if key.cur >= len(key.list) {
		key.cur = 0
	}
	return key.list[key.cur]
}

var keys = make(Keys)

func init() {

	configPath := flag.String("k", "~/.config/keys.json", "api key config file.")
	flag.Parse()
	f, err := os.Open(*configPath)
	if err != nil {
		slog.Error("please use ~/.config/keys.json or -k cli flag")
		os.Exit(1)
	}
	defer f.Close()
	var data = make(map[string][]string)
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		slog.Error("Unmarshal config.")
		os.Exit(1)
	}

	for n, list := range data {
		var network explorer.Network
		switch n {
		case "bscscan":
			network = explorer.BinanceSmartChain
		case "bscscan-test":
			network = explorer.BinanceTestnet
		case "etherscan":
			network = explorer.Ethereum
		case "goerli":
			network = explorer.GoerliTestnet
		case "sepolia":
			network = explorer.SepoliaTestnet
		case "polygonscan":
			network = explorer.Polygon
		case "arbiscan":
			network = explorer.Arbitrum
		}

		keys[network] = &Key{
			list: list,
			cur:  0,
		}
	}

}

type ScanProxy struct{}

func (*ScanProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := r.Clone(ctx)
	// 设置X-Forward-For 头部
	if clientIp, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		if prior, ok := req.Header["X-Forward-For"]; ok {
			clientIp = strings.Join(prior, ", ") + ", " + clientIp
		}
		req.Header.Set("X-Forward-For", clientIp)
	}
	var dstNetwork explorer.Network
	switch req.URL.Path {
	case "/bscscan":
		dstNetwork = explorer.BinanceSmartChain
	case "/bscscan-test":
		dstNetwork = explorer.BinanceTestnet
	case "/etherscan":
		dstNetwork = explorer.Ethereum
	case "/goerli":
		dstNetwork = explorer.GoerliTestnet
	case "/sepolia":
		dstNetwork = explorer.SepoliaTestnet
	case "/polygonscan":
		dstNetwork = explorer.Polygon
	case "/arbiscan":
		dstNetwork = explorer.Arbitrum
	default:
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("unsupported scan"))
		return
	}
	url := dstNetwork.URL()
	q := req.URL.Query()
	q.Set("apikey", keys[dstNetwork].Next())
	url.RawQuery = q.Encode()
	req.URL = url
	req.Host = url.Host
	// 构造新请求
	response, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bytes.NewBufferString("round trip error").Bytes())
		slog.Error("RoundTrip", err)
		return
	}

	// 获取响应数据并返回
	for k, v := range response.Header {
		for _, v1 := range v {
			w.Header().Add(k, v1)
		}
	}
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}
func main() {
	log.Fatalln(http.ListenAndServe(":9000", &ScanProxy{}))
}
