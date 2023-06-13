package proxy

import (
	"fmt"
	"github.com/Jonescy/explorer-api"
)

type Service explorer.Service

func (*Service) Name() string { return "proxy" }

func (s *Service) GetBlockNumber() (blockNumber string, err error) {
	err = s.Client.Call(s.Name(), "eth_blockNumber", nil, &blockNumber)
	return
}

func (s *Service) GetBlockByNumber(blockNumber int) (block Block, err error) {

	err = s.Client.Call(s.Name(), "eth_getBlockByNumber", map[string]string{"tag": fmt.Sprintf("0x%x", blockNumber)}, &block)
	return
}

func (s *Service) GetUncleByBlockNumberAndIndex(tag, index int) (uncle Uncle, err error) {
	err = s.Client.Call(s.Name(), "eth_getUncleByBlockNumberAndIndex", map[string]string{"tag": fmt.Sprintf("0x%x", tag), "index": fmt.Sprintf("0x%x", index)}, &uncle)
	return
}

func (s *Service) GetTxCountByBlockNumber(tag int) (txCount string, err error) {
	tagStr := fmt.Sprintf("0x%x", tag)
	if tag == 0 {
		tagStr = "latest"
	}
	err = s.Client.Call(s.Name(), "eth_getBlockTransactionCountByNumber", map[string]string{"tag": tagStr}, &txCount)
	return
}

func (s *Service) GetTxByBlockHash(hash string) (tx Tx, err error) {
	err = s.Client.Call(s.Name(), "eth_getTransactionByBlockHashAndIndex", map[string]string{"txhash": hash}, &tx)
	return
}

func (s *Service) GetTxByBlockNumberAndIndex(tag, index int) (tx Tx, err error) {
	err = s.Client.Call(s.Name(), "eth_getTransactionByBlockNumberAndIndex", map[string]string{"tag": fmt.Sprintf("0x%x", tag), "index": fmt.Sprintf("0x%x", index)}, &tx)
	return
}

func (s *Service) GetTxCountByAddress(address string) (txCount string, err error) {
	err = s.Client.Call(s.Name(), "eth_getTransactionCount", map[string]string{"address": address}, &txCount)
	return
}

func (s *Service) GetTxReceiptByHash(hash string) (txReceipt TxReceipt, err error) {
	err = s.Client.Call(s.Name(), "eth_getTransactionReceipt", map[string]string{"txhash": hash}, &txReceipt)
	return
}

func (s *Service) EthCall(to string, data []byte, tag int) (result string, err error) {
	tagStr := fmt.Sprintf("0x%x", tag)
	if tag == 0 {
		tagStr = "latest"
	}
	err = s.Client.Call(s.Name(), "eth_call", map[string]string{"to": to, "data": fmt.Sprintf("0x%x", data), "tag": tagStr}, &result)
	return
}

func (s *Service) GetCode(address string, tag int) (code string, err error) {
	tagStr := fmt.Sprintf("0x%x", tag)
	if tag == 0 {
		tagStr = "latest"
	}
	err = s.Client.Call(s.Name(), "eth_getCode", map[string]string{"address": address, "tag": tagStr}, &code)
	return
}

func (s *Service) GetStorageAt(address string, position int, tag int) (storage string, err error) {
	tagStr := fmt.Sprintf("0x%x", tag)
	if tag == 0 {
		tagStr = "latest"
	}
	err = s.Client.Call(s.Name(), "eth_getStorageAt", map[string]string{"address": address, "position": fmt.Sprintf("0x%x", position), "tag": tagStr}, &storage)
	return
}

func (s *Service) GetGasPrice() (gasPrice string, err error) {
	err = s.Client.Call(s.Name(), "eth_gasPrice", nil, &gasPrice)
	return
}

func (s *Service) GetEstimatedGas(from, to string, value string, data []byte, tag int) (gas string, err error) {
	tagStr := fmt.Sprintf("0x%x", tag)
	if tag == 0 {
		tagStr = "latest"
	}
	err = s.Client.Call(s.Name(), "eth_estimateGas", map[string]string{"from": from, "to": to, "value": value, "data": fmt.Sprintf("0x%x", data), "tag": tagStr}, &gas)
	return
}
