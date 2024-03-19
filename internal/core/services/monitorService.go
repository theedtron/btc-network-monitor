package services

import (
	"btc-network-monitor/internal/adapter/api/rpc"
	mysql_repo "btc-network-monitor/internal/adapter/repositories/mysql"
	"btc-network-monitor/internal/ports"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
)

type MonitorService struct {
	port   ports.Repository
	client *rpcclient.Client
}

func NewMonitorService() *MonitorService {
	return &MonitorService{
		port:   mysql_repo.NewMonitorRepository(),
		client: rpc.Config,
	}
}

func (m *MonitorService) GetBlockChainInfo() (interface{}, error) {
	info, err := m.client.GetBlockChainInfo()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"current_block_height": info.Blocks,
		"difficulty":           info.Difficulty,
		"network_hash_rate":    info.BestBlockHash,
	}, nil

}

func (m *MonitorService) GetBlockByHash(hash *chainhash.Hash) (interface{}, error) {
	block, err := m.client.GetBlockVerbose(hash)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"block": block,
	}, nil

}

func (m *MonitorService) GetBlockByHeight(height int64) (interface{}, error) {
	blockHash, err := m.client.GetBlockHash(height)
	if err != nil {
		return nil, err
	}

	block, err := m.client.GetBlockVerbose(blockHash)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"block": block,
	}, nil
}

func (m *MonitorService) GetTransactionByTransactionID(hash *chainhash.Hash) (interface{}, error) {
	tx, err := m.client.GetRawTransactionVerbose(hash)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"transaction": tx,
	}, nil

}

func (m *MonitorService) GetLatestTransactions() (interface{}, error) {
	transactions, err := m.client.GetRawMempool()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"transactions": transactions,
	}, nil
}

func (m *MonitorService) GetLatestBlocks() (interface{}, error) {
	chainInfo, err := m.client.GetBlockChainInfo()
	if err != nil {
		return nil, err
	}

	var blocks []interface{}
	for i := 0; i < 10; i++ {
		blockHash, err := m.client.GetBlockHash(int64(chainInfo.Blocks - int32(i)))
		if err != nil {
			return nil, err
		}

		block, err := m.client.GetBlockVerbose(blockHash)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)

	}

	return map[string]interface{}{
		"blocks": blocks,
	}, nil

}

func (m *MonitorService) GetAddressInfo(address string) (interface{}, error) {
	info, err := m.client.GetAddressInfo(address)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"info": info,
	}, nil

}
