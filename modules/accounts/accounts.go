package accounts

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"strings"

	"github.com/xavierzho/explorer-api/iface"
	"github.com/xavierzho/explorer-api/utils"
)

// Service is the module to access account related APIs.
// implement etherscan.Blocks
type Service iface.Service

func (*Service) Name() string { return "account" }

// EtherBalance Get Ether Balance for a Single Address.
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-ether-balance-for-a-single-address
//
// Returns the Ether balance of a given address.
func (s *Service) EtherBalance(address common.Address) (balance *utils.BN, err error) {
	err = s.Client.Call(s, "balance", utils.M{
		"address": address.Hex(),
	}, &balance)
	return
}

// MultiAccountBalance Get Ether Balance for Multiple Addresses in a Single Call
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-ether-balance-for-multiple-addresses-in-a-single-call
//
// Returns the balance of the accounts from a list of addresses.
func (s *Service) MultiAccountBalance(addresses []string) (balances []iface.Balance, err error) {
	err = s.Client.Call(s, "balancemulti", utils.M{
		"address": strings.Join(addresses, ","),
	}, &balances)

	return
}

// NormalTxByAddress Get a list of 'Normal' Transactions By Address
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-a-list-of-normal-transactions-by-address
//
// Returns the list of NormalTx performed by an address, with optional pagination.
func (s *Service) NormalTxByAddress(address common.Address, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []iface.NormalTx, err error) {
	var m = utils.M{
		"address": address.Hex(),
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
	err = s.Client.Call(s, "txlist", m, &txs)

	return
}

// InternalTxByAddress Get a list of 'Internal' Transactions by Address
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-a-list-of-internal-transactions-by-address
//
// Returns the list of InternalTx performed by an address, with optional pagination.
func (s *Service) InternalTxByAddress(address common.Address, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []iface.InternalTx, err error) {
	var m = utils.M{
		"address": address.Hex(),
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
	err = s.Client.Call(s, "txlistinternal", m, &txs)
	return
}

// InternalTxsByHash Get 'Internal Transactions' by Transaction Hash
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-internal-transactions-by-transaction-hash
//
// Returns the list of InternalTx performed within a transaction.
func (s *Service) InternalTxsByHash(txHash common.Hash) (txs []iface.InternalTx, err error) {
	var m = utils.M{
		"txhash": txHash.Hex(),
	}
	err = s.Client.Call(s, "txlistinternal", m, &txs)
	return
}

// InternalTxsByBlockRange Get "Internal Transactions" by Block Range
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-internal-transactions-by-block-range
//
// Returns the list of internal transactions performed within a block range, with optional pagination.
func (s *Service) InternalTxsByBlockRange(startBlock int, endBlock int, page int, offset int, desc bool) (txs []iface.InternalTx, err error) {
	var m = utils.M{
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
	err = s.Client.Call(s, "txlistinternal", m, &txs)
	return
}

// ERC20TransferEventsByAddress Get a list of 'ERC20 - Token Transfer Events' by Address
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-a-list-of-erc20-token-transfer-events-by-address
//
// Returns the list of ERC20Transfer by an address, with optional filtering by token contract.
func (s *Service) ERC20TransferEventsByAddress(address common.Address, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []iface.ERC20Transfer, err error) {
	var m = utils.M{
		"address": address.Hex(),
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
	err = s.Client.Call(s, "tokentx", m, &txs)
	return
}

// ERC721TransferEventsByAddress Get a list of 'ERC721 - Token Transfer Events' by Address
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-a-list-of-erc721-token-transfer-events-by-address
//
// Returns the list of ERC721Transfer by an address, with optional filtering by token contract.
func (s *Service) ERC721TransferEventsByAddress(address common.Address, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []iface.ERC721Transfer, err error) {
	var m = utils.M{
		"address": address.Hex(),
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
	err = s.Client.Call(s, "tokennfttx", m, &txs)
	return
}

// ERC1155TransferEventsByAddress Get a list of 'ERC1155 - Token Transfer Events' by Address
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-a-list-of-erc1155-token-transfer-events-by-address
//
// Returns the list of ERC1155Transfer by an address, with optional filtering by token contract.
func (s *Service) ERC1155TransferEventsByAddress(address common.Address, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []iface.ERC1155Transfer, err error) {
	var m = utils.M{
		"address": address.Hex(),
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
	err = s.Client.Call(s, "tokennfttx", m, &txs)
	return
}

// MinedBlocksByAddress Get a list of 'Mined Blocks' by Address
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-a-list-of-mined-blocks-by-address
//
// Returns the list of MinedBlock by an address, with optional pagination.
func (s *Service) MinedBlocksByAddress(address common.Address, page int, offset int) (blocks []iface.MinedBlock, err error) {
	var m = utils.M{
		"address":   address.Hex(),
		"page":      strconv.Itoa(page),
		"offset":    strconv.Itoa(offset),
		"blocktype": "blocks",
	}
	err = s.Client.Call(s, "getminedblocks", m, &blocks)
	return
}

// HistoricalByBlockNo Get Historical Ether Balance for a Single Address By BlockNo. [Pro]
//
// description: https://docs.etherscan.io/api-endpoints/accounts#get-historical-ether-balance-for-a-single-address-by-blockno
//
// Returns the balance of an address at a given block number.
func (s *Service) HistoricalByBlockNo(address common.Address, blockNo int) (balance iface.HistoricalBalance, err error) {
	var m = utils.M{
		"address": address.Hex(),
		"blockno": strconv.Itoa(blockNo),
	}
	err = s.Client.Call(s, "balancehistory", m, &balance)
	return
}

// ERC20Balance Get ERC20-Token Account Balance for TokenContractAddress
//
// description: https://docs.etherscan.io/api-endpoints/tokens#get-erc20-token-account-balance-for-tokencontractaddress
//
// Returns the current balance of an ERC-20 token of an address.
func (s *Service) ERC20Balance(address common.Address, contractAddress string) (balance *utils.BN, err error) {
	err = s.Client.Call(s, "tokenbalance", utils.M{
		"address":         address.Hex(),
		"contractaddress": contractAddress,
	}, &balance)
	return
}

// HistoricalERC20Balance Get Historical ERC20-Token Account Balance for TokenContractAddress by BlockNo. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/tokens#get-historical-erc20-token-account-balance-for-tokencontractaddress-by-blockno
//
// Returns the balance of an ERC-20 token of an address at a certain block height.
func (s *Service) HistoricalERC20Balance(address common.Address, contractAddress common.Address, blockNo int) (balance iface.HistoricalBalance, err error) {
	err = s.Client.Call(s, "tokenbalancehistory", utils.M{
		"address":         address.Hex(),
		"contractaddress": contractAddress.Hex(),
		"blockno":         strconv.Itoa(blockNo),
	}, &balance)
	return
}

// GetHoldingERC20Tokens Get Address ERC20 Token Holding. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/tokens#get-address-erc20-token-holding
//
// Returns the ERC-20 tokens and amount held by an address.
func (s *Service) GetHoldingERC20Tokens(address common.Address, page, offset *int) (tokens []iface.HoldingERC20, err error) {
	if page == nil {
		*page = 1
	}
	if offset == nil {
		*offset = 100
	}
	err = s.Client.Call(s, "addresstokenbalance", utils.M{
		"address": address.Hex(),
		"page":    strconv.Itoa(*page),
		"offset":  strconv.Itoa(*offset),
	}, &tokens)
	return
}

// GetHoldingERC721Tokens Get Address ERC721 Token Holding. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/tokens#get-address-erc721-token-holding
//
// Returns the ERC-721 tokens and amount held by an address.
func (s *Service) GetHoldingERC721Tokens(address common.Address, page, offset *int) (tokens []iface.HoldingERC721, err error) {
	if page == nil {
		*page = 1
	}
	if offset == nil {
		*offset = 100
	}
	err = s.Client.Call(s, "addresstokennftbalance", utils.M{
		"address": address.Hex(),
		"page":    strconv.Itoa(*page),
		"offset":  strconv.Itoa(*offset),
	}, &tokens)
	return
}

// GetERC721Inventory Get Address ERC721 Token Inventory By Contract Address. [PRO]
//
// description: https://docs.etherscan.io/api-endpoints/tokens#get-address-erc721-token-inventory-by-contractaddress
//
// Returns the ERC-721 token inventory of an address, filtered by contract address.
func (s *Service) GetERC721Inventory(wallet, contract common.Address, page, offset *int) (tokens []iface.ERC721Inventory, err error) {
	if page == nil {
		*page = 1
	}
	if offset == nil {
		*offset = 100
	}
	err = s.Client.Call(s, "addresstokennftinventory", utils.M{
		"address":         wallet.Hex(),
		"page":            strconv.Itoa(*page),
		"offset":          strconv.Itoa(*offset),
		"contractaddress": contract.Hex(),
	}, &tokens)
	return
}
