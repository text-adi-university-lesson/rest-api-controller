package service

import (
	"rest-api-controller/internal/module/sensors/model"
	"rest-api-controller/internal/module/sensors/repository"

	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Repository repository.Repository
}

type Service struct {
	repo repository.Repository
}

func NewService(p Params) *Service {
	return &Service{
		repo: p.Repository,
	}
}

func (receiver Service) DeleteTemperatureItem(id string) error {
	return receiver.repo.DeleteTemperatureByID(id)
}
func (receiver Service) CreateTemperatureItem(item *model.CreateTemperatureSensorsDTO) (*model.TemperatureSensor, error) {
	data := item.ToModel()
	return receiver.repo.AddTemperatureItem(&data)
}
func (receiver Service) UpdateTemperatureItem(id string, item *model.UpdateTemperatureSensorsDTO) (*model.TemperatureSensor, error) {
	data := item.ToModel(id)
	return receiver.repo.UpdateTemperatureItem(&data)
}
func (receiver Service) GetTemperatureItemByID(id string) (*model.TemperatureSensor, error) {
	return receiver.repo.GetTemperatureItemByID(id)
}
func (receiver Service) GetAllTemperature() ([]model.TemperatureSensor, error) {
	return receiver.repo.GetAllTemperatureItems()
}

func (receiver Service) DeleteHumidityItem(id string) error {
	return receiver.repo.DeleteHumidityByID(id)
}
func (receiver Service) CreateHumidityItem(item *model.CreateHumiditySensorsDTO) (*model.HumiditySensor, error) {
	data := item.ToModel()
	return receiver.repo.AddHumidityItem(&data)
}
func (receiver Service) UpdateHumidityItem(id string, item *model.UpdateHumiditySensorsDTO) (*model.HumiditySensor, error) {
	data := item.ToModel(id)
	return receiver.repo.UpdateHumidityItem(&data)
}
func (receiver Service) GetHumidityItemByID(id string) (*model.HumiditySensor, error) {
	return receiver.repo.GetHumidityItemByID(id)
}
func (receiver Service) GetAllHumidity() ([]model.HumiditySensor, error) {
	return receiver.repo.GetAllHumidityItems()
}
