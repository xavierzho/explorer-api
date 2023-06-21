package proxy

import (
	"fmt"
	"github.com/Jonescy/explorer-api/modules"
)

type Service modules.Service

func (*Service) Name() string { return "proxy" }

func (s *Service) GetBlockNumber() (blockNumber string, err error) {
	err = s.Client.Call(s, "eth_blockNumber", nil, &blockNumber)
	return
}

func (s *Service) GetBlockByNumber(blockNumber int) (block Block, err error) {
	err = s.Client.Call(s, "eth_getBlockByNumber", map[string]string{"tag": fmt.Sprintf("0x%x", blockNumber), "boolean": "true"}, &block)
	return
}

func (s *Service) GetUncleByBlockNumberAndIndex(tag, index int) (uncle Uncle, err error) {
	err = s.Client.Call(s, "eth_getUncleByBlockNumberAndIndex", map[string]string{"tag": fmt.Sprintf("0x%x", tag), "index": fmt.Sprintf("0x%x", index)}, &uncle)
	return
}

func (s *Service) GetTxCountByBlockNumber(tag int) (txCount string, err error) {
	tagStr := fmt.Sprintf("0x%x", tag)
	if tag == 0 {
		tagStr = "latest"
	}
	err = s.Client.Call(s, "eth_getBlockTransactionCountByNumber", map[string]string{"tag": tagStr}, &txCount)
	return
}

func (s *Service) GetTxByHash(hash string) (tx Tx, err error) {
	err = s.Client.Call(s, "eth_getTransactionByHash", map[string]string{"txhash": hash}, &tx)
	return
}

func (s *Service) GetTxByBlockNumberAndIndex(tag, index int) (tx Tx, err error) {
	err = s.Client.Call(s, "eth_getTransactionByBlockNumberAndIndex", map[string]string{"tag": fmt.Sprintf("0x%x", tag), "index": fmt.Sprintf("0x%x", index)}, &tx)
	return
}

func (s *Service) GetTxCountByAddress(address string) (txCount string, err error) {
	err = s.Client.Call(s, "eth_getTransactionCount", map[string]string{"address": address}, &txCount)
	return
}

func (s *Service) GetTxReceiptByHash(hash string) (txReceipt TxReceipt, err error) {
	err = s.Client.Call(s, "eth_getTransactionReceipt", map[string]string{"txhash": hash}, &txReceipt)
	return
}

func (s *Service) EthCall(to string, data []byte, tag int) (result string, err error) {
	tagStr := fmt.Sprintf("0x%x", tag)
	if tag == 0 {
		tagStr = "latest"
	}
	err = s.Client.Call(s, "eth_call", map[string]string{"to": to, "data": fmt.Sprintf("0x%x", data), "tag": tagStr}, &result)
	return
}

func (s *Service) GetCode(address string, tag int) (code string, err error) {
	tagStr := fmt.Sprintf("0x%x", tag)
	if tag == 0 {
		tagStr = "latest"
	}
	err = s.Client.Call(s, "eth_getCode", map[string]string{"address": address, "tag": tagStr}, &code)
	return
}

func (s *Service) GetStorageAt(address string, position int, tag int) (storage string, err error) {
	tagStr := fmt.Sprintf("0x%x", tag)
	if tag == 0 {
		tagStr = "latest"
	}
	err = s.Client.Call(s, "eth_getStorageAt", map[string]string{"address": address, "position": fmt.Sprintf("0x%x", position), "tag": tagStr}, &storage)
	return
}

func (s *Service) GetGasPrice() (gasPrice string, err error) {
	err = s.Client.Call(s, "eth_gasPrice", nil, &gasPrice)
	return
}

func (s *Service) GetEstimatedGas(to string, value string, data []byte, g, gasPrice int) (gas string, err error) {
	err = s.Client.Call(s, "eth_estimateGas", map[string]string{
		"to":       to,
		"value":    value,
		"data":     fmt.Sprintf("0x%x", data),
		"gas":      fmt.Sprintf("0x%x", g),
		"gasPrice": fmt.Sprintf("0x%x", gasPrice),
	}, &gas)
	return
}
