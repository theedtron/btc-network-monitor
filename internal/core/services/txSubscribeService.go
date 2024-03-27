package services

import (
	"btc-network-monitor/internal/adapter/api/requests"
	mysql_repo "btc-network-monitor/internal/adapter/repositories/mysql"
	"btc-network-monitor/internal/core/domain"
	"btc-network-monitor/internal/logger"
	"btc-network-monitor/internal/ports"
	"errors"
)

type TxSubscribeService struct {
	repo ports.Repository
}

func NewTxSubscribeService() *TxSubscribeService {
	return &TxSubscribeService{
		repo: mysql_repo.NewTxSubscribeRepository(),
	}
}

func (s *TxSubscribeService) Save(input domain.TxSubscribe) (interface{}, error) {
	txSub := domain.TxSubscribe{
		TxID:           input.TxID,
		TargetConfirms: input.TargetConfirms,
		UserID:         input.UserID,
	}
	return s.repo.Create(&txSub)
}

func (s *TxSubscribeService) Update(id string, input requests.UpdateTxSubscribeRequest) (interface{}, error) {
	return s.repo.Update(id, input)
}

func (s *TxSubscribeService) GetAll(param map[string]interface{}) (interface{}, error) {
	return s.repo.GetAll(param)
}

func (s *TxSubscribeService) GetFalseStatus() (interface{}, error) {
	return s.repo.GetFalseStatus()
}

func (s *TxSubscribeService) Find(id string) (interface{}, error) {
	return s.repo.Find(id)
}

func (s *TxSubscribeService) FindByTxId(param string) (*domain.TxSubscribe, error) {
	var tx []domain.TxSubscribe
	data, err := s.GetAll(map[string]interface{}{"txid": param})

	if err != nil {
		logger.Error("transaction ID already exists")
		return nil, errors.New("transaction ID already exists")
	}

	tx, exists := data.([]domain.TxSubscribe)
	if !exists {
		logger.Error("error decoding transaction")
		return nil, errors.New("error decoding transaction")
	}

	if len(tx) == 0 {
		logger.Error("transaction not found")
		return nil, errors.New("transaction not found")
	}

	return &tx[0], nil
}
