package proxy

import (
	"github.com/Jonescy/explorer-api"
	"github.com/Jonescy/explorer-api/utils"
	"github.com/ethereum/go-ethereum/common"
)

type Action interface {
	explorer.Module
	GetBlockNumber() (blockNumber utils.BN, err error)
	GetBlockByNumber(blockNumber *utils.BN) (block Block, err error)
	GetUncleByBlockNumberAndIndex(tag, index *utils.BN) (uncle Uncle, err error)
	GetTxCountByBlockNumber(tag *utils.BN) (txCount utils.BN, err error)
	GetTxByHash(hash common.Hash) (tx Tx, err error)
	GetTxByBlockNumberAndIndex(tag, index *utils.BN) (tx Tx, err error)
	GetTxCountByAddress(address common.Address) (txCount utils.BN, err error)
	GetTxReceiptByHash(hash common.Hash) (txReceipt TxReceipt, err error)
	EthCall(to common.Address, data []byte, tag *utils.BN) (result string, err error)
	GetCode(address common.Address, tag *utils.BN) (code string, err error)
	GetStorageAt(address common.Address, position, tag *utils.BN) (storage string, err error)
	GetGasPrice() (gasPrice utils.BN, err error)
	GetEstimatedGas(to common.Address, value string, data []byte, g, gasPrice *utils.BN) (gas string, err error)
}

type Block struct {
	LogsBloom       string      `json:"logsBloom"`
	TotalDifficulty utils.BN    `json:"totalDifficulty"`
	ReceiptsRoot    common.Hash `json:"receiptsRoot"`
	ExtraData       string      `json:"extraData"`
	BaseFeePerGas   utils.BN    `json:"baseFeePerGas"`
	Transactions    []struct {
		BlockHash        common.Hash     `json:"blockHash"`
		TransactionIndex string          `json:"transactionIndex"`
		Type             string          `json:"type"`
		Nonce            string          `json:"nonce"`
		Input            string          `json:"input"`
		R                string          `json:"r"`
		S                string          `json:"s"`
		V                string          `json:"v"`
		BlockNumber      string          `json:"blockNumber"`
		Gas              utils.BN        `json:"gas"`
		From             common.Address  `json:"from"`
		To               *common.Address `json:"to"`
		Value            utils.BN        `json:"value"`
		Hash             common.Hash     `json:"hash"`
		GasPrice         utils.BN        `json:"gasPrice"`
	} `json:"transactions"`
	Nonce            utils.BN       `json:"nonce"`
	Miner            common.Address `json:"miner"`
	Difficulty       utils.BN       `json:"difficulty"`
	GasLimit         utils.BN       `json:"gasLimit"`
	Number           utils.BN       `json:"number"`
	GasUsed          utils.BN       `json:"gasUsed"`
	Uncles           []common.Hash  `json:"uncles"`
	Sha3Uncles       common.Hash    `json:"sha3Uncles"`
	Size             utils.BN       `json:"size"`
	TransactionsRoot common.Hash    `json:"transactionsRoot"`
	StateRoot        common.Hash    `json:"stateRoot"`
	MixHash          common.Hash    `json:"mixHash"`
	ParentHash       common.Hash    `json:"parentHash"`
	Hash             common.Hash    `json:"hash"`
	Timestamp        utils.BN       `json:"timestamp"`
}

type Uncle struct {
	LogsBloom        string         `json:"logsBloom"`
	ReceiptsRoot     common.Hash    `json:"receiptsRoot"`
	ExtraData        string         `json:"extraData"`
	BaseFeePerGas    utils.BN       `json:"baseFeePerGas"`
	Nonce            utils.BN       `json:"nonce"`
	Miner            common.Address `json:"miner"`
	Difficulty       utils.BN       `json:"difficulty"`
	GasLimit         utils.BN       `json:"gasLimit"`
	Number           utils.BN       `json:"number"`
	GasUsed          utils.BN       `json:"gasUsed"`
	Uncles           []common.Hash  `json:"uncles"`
	Sha3Uncles       common.Hash    `json:"sha3Uncles"`
	Size             utils.BN       `json:"size"`
	TransactionsRoot common.Hash    `json:"transactionsRoot"`
	StateRoot        common.Hash    `json:"stateRoot"`
	MixHash          common.Hash    `json:"mixHash"`
	ParentHash       common.Hash    `json:"parentHash"`
	Hash             common.Hash    `json:"hash"`
	Timestamp        utils.BN       `json:"timestamp"`
}

type Tx struct {
	BlockHash            common.Hash     `json:"blockHash"`
	AccessList           []interface{}   `json:"accessList"`
	TransactionIndex     utils.BN        `json:"transactionIndex"`
	Type                 string          `json:"type"`
	Nonce                utils.BN        `json:"nonce"`
	Input                string          `json:"input"`
	R                    string          `json:"r"`
	S                    string          `json:"s"`
	ChainID              utils.BN        `json:"chainId"`
	V                    string          `json:"v"`
	BlockNumber          string          `json:"blockNumber"`
	Gas                  utils.BN        `json:"gas"`
	MaxPriorityFeePerGas utils.BN        `json:"maxPriorityFeePerGas"`
	From                 common.Address  `json:"from"`
	To                   *common.Address `json:"to"`
	MaxFeePerGas         utils.BN        `json:"maxFeePerGas"`
	Value                string          `json:"value"`
	Hash                 common.Hash     `json:"hash"`
	GasPrice             utils.BN        `json:"gasPrice"`
}

type TxReceipt struct {
	BlockHash         common.Hash     `json:"blockHash"`
	LogsBloom         string          `json:"logsBloom"`
	ContractAddress   string          `json:"contractAddress"`
	TransactionIndex  string          `json:"transactionIndex"`
	Type              string          `json:"type"`
	TransactionHash   common.Hash     `json:"transactionHash"`
	GasUsed           utils.BN        `json:"gasUsed"`
	BlockNumber       utils.BN        `json:"blockNumber"`
	CumulativeGasUsed utils.BN        `json:"cumulativeGasUsed"`
	From              common.Address  `json:"from"`
	To                *common.Address `json:"to"`
	EffectiveGasPrice utils.BN        `json:"effectiveGasPrice"`
	Logs              []struct {
		BlockHash        common.Hash    `json:"blockHash"`
		Address          common.Address `json:"address"`
		LogIndex         string         `json:"logIndex"`
		Data             string         `json:"data"`
		Removed          bool           `json:"removed"`
		Topics           []common.Hash  `json:"topics"`
		BlockNumber      utils.BN       `json:"blockNumber"`
		TransactionIndex string         `json:"transactionIndex"`
		TransactionHash  common.Hash    `json:"transactionHash"`
	} `json:"logs"`
	Status string `json:"status"`
}
