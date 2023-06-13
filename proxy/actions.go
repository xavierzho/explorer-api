package proxy

import "github.com/Jonescy/explorer-api"

type Action interface {
	explorer.Endpoint
	GetBlockNumber() (blockNumber string, err error)
	GetBlockByNumber(tag int) (block Block, err error)
	GetUncleByBlockNumberAndIndex(tag int, index int) (uncle Uncle, err error)
	GetTxCountByBlockNumber(tag int) (txCount string, err error)
	GetTxByBlockHash(hash string) (tx Tx, err error)
	GetTxByBlockNumberAndIndex(tag int, index int) (tx Tx, err error)
	GetTxCountByAddress(address string) (txCount string, err error)
	GetTxReceiptByHash(hash string) (txReceipt TxReceipt, err error)
	EthCall(to string, data []byte, tag int) (result string, err error)
	GetCode(address string, tag int) (code string, err error)
	GetStorageAt(address string, position int, tag int) (data string, err error)
	GetGasPrice() (gasPrice string, err error)
	GetEstimatedGas(from, to string, value string, data []byte, tag int) (gas string, err error)
}

type Block struct {
	LogsBloom        string   `json:"logsBloom"`
	TotalDifficulty  string   `json:"totalDifficulty"`
	ReceiptsRoot     string   `json:"receiptsRoot"`
	ExtraData        string   `json:"extraData"`
	BaseFeePerGas    string   `json:"baseFeePerGas"`
	Transactions     []string `json:"transactions"`
	Nonce            string   `json:"nonce"`
	Miner            string   `json:"miner"`
	Difficulty       string   `json:"difficulty"`
	GasLimit         string   `json:"gasLimit"`
	Number           string   `json:"number"`
	GasUsed          string   `json:"gasUsed"`
	Uncles           []string `json:"uncles"`
	Sha3Uncles       string   `json:"sha3Uncles"`
	Size             string   `json:"size"`
	TransactionsRoot string   `json:"transactionsRoot"`
	StateRoot        string   `json:"stateRoot"`
	MixHash          string   `json:"mixHash"`
	ParentHash       string   `json:"parentHash"`
	Hash             string   `json:"hash"`
	Timestamp        string   `json:"timestamp"`
}

type Uncle struct {
	LogsBloom        string        `json:"logsBloom"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	ExtraData        string        `json:"extraData"`
	BaseFeePerGas    string        `json:"baseFeePerGas"`
	Nonce            string        `json:"nonce"`
	Miner            string        `json:"miner"`
	Difficulty       string        `json:"difficulty"`
	GasLimit         string        `json:"gasLimit"`
	Number           string        `json:"number"`
	GasUsed          string        `json:"gasUsed"`
	Uncles           []interface{} `json:"uncles"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	TransactionsRoot string        `json:"transactionsRoot"`
	StateRoot        string        `json:"stateRoot"`
	MixHash          string        `json:"mixHash"`
	ParentHash       string        `json:"parentHash"`
	Hash             string        `json:"hash"`
	Timestamp        string        `json:"timestamp"`
}

type Tx struct {
	BlockHash            string        `json:"blockHash"`
	AccessList           []interface{} `json:"accessList"`
	TransactionIndex     string        `json:"transactionIndex"`
	Type                 string        `json:"type"`
	Nonce                string        `json:"nonce"`
	Input                string        `json:"input"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
	ChainID              string        `json:"chainId"`
	V                    string        `json:"v"`
	BlockNumber          string        `json:"blockNumber"`
	Gas                  string        `json:"gas"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	From                 string        `json:"from"`
	To                   string        `json:"to"`
	MaxFeePerGas         string        `json:"maxFeePerGas"`
	Value                string        `json:"value"`
	Hash                 string        `json:"hash"`
	GasPrice             string        `json:"gasPrice"`
}

type TxReceipt struct {
	BlockHash         string      `json:"blockHash"`
	LogsBloom         string      `json:"logsBloom"`
	ContractAddress   interface{} `json:"contractAddress"`
	TransactionIndex  string      `json:"transactionIndex"`
	Type              string      `json:"type"`
	TransactionHash   string      `json:"transactionHash"`
	GasUsed           string      `json:"gasUsed"`
	BlockNumber       string      `json:"blockNumber"`
	CumulativeGasUsed string      `json:"cumulativeGasUsed"`
	From              string      `json:"from"`
	To                string      `json:"to"`
	EffectiveGasPrice string      `json:"effectiveGasPrice"`
	Logs              []struct {
		BlockHash        string   `json:"blockHash"`
		Address          string   `json:"address"`
		LogIndex         string   `json:"logIndex"`
		Data             string   `json:"data"`
		Removed          bool     `json:"removed"`
		Topics           []string `json:"topics"`
		BlockNumber      string   `json:"blockNumber"`
		TransactionIndex string   `json:"transactionIndex"`
		TransactionHash  string   `json:"transactionHash"`
	} `json:"logs"`
	Status string `json:"status"`
}
