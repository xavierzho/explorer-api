package accounts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/xavierzho/explorer-api"
	"github.com/xavierzho/explorer-api/utils"
)

type ProAction interface {
	// HistoricalByBlockNo gets historical block information by block number
	HistoricalByBlockNo(address common.Address, blockNo int) (balance HistoricalBalance, err error)
	HistoricalERC20Balance(address, contractAddress common.Address, blockNo int) (balance HistoricalBalance, err error)
	GetHoldingERC20Tokens(address common.Address, page, offset *int) (tokens []HoldingERC20, err error)
	GetHoldingERC721Tokens(address common.Address, page, offset *int) (tokens []HoldingERC721, err error)
	GetERC721Inventory(wallet, contract common.Address, page, offset *int) (tokens []ERC721Inventory, err error)
}

type Action interface {
	explorer.Module
	// EtherBalance gets ether balance for a single address
	EtherBalance(address common.Address) (balance *utils.BN, err error)
	// MultiAccountBalance gets ether balance for multiple addresses in a single call
	MultiAccountBalance(addresses []string) (balances []Balance, err error)
	// NormalTxByAddress gets a list of "normal" transactions by address
	NormalTxByAddress(address common.Address, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []NormalTx, err error)
	// InternalTxByAddress gets a list of "internal" transactions by address
	InternalTxByAddress(address common.Address, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []InternalTx, err error)
	// InternalTxsByHash gets a list of "internal" transactions by transaction hash
	InternalTxsByHash(txHash common.Hash) (txs []InternalTx, err error)
	// InternalTxsByBlockRange gets a list of "internal" transactions by block range
	InternalTxsByBlockRange(startBlock int, endBlock int, page int, offset int, desc bool) (txs []InternalTx, err error)
	// ERC20TransferEventsByAddress gets a list of ERC20 token transfer events by address
	ERC20TransferEventsByAddress(address common.Address, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC20Transfer, err error)
	// ERC721TransferEventsByAddress gets a list of ERC721 token transfer events by address
	ERC721TransferEventsByAddress(address common.Address, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC721Transfer, err error)
	// ERC1155TransferEventsByAddress gets a list of ERC1155 token transfer events by address
	ERC1155TransferEventsByAddress(address common.Address, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC1155Transfer, err error)
	// MinedBlocksByAddress gets a list of mined blocks by address
	MinedBlocksByAddress(address common.Address, page int, offset int) (blocks []MinedBlock, err error)

	ERC20Balance(address common.Address, contractAddress string) (balance *utils.BN, err error)
	ProAction
}

// Balance account and its balance in pair
type Balance struct {
	Account string    `json:"account"`
	Balance *utils.BN `json:"balance"`
}

// NormalTx holds info from normal tx query
type NormalTx struct {
	BlockNumber       utils.BN   `json:"blockNumber"`
	TimeStamp         utils.Time `json:"timeStamp"`
	Hash              string     `json:"hash"`
	Nonce             utils.BN   `json:"nonce"`
	BlockHash         string     `json:"blockHash"`
	TransactionIndex  string     `json:"transactionIndex"`
	From              string     `json:"from"`
	To                string     `json:"to"`
	Value             utils.BN   `json:"value"`
	Gas               int        `json:"gas,string"`
	GasPrice          utils.BN   `json:"gasPrice"`
	IsError           string     `json:"isError"`
	TxReceiptStatus   string     `json:"txreceipt_status"`
	Input             string     `json:"input"`
	ContractAddress   string     `json:"contractAddress"`
	CumulativeGasUsed utils.BN   `json:"cumulativeGasUsed"`
	GasUsed           utils.BN   `json:"gasUsed"`
	Confirmations     utils.BN   `json:"confirmations"`
}

// InternalTx holds info from internal tx query
type InternalTx struct {
	BlockNumber     utils.BN   `json:"blockNumber"`
	TimeStamp       utils.Time `json:"timeStamp"`
	Hash            string     `json:"hash"`
	From            string     `json:"from"`
	To              string     `json:"to"`
	Value           utils.BN   `json:"value"`
	ContractAddress string     `json:"contractAddress"`
	Input           string     `json:"input"`
	Type            string     `json:"type"`
	Gas             utils.BN   `json:"gas"`
	GasUsed         utils.BN   `json:"gasUsed"`
	TraceID         string     `json:"traceId"`
	IsError         string     `json:"isError"`
	ErrCode         string     `json:"errCode"`
}

