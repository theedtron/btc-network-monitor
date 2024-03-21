package services

import (
	"btc-network-monitor/internal/adapter/api/requests"
	mysql_repo "btc-network-monitor/internal/adapter/repositories/mysql"
	"btc-network-monitor/internal/core/domain"
	"btc-network-monitor/internal/ports"
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
		TxID: input.TxID,
		TargetConfirms: input.TargetConfirms,
		UserID: input.UserID,
	}
	return s.repo.Create(&txSub)
}

func (s *TxSubscribeService) Update(id string, input requests.UpdateTxSubscribeRequest) (interface{}, error) {
	return s.repo.Update(id, input)
}

func (s *TxSubscribeService) GetAll(param map[string]interface{}) (interface{}, error) {
	return s.repo.GetAll(param)
}

func (s *TxSubscribeService) Find(id string) (interface{}, error) {
	return s.repo.Find(id)
}