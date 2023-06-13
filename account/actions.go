package account

import (
	"github.com/Jonescy/explorer-api"
	"github.com/Jonescy/explorer-api/utils"
)

type Action interface {
	explorer.Endpoint
	// AccountBalance gets ether balance for a single address
	EtherBalance(address string) (balance *utils.BigInt, err error)
	// MultiAccountBalance gets ether balance for multiple addresses in a single call
	MultiAccountBalance(addresses ...string) (balances []Balance, err error)
	// NormalTxByAddress gets a list of "normal" transactions by address
	NormalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []NormalTx, err error)
	// InternalTxByAddress gets a list of "internal" transactions by address
	InternalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []InternalTx, err error)
	// InternalTxsByHash gets a list of "internal" transactions by transaction hash
	InternalTxsByHash(txHash string) (txs []InternalTx, err error)
	// InternalTxsByBlockRange gets a list of "internal" transactions by block range
	InternalTxsByBlockRange(startBlock int, endBlock int, page int, offset int, desc bool) (txs []InternalTx, err error)
	// ERC20TransferEventsByAddress gets a list of ERC20 token transfer events by address
	ERC20TransferEventsByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC20Transfer, err error)
	// ERC721TransferEventsByAddress gets a list of ERC721 token transfer events by address
	ERC721TransferEventsByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC721Transfer, err error)
	// ERC1155TransferEventsByAddress gets a list of ERC1155 token transfer events by address
	ERC1155TransferEventsByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC1155Transfer, err error)
	// MinedBlocksByAddress gets a list of mined blocks by address
	MinedBlocksByAddress(address string, page int, offset int) (blocks []MinedBlock, err error)
	// HistoricalByBlockNo gets historical block information by block number
	HistoricalByBlockNo(address string, blockNo int) (balance HistoricalBalance, err error)

	ERC20Balance(address string, contractAddress string) (balance *utils.BigInt, err error)
	HistoricalERC20Balance(address string, contractAddress string, blockNo int) (balance HistoricalBalance, err error)
}

// Balance account and its balance in pair
type Balance struct {
	Account string        `json:"account"`
	Balance *utils.BigInt `json:"balance"`
}

// NormalTx holds info from normal tx query
type NormalTx struct {
	BlockNumber       int           `json:"blockNumber,string"`
	TimeStamp         utils.Time    `json:"timeStamp"`
	Hash              string        `json:"hash"`
	Nonce             int           `json:"nonce,string"`
	BlockHash         string        `json:"blockHash"`
	TransactionIndex  int           `json:"transactionIndex,string"`
	From              string        `json:"from"`
	To                string        `json:"to"`
	Value             *utils.BigInt `json:"value"`
	Gas               int           `json:"gas,string"`
	GasPrice          *utils.BigInt `json:"gasPrice"`
	IsError           int           `json:"isError,string"`
	TxReceiptStatus   string        `json:"txreceipt_status"`
	Input             string        `json:"input"`
	ContractAddress   string        `json:"contractAddress"`
	CumulativeGasUsed int           `json:"cumulativeGasUsed,string"`
	GasUsed           int           `json:"gasUsed,string"`
	Confirmations     int           `json:"confirmations,string"`
}

// InternalTx holds info from internal tx query
type InternalTx struct {
	BlockNumber     int           `json:"blockNumber,string"`
	TimeStamp       utils.Time    `json:"timeStamp"`
	Hash            string        `json:"hash"`
	From            string        `json:"from"`
	To              string        `json:"to"`
	Value           *utils.BigInt `json:"value"`
	ContractAddress string        `json:"contractAddress"`
	Input           string        `json:"input"`
	Type            string        `json:"type"`
	Gas             int           `json:"gas,string"`
	GasUsed         int           `json:"gasUsed,string"`
	TraceID         string        `json:"traceId"`
	IsError         int           `json:"isError,string"`
	ErrCode         string        `json:"errCode"`
}

