package services

import (
	"btc-network-monitor/internal/adapter/api/rpc"
	mysql_repo "btc-network-monitor/internal/adapter/repositories/mysql"
	"btc-network-monitor/internal/logger"
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

func (m *MonitorService) GetBlockByHash(str string) (interface{}, error) {
	blockHash, err := chainhash.NewHashFromStr(str)
	if err != nil {
		logger.Error("Invalid block hash: " + err.Error())
		return nil, err
	}

	block, err := m.client.GetBlockVerbose(blockHash)
	if err != nil {
		logger.Error("Error getting block : " + err.Error())

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

func (m *MonitorService) GetTransactionByTransactionID(id string) (interface{}, error) {
	txHash, err := chainhash.NewHashFromStr(id)
	if err != nil {
		logger.Error("Invalid transaction ID:" + err.Error())
		return nil, err
	}

	tx, err := m.client.GetRawTransactionVerbose(txHash)
	if err != nil {
		logger.Error("Error getting transaction: " + err.Error())
		return nil, err
	}

	return map[string]interface{}{
		"transaction": tx,
	}, nil

}

func (m *MonitorService) GetLatestTransactions() (interface{}, error) {
	transactions, err := m.client.GetRawMempool()
	if err != nil {
		logger.Error("Error getting latest transactions: " + err.Error())
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

	var blockHashes []string
	for i := 0; i < 10; i++ { // Assuming we want the latest 10 blocks
		blockHash, err := m.client.GetBlockHash(int64(chainInfo.Blocks - int32(i)))
		if err != nil {
			return nil, err
		}
		blockHashes = append(blockHashes, blockHash.String())
	}

	return map[string]interface{}{
		"blocks": blockHashes,
	}, nil

}
