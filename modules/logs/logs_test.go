package logs

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xavierzho/explorer-api"
	"testing"
)

var client = explorer.NewClient(explorer.WithAPIKey("YourApiKeyToken"),
	explorer.WithLimitTier(4))
var s = &Service{
	Client: client,
}

func TestMarshalIntoMap(t *testing.T) {
	var param = Params{
		FromBlock:       "0x1",
		Topic0:          "0x2",
		Topic1:          "0x3",
		Topic01Operator: OperatorOr,
		Topic12Operator: OperatorAnd,
	}
	fmt.Println(param.MarshalIntoMap())
}

func TestGetLogs(t *testing.T) {
	res, err := s.GetLogs(Params{
		Address:         "0x59728544b08ab483533076417fbbb2fd0b17ce3a",
		FromBlock:       "15073139",
		ToBlock:         "15074139",
		Page:            "1",
		Offset:          "1000",
		Topic0:          "0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d",
		Topic01Operator: "and",
		Topic1:          "0x00000000000000000000000023581767a106ae21c074b2276d25e5c3e136a68b",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(len(res))
	t.Logf("%+v", res)
	if res[0].TransactionHash != common.HexToHash("0x26fe1a0a403fd44ef11ee72f3b4ceff590b6ea533684cb279cb4242be463304c") {
		t.Error("TransactionHash is not correct")
	}
}
