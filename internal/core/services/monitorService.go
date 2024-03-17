package services

import (
	"btc-network-monitor/internal/adapter/api/rpc"
	mysql_repo "btc-network-monitor/internal/adapter/repositories/mysql"
	"btc-network-monitor/internal/ports"
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