// ERC20Transfer holds info from ERC20 token transfer event query
type ERC20Transfer struct {
	BlockNumber       utils.BN   `json:"blockNumber"`
	TimeStamp         utils.Time `json:"timeStamp"`
	Hash              string     `json:"hash"`
	Nonce             utils.BN   `json:"nonce"`
	BlockHash         string     `json:"blockHash"`
	From              string     `json:"from"`
	ContractAddress   string     `json:"contractAddress"`
	To                string     `json:"to"`
	Value             utils.BN   `json:"value"`
	TokenName         string     `json:"tokenName"`
	TokenSymbol       string     `json:"tokenSymbol"`
	TokenDecimal      uint8      `json:"tokenDecimal,string"`
	TransactionIndex  int        `json:"transactionIndex,string"`
	Gas               int        `json:"gas,string"`
	GasPrice          utils.BN   `json:"gasPrice"`
	GasUsed           int        `json:"gasUsed,string"`
	CumulativeGasUsed int        `json:"cumulativeGasUsed,string"`
	Input             string     `json:"input"`
	Confirmations     int        `json:"confirmations,string"`
}

// ERC721Transfer holds info from ERC721 token transfer event query
type ERC721Transfer struct {
	BlockNumber       int        `json:"blockNumber,string"`
	TimeStamp         utils.Time `json:"timeStamp"`
	Hash              string     `json:"hash"`
	Nonce             int        `json:"nonce,string"`
	BlockHash         string     `json:"blockHash"`
	From              string     `json:"from"`
	ContractAddress   string     `json:"contractAddress"`
	To                string     `json:"to"`
	TokenID           utils.BN   `json:"tokenID"`
	TokenName         string     `json:"tokenName"`
	TokenSymbol       string     `json:"tokenSymbol"`
	TokenDecimal      uint8      `json:"tokenDecimal,string"`
	TransactionIndex  int        `json:"transactionIndex,string"`
	Gas               int        `json:"gas,string"`
	GasPrice          utils.BN   `json:"gasPrice"`
	GasUsed           int        `json:"gasUsed,string"`
	CumulativeGasUsed int        `json:"cumulativeGasUsed,string"`
	Input             string     `json:"input"`
	Confirmations     int        `json:"confirmations,string"`
}

// ERC1155Transfer holds info from ERC1155 token transfer event query
type ERC1155Transfer struct {
	BlockNumber       int            `json:"blockNumber,string"`
	TimeStamp         utils.Time     `json:"timeStamp"`
	Hash              common.Hash    `json:"hash"`
	Nonce             int            `json:"nonce,string"`
	BlockHash         common.Hash    `json:"blockHash"`
	From              common.Address `json:"from"`
	ContractAddress   string         `json:"contractAddress"`
	To                common.Address `json:"to"`
	TokenID           utils.BN       `json:"tokenID"`
	TokenName         string         `json:"tokenName"`
	TokenSymbol       string         `json:"tokenSymbol"`
	TokenDecimal      uint8          `json:"tokenDecimal,string"`
	TokenValue        uint8          `json:"tokenValue,string"`
	TransactionIndex  int            `json:"transactionIndex,string"`
	Gas               utils.BN       `json:"gas"`
	GasPrice          utils.BN       `json:"gasPrice"`
	GasUsed           utils.BN       `json:"gasUsed"`
	CumulativeGasUsed utils.BN       `json:"cumulativeGasUsed"`
	Input             string         `json:"input"`
	Confirmations     int            `json:"confirmations,string"`
}

// MinedBlock holds info from query for mined block by address
type MinedBlock struct {
	BlockNumber int        `json:"blockNumber,string"`
	TimeStamp   utils.Time `json:"timeStamp"`
	BlockReward utils.BN   `json:"blockReward"`
}

// HistoricalBalance holds info from query for historical balance by block number
type HistoricalBalance utils.BN

type HoldingERC20 struct {
	TokenDivisor  string         `json:"TokenDivisor"`
	TokenName     string         `json:"TokenName"`
	TokenQuantity string         `json:"TokenQuantity"`
	TokenAddress  common.Address `json:"TokenAddress"`
	TokenSymbol   string         `json:"TokenSymbol"`
}

type HoldingERC721 struct {
	TokenAddress  common.Address `json:"TokenAddress"`
	TokenName     string         `json:"TokenName"`
	TokenSymbol   string         `json:"TokenSymbol"`
	TokenQuantity string         `json:"TokenQuantity"`
}
type ERC721Inventory struct {
	TokenAddress common.Address `json:"TokenAddress"`
	TokenId      string         `json:"TokenId"`
}
