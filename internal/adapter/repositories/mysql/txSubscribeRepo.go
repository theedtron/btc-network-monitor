package mysql_repo

import (
	"btc-network-monitor/internal/core/domain"
	"btc-network-monitor/internal/ports"
	"gorm.io/gorm/clause"
)

type TxSubscribeRepository struct {}

func NewTxSubscribeRepository() ports.Repository {
	return &TxSubscribeRepository{}
}

func (repo *TxSubscribeRepository) GetAll(param map[string]interface{}) (interface{}, error) {
	model := repo.ArrayModel()

	q := db.Preload(clause.Associations)

	if param["txid"] != nil {
		q.Where("tx_id = ?", param["txid"])
	}

	q.Find(&model)
	return model, q.Error
}

func (repo *TxSubscribeRepository) Find(id string) (interface{}, error) {
	model := repo.Model()
	q := db.Preload(clause.Associations).Where("id = ?", id).First(&model)

	return model, q.Error
}

func (repo *TxSubscribeRepository) Create(data interface{}) (interface{}, error) {
	model := repo.Model()
	q := db.Model(model).Create(data)
	return data, q.Error

}

func (repo *TxSubscribeRepository) Update(id string, data interface{}) (interface{}, error) {
	model := repo.Model()
	q := db.Model(&model).Where("id = ?", id).Updates(data)
	if q.Error != nil {
		return nil, q.Error
	}
	return repo.Find(id)
}

func (repo *TxSubscribeRepository) Delete(id string) (interface{}, error) {
	model := repo.Model()
	q := db.Model(&model).Where("id = ?", id).Delete(model)
	return model, q.Error
}

func (repo *TxSubscribeRepository) Model() domain.TxSubscribe {
	return domain.TxSubscribe{}
}

func (repo *TxSubscribeRepository) ArrayModel() []domain.TxSubscribe {
	return []domain.TxSubscribe{}
}
