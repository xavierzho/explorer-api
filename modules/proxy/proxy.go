package proxy

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xavierzho/explorer-api/modules"
	"github.com/xavierzho/explorer-api/utils"
)

type Service modules.Service

func (*Service) Name() string { return "proxy" }

func (s *Service) GetBlockNumber() (blockNumber utils.BN, err error) {
	err = s.Client.Call(s, "eth_blockNumber", nil, &blockNumber)
	return
}

func (s *Service) GetBlockByNumber(blockNumber *utils.BN) (block Block, err error) {
	err = s.Client.Call(s, "eth_getBlockByNumber", utils.M{"tag": blockNumber.Hex(), "boolean": "true"}, &block)
	return
}

func (s *Service) GetUncleByBlockNumberAndIndex(tag, index *utils.BN) (uncle Uncle, err error) {
	err = s.Client.Call(s, "eth_getUncleByBlockNumberAndIndex", utils.M{"tag": tag.Hex(), "index": index.Hex()}, &uncle)
	return
}

func (s *Service) GetTxCountByBlockNumber(tag *utils.BN) (txCount utils.BN, err error) {
	tagStr := tag.Hex()
	if tag.Int().Int64() == 0 || tag == nil {
		tagStr = "latest"
	}
	err = s.Client.Call(s, "eth_getBlockTransactionCountByNumber", utils.M{"tag": tagStr}, &txCount)
	return
}

func (s *Service) GetTxByHash(hash common.Hash) (tx Tx, err error) {
	err = s.Client.Call(s, "eth_getTransactionByHash", utils.M{"txhash": hash.Hex()}, &tx)
	return
}

func (s *Service) GetTxByBlockNumberAndIndex(tag, index *utils.BN) (tx Tx, err error) {
	err = s.Client.Call(s, "eth_getTransactionByBlockNumberAndIndex", utils.M{"tag": tag.Hex(), "index": index.Hex()}, &tx)
	return
}

func (s *Service) GetTxCountByAddress(address common.Address) (txCount utils.BN, err error) {
	err = s.Client.Call(s, "eth_getTransactionCount", utils.M{"address": address.Hex()}, &txCount)
	return
}

func (s *Service) GetTxReceiptByHash(hash common.Hash) (txReceipt TxReceipt, err error) {
	err = s.Client.Call(s, "eth_getTransactionReceipt", utils.M{"txhash": hash.Hex()}, &txReceipt)
	return
}

func (s *Service) EthCall(to common.Address, data []byte, tag *utils.BN) (result string, err error) {
	tagStr := tag.Hex()
	if tag.Int().Int64() == 0 || tag == nil {
		tagStr = "latest"
	}
	err = s.Client.Call(s, "eth_call", utils.M{"to": to.Hex(), "data": fmt.Sprintf("0x%x", data), "tag": tagStr}, &result)
	return
}

func (s *Service) GetCode(address common.Address, tag *utils.BN) (code string, err error) {
	tagStr := tag.Hex()
	if tag.Int().Int64() == 0 || tag == nil {
		tagStr = "latest"
	}
	err = s.Client.Call(s, "eth_getCode", utils.M{"address": address.Hex(), "tag": tagStr}, &code)
	return
}

func (s *Service) GetStorageAt(address common.Address, position, tag *utils.BN) (storage string, err error) {
	tagStr := tag.Hex()
	if tag.Int().Int64() == 0 || tag == nil {
		tagStr = "latest"
	}
	err = s.Client.Call(s, "eth_getStorageAt", utils.M{
		"address":  address.Hex(),
		"position": position.Hex(),
		"tag":      tagStr},
		&storage)
	return
}

func (s *Service) GetGasPrice() (gasPrice utils.BN, err error) {
	err = s.Client.Call(s, "eth_gasPrice", nil, &gasPrice)
	return
}

func (s *Service) GetEstimatedGas(to common.Address, value string, data []byte, g, gasPrice *utils.BN) (gas string, err error) {
	err = s.Client.Call(s, "eth_estimateGas", utils.M{
		"to":       to.Hex(),
		"value":    value,
		"data":     fmt.Sprintf("0x%x", data),
		"gas":      g.Hex(),
		"gasPrice": gasPrice.Hex(),
	}, &gas)
	return
}
