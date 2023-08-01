package accounts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/xavierzho/explorer-api"
	"math/big"
	"testing"
)

var client = explorer.NewClient("YouApiKeyToken", explorer.Ethereum, nil)
var s = &Service{
	Client: client,
}
var accounts = []string{
	"0xBE0eB53F46cd790Cd13851d5EFf43D12404d33E8",
}

func TestService_EtherBalance(t *testing.T) {

	if s.Name() != "account" {
		t.Error("Service.Name() != account")
	}
	// single account balance
	eb, err := s.EtherBalance(common.HexToAddress(accounts[0]))
	if err != nil {
		t.Error(err)
	}
	if eb.Int().Cmp(big.NewInt(0)) != 1 {
		t.Error("eb.Int().Cmp(big.NewInt(0)) != 1")
	}
	t.Logf("account %s balance of %s\n", accounts[0], eb.Int().String())
	// multi account balance
	bals, err := s.MultiAccountBalance(accounts)
	if err != nil {
		t.Error(err)
	}
	for _, bal := range bals {
		t.Logf("account %s balance of %s\n", bal.Account, bal.Balance.Int().String())
	}
}

func TestService_NormalTxByAddress(t *testing.T) {
	var endBlock int = 100000000000
	txs, err := s.NormalTxByAddress(common.HexToAddress(accounts[0]), nil, &endBlock, 1, 10, true)
	if err != nil {
		t.Error(err)
	}
	for _, tx := range txs {
		t.Logf("tx %+v\n", tx)
	}
}

func TestService_InternalTx(t *testing.T) {
	var endBlock int = 100000000000
	txs, err := s.InternalTxByAddress(common.HexToAddress(accounts[0]), nil, &endBlock, 1, 10, true)
	if err != nil {
		t.Error(err)
	}
	for _, tx := range txs {
		t.Logf("tx %+v\n", tx)
	}

	txs, err = s.InternalTxsByHash(common.HexToHash("0xfb0f81c11248669ecd85bd506cf4ad3e8d9ff9a550cb46f3a873b97dc9f8c481"))
	if err != nil {
		t.Error(err)
	}

	for _, tx := range txs {
		t.Logf("tx %+v\n", tx)
	}

	txs, err = s.InternalTxsByBlockRange(0, 16145605, 1, 10, true)
	if err != nil {
		t.Error(err)
	}
	t.Log("txs by block range len:", len(txs))
	for _, tx := range txs {
		t.Logf("tx %+v\n", tx)
	}

}

func TestService_TransferEvents(t *testing.T) {
	var endBlock int = 100000000000
	erc20, err := s.ERC20TransferEventsByAddress(common.HexToAddress(accounts[0]), nil, &endBlock, 1, 10, true)
	if err != nil {
		t.Error(err)
	}
	for _, tx := range erc20 {
		t.Logf("tx %+v\n", tx)
	}

	erc721, err := s.ERC721TransferEventsByAddress(common.HexToAddress(accounts[0]), nil, &endBlock, 1, 10, true)
	if err != nil {
		t.Error(err)
	}
	for _, tx := range erc721 {
		t.Logf("tx %+v\n", tx)
	}

	erc1155, err := s.ERC1155TransferEventsByAddress(common.HexToAddress(accounts[0]), nil, &endBlock, 1, 10, true)
	if err != nil {
		t.Error(err)
	}
	for _, tx := range erc1155 {
		t.Logf("tx %+v\n", tx)
	}
}

func TestService_MinedBlocksByAddress(t *testing.T) {
	blocks, err := s.MinedBlocksByAddress(common.HexToAddress("0x9dd134d14d1e65f84b706d6f205cd5b1cd03a46b"), 1, 10)
	if err != nil {
		t.Error(err)
	}
	for _, block := range blocks {
		t.Logf("block %+v\n", block)
	}
}

func TestService_HistoricalByBlockNo(t *testing.T) {
	block, err := s.HistoricalByBlockNo(common.HexToAddress("0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae"), 8000000)
	if err != nil {
		t.Error(err)
	}
	t.Logf("block %+v\n", block)
}
