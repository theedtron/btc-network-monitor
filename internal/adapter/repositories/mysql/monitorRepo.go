package mysql_repo

import "btc-network-monitor/internal/ports"

type MonitorRepository struct {
}

func NewMonitorRepository() ports.Repository {
	return &MonitorRepository{}
}

func (m MonitorRepository) Create(data interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (m MonitorRepository) Find(id string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (m MonitorRepository) GetAll(param map[string]interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (m MonitorRepository) Update(id string, data interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (m MonitorRepository) Delete(id string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}
