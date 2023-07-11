package proxy

import (
	"fmt"
	"github.com/xavierzho/explorer-api"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/xavierzho/explorer-api/utils"
)

var client = explorer.NewClient(explorer.WithAPIKey("3T44IRF46WQUF4MQXIE2AYIFRSGCTNYAMS"),
	explorer.WithLimitTier(4))
var s = &Service{
	Client: client,
}

func TestNumberString2Hex(t *testing.T) {
	var s = "88888"
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {

	}
	fmt.Println(i)
	var h = fmt.Sprintf("%s", s)
	fmt.Println(h)
}

func TestGetBlockNumber(t *testing.T) {
	res, err := s.GetBlockNumber()
	if err != nil {
		t.Error(err)
	}
	t.Log(res.Int().String())
}

func TestGetBlockByNumber(t *testing.T) {

	res, err := s.GetBlockByNumber(utils.NewBNFromInt64(17513464))
	if err != nil {
		t.Error(err)
	}
	t.Log(res.Miner.String())
}

func TestGetUncleByBlockNo(t *testing.T) {
	res, err := s.GetUncleByBlockNumberAndIndex(utils.NewBNFromInt64(12989046), utils.NewBNFromInt64(0))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestGetBlockTxCount(t *testing.T) {
	res, err := s.GetTxCountByBlockNumber(utils.NewBNFromInt64(1112952))
	if err != nil {
		t.Error(err)
	}
	t.Log(res.Int().Text(16))
}

func TestGetTxByHash(t *testing.T) {
	res, err := s.GetTxByHash(common.HexToHash("0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44"))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestTxByNoAndIndex(t *testing.T) {
	res, err := s.GetTxByBlockNumberAndIndex(utils.NewBNFromInt64(12989213), utils.NewBNFromInt64(282))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestTxCountByAddress(t *testing.T) {
	res, err := s.GetTxCountByAddress(common.HexToAddress("0x4bd5900Cb274ef15b153066D736bf3e83A9ba44e"))
	if err != nil {
		t.Error(err)
	}
	t.Log(res.Int().String())
}

func TestGetTxReceipt(t *testing.T) {
	res, err := s.GetTxReceiptByHash(common.HexToHash("0xadb8aec59e80db99811ac4a0235efa3e45da32928bcff557998552250fa672eb"))
	if err != nil {
		t.Error(err)
	}
	t.Log(res.Logs[0].Topics)
}

func TestEthCall(t *testing.T) {
	res, err := s.EthCall(common.HexToAddress("0xAEEF46DB4855E25702F8237E8f403FddcaF931C0"), []byte("0x70a08231000000000000000000000000e16359506c028e51f16be38986ec5746251e9724"), utils.NewBNFromInt64(0))
	if err != nil {
		t.Error(err)
	}
	if res != "0x00000000000000000000000000000000000000000000000000601d8888141c00" {
		t.Error("not equal")
	}
}

func TestGetCode(t *testing.T) {
	res, err := s.GetCode(common.HexToAddress("0xf75e354c5edc8efed9b59ee9f67a80845ade7d0c"), utils.NewBNFromInt64(0))
	if err != nil {
		t.Error(err)
	}
	if res != "0x3660008037602060003660003473273930d21e01ee25e4c219b63259d214872220a261235a5a03f21560015760206000f3" {
		t.Error("not equal")
	}
}
func TestGetStorageAt(t *testing.T) {
	res, err := s.GetStorageAt(common.HexToAddress("0x6e03d9cce9d60f3e9f2597e13cd4c54c55330cfd"), utils.NewBNFromInt64(0), utils.NewBNFromInt64(0))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestGetGasPrice(t *testing.T) {
	res, err := s.GetGasPrice()
	if err != nil {
		t.Error(err)
	}
	t.Log(res.Int().String())
}

func TestGetEstimatedGas(t *testing.T) {
	res, err := s.GetEstimatedGas(common.HexToAddress("0xf0160428a8552ac9bb7e050d90eeade4ddd52843"), "0xff22", []byte("0x4e71d92d"), utils.NewBNFromInt64(99999999), utils.NewBNFromInt64(21971876044))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
