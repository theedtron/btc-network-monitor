package mysql_repo

import (
	"btc-network-monitor/internal/core/domain"
	"btc-network-monitor/internal/ports"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
}

func (repo *UserRepository) GetFalseStatus() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository() ports.Repository {
	return &UserRepository{}
}

func (repo *UserRepository) GetAll(param map[string]interface{}) (interface{}, error) {
	model := repo.ArrayModel()

	q := db.Preload(clause.Associations)

	if param["name"] != nil {
		q.Where("name = ?", param["name"])
	}

	if param["email"] != nil {
		q.Where("email = ?", param["email"])
	}

	q.Find(&model)
	return model, q.Error
}

func (repo *UserRepository) Find(id string) (interface{}, error) {
	model := repo.Model()
	q := db.Preload(clause.Associations).Where("id = ?", id).First(&model)

	return model, q.Error
}

func (repo *UserRepository) Create(data interface{}) (interface{}, error) {
	model := repo.Model()
	q := db.Model(model).Create(data)
	return data, q.Error

}

func (repo *UserRepository) Update(id string, data interface{}) (interface{}, error) {
	model := repo.Model()
	q := db.Model(&model).Where("id = ?", id).Updates(data)
	if q.Error != nil {
		return nil, q.Error
	}
	return repo.Find(id)
}

func (repo *UserRepository) Delete(id string) (interface{}, error) {
	model := repo.Model()
	q := db.Model(&model).Where("id = ?", id).Delete(model)
	return model, q.Error
}

func (repo *UserRepository) Model() domain.User {
	return domain.User{}
}

func (repo *UserRepository) ArrayModel() []domain.User {
	return []domain.User{}
}