// ERC20Transfer holds info from ERC20 token transfer event query
type ERC20Transfer struct {
	BlockNumber       int           `json:"blockNumber,string"`
	TimeStamp         utils.Time    `json:"timeStamp"`
	Hash              string        `json:"hash"`
	Nonce             int           `json:"nonce,string"`
	BlockHash         string        `json:"blockHash"`
	From              string        `json:"from"`
	ContractAddress   string        `json:"contractAddress"`
	To                string        `json:"to"`
	Value             *utils.BigInt `json:"value"`
	TokenName         string        `json:"tokenName"`
	TokenSymbol       string        `json:"tokenSymbol"`
	TokenDecimal      uint8         `json:"tokenDecimal,string"`
	TransactionIndex  int           `json:"transactionIndex,string"`
	Gas               int           `json:"gas,string"`
	GasPrice          *utils.BigInt `json:"gasPrice"`
	GasUsed           int           `json:"gasUsed,string"`
	CumulativeGasUsed int           `json:"cumulativeGasUsed,string"`
	Input             string        `json:"input"`
	Confirmations     int           `json:"confirmations,string"`
}

// ERC721Transfer holds info from ERC721 token transfer event query
type ERC721Transfer struct {
	BlockNumber       int           `json:"blockNumber,string"`
	TimeStamp         utils.Time    `json:"timeStamp"`
	Hash              string        `json:"hash"`
	Nonce             int           `json:"nonce,string"`
	BlockHash         string        `json:"blockHash"`
	From              string        `json:"from"`
	ContractAddress   string        `json:"contractAddress"`
	To                string        `json:"to"`
	TokenID           *utils.BigInt `json:"tokenID"`
	TokenName         string        `json:"tokenName"`
	TokenSymbol       string        `json:"tokenSymbol"`
	TokenDecimal      uint8         `json:"tokenDecimal,string"`
	TransactionIndex  int           `json:"transactionIndex,string"`
	Gas               int           `json:"gas,string"`
	GasPrice          *utils.BigInt `json:"gasPrice"`
	GasUsed           int           `json:"gasUsed,string"`
	CumulativeGasUsed int           `json:"cumulativeGasUsed,string"`
	Input             string        `json:"input"`
	Confirmations     int           `json:"confirmations,string"`
}

// ERC1155Transfer holds info from ERC1155 token transfer event query
type ERC1155Transfer struct {
	BlockNumber       int           `json:"blockNumber,string"`
	TimeStamp         utils.Time    `json:"timeStamp"`
	Hash              string        `json:"hash"`
	Nonce             int           `json:"nonce,string"`
	BlockHash         string        `json:"blockHash"`
	From              string        `json:"from"`
	ContractAddress   string        `json:"contractAddress"`
	To                string        `json:"to"`
	TokenID           *utils.BigInt `json:"tokenID"`
	TokenName         string        `json:"tokenName"`
	TokenSymbol       string        `json:"tokenSymbol"`
	TokenDecimal      uint8         `json:"tokenDecimal,string"`
	TokenValue        uint8         `json:"tokenValue,string"`
	TransactionIndex  int           `json:"transactionIndex,string"`
	Gas               int           `json:"gas,string"`
	GasPrice          *utils.BigInt `json:"gasPrice"`
	GasUsed           int           `json:"gasUsed,string"`
	CumulativeGasUsed int           `json:"cumulativeGasUsed,string"`
	Input             string        `json:"input"`
	Confirmations     int           `json:"confirmations,string"`
}

// MinedBlock holds info from query for mined block by address
type MinedBlock struct {
	BlockNumber int           `json:"blockNumber,string"`
	TimeStamp   utils.Time    `json:"timeStamp"`
	BlockReward *utils.BigInt `json:"blockReward"`
}

// HistoricalBalance holds info from query for historical balance by block number
type HistoricalBalance *utils.BigInt
