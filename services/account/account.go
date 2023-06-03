package account

import (
	"strconv"
	"strings"

	"github.com/Jonescy/explorer-api/services"
	"github.com/Jonescy/explorer-api/utils"
)

// Service is the module to access account related APIs.
// implement etherscan.Action
type Service services.Service

func (*Service) Name() string { return "account" }

// EtherBalance returns the balance of a given address.
func (s *Service) EtherBalance(address string) (balance *utils.BigInt, err error) {
	var m = make(map[string]string)
	m["address"] = address
	err = s.Client.Call(s.Name(), "balance", m, &balance)
	return
}
func (s *Service) MultiAccountBalance(addresses ...string) (balances []Balance, err error) {
	err = s.Client.Call(s.Name(), "balancemulti", map[string]string{
		"address": strings.Join(addresses, ","),
	}, &balances)

	return
}

func (s *Service) NormalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []NormalTx, err error) {
	var m = map[string]string{
		"address": address,
		"page":    strconv.Itoa(page),
		"offset":  strconv.Itoa(offset),
	}

	if desc {
		m["sort"] = "desc"
	} else {
		m["sort"] = "asc"
	}
	if startBlock != nil {
		m["startblock"] = strconv.Itoa(*startBlock)
	} else {
		m["startblock"] = "0"
	}
	if endBlock != nil {
		m["endblock"] = strconv.Itoa(*endBlock)
	}
	err = s.Client.Call(s.Name(), "txlist", m, &txs)

	return
}
func (s *Service) InternalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []InternalTx, err error) {
	var m = map[string]string{
		"address": address,
		"page":    strconv.Itoa(page),
		"offset":  strconv.Itoa(offset),
	}
	if startBlock != nil {
		m["startblock"] = strconv.Itoa(*startBlock)
	} else {
		m["startblock"] = "0"
	}
	if endBlock != nil {
		m["endblock"] = strconv.Itoa(*endBlock)
	}
	if desc {
		m["sort"] = "desc"
	} else {
		m["sort"] = "asc"
	}
	err = s.Client.Call(s.Name(), "txlistinternal", m, &txs)
	return
}
func (s *Service) InternalTxsByHash(txHash string) (txs []InternalTx, err error) {
	var m = map[string]string{
		"txhash": txHash,
	}
	err = s.Client.Call(s.Name(), "txlistinternal", m, &txs)
	return
}
func (s *Service) InternalTxsByBlockRange(startBlock int, endBlock int, page int, offset int, desc bool) (txs []InternalTx, err error) {
	var m = map[string]string{
		"startblock": strconv.Itoa(startBlock),
		"endblock":   strconv.Itoa(endBlock),
		"page":       strconv.Itoa(page),
		"offset":     strconv.Itoa(offset),
	}
	if desc {
		m["sort"] = "desc"
	} else {
		m["sort"] = "asc"
	}
	err = s.Client.Call(s.Name(), "txlistinternal", m, &txs)
	return
}
func (s *Service) ERC20TransferEventsByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC20Transfer, err error) {
	var m = map[string]string{
		"address": address,
		"page":    strconv.Itoa(page),
		"offset":  strconv.Itoa(offset),
	}
	if startBlock != nil {
		m["startblock"] = strconv.Itoa(*startBlock)
	} else {
		m["startblock"] = "0"
	}
	if endBlock != nil {
		m["endblock"] = strconv.Itoa(*endBlock)
	}
	if desc {
		m["sort"] = "desc"
	} else {
		m["sort"] = "asc"
	}
	err = s.Client.Call(s.Name(), "tokentx", m, &txs)
	return
}
func (s *Service) ERC721TransferEventsByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC721Transfer, err error) {
	var m = map[string]string{
		"address": address,
		"page":    strconv.Itoa(page),
		"offset":  strconv.Itoa(offset),
	}
	if startBlock != nil {
		m["startblock"] = strconv.Itoa(*startBlock)
	} else {
		m["startblock"] = "0"
	}
	if endBlock != nil {
		m["endblock"] = strconv.Itoa(*endBlock)
	}
	if desc {
		m["sort"] = "desc"
	} else {
		m["sort"] = "asc"
	}
	err = s.Client.Call(s.Name(), "tokennfttx", m, &txs)
	return
}
func (s *Service) ERC1155TransferEventsByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC1155Transfer, err error) {
	var m = map[string]string{
		"address": address,
		"page":    strconv.Itoa(page),
		"offset":  strconv.Itoa(offset),
	}
	if startBlock != nil {
		m["startblock"] = strconv.Itoa(*startBlock)
	} else {
		m["startblock"] = "0"
	}
	if endBlock != nil {
		m["endblock"] = strconv.Itoa(*endBlock)
	}
	if desc {
		m["sort"] = "desc"
	} else {
		m["sort"] = "asc"
	}
	err = s.Client.Call(s.Name(), "tokennfttx", m, &txs)
	return
}
func (s *Service) MinedBlocksByAddress(address string, page int, offset int) (blocks []MinedBlock, err error) {
	var m = map[string]string{
		"address":   address,
		"page":      strconv.Itoa(page),
		"offset":    strconv.Itoa(offset),
		"blocktype": "blocks",
	}
	err = s.Client.Call(s.Name(), "getminedblocks", m, &blocks)
	return
}
func (s *Service) HistoricalByBlockNo(address string, blockNo int) (balance HistoricalBalance, err error) {
	var m = map[string]string{
		"address": address,
		"blockno": strconv.Itoa(blockNo),
	}
	err = s.Client.Call(s.Name(), "balancehistory", m, &balance)
	return
}

func (s *Service) ERC20Balance(address string, contractAddress string) (balance *utils.BigInt, err error) {
	err = s.Client.Call(s.Name(), "tokenbalance", map[string]string{
		"address":         address,
		"contractaddress": contractAddress,
	}, &balance)
	return
}

func (s *Service) HistoricalERC20Balance(address string, contractAddress string, blockNo int) (balance HistoricalBalance, err error) {
	err = s.Client.Call(s.Name(), "tokenbalancehistory", map[string]string{
		"address":         address,
		"contractaddress": contractAddress,
		"blockno":         strconv.Itoa(blockNo),
	}, &balance)
	return
}
