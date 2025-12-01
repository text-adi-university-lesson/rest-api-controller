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
func (receiver Service) CreateTemperatureItem(item *model.TemperatureSensor) (*model.TemperatureSensor, error) {
	return receiver.repo.AddTemperatureItem(item)
}
func (receiver Service) UpdateTemperatureItem(item *model.TemperatureSensor) (*model.TemperatureSensor, error) {
	return receiver.repo.UpdateTemperatureItem(item)
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
func (receiver Service) CreateHumidityItem(item *model.HumiditySensor) (*model.HumiditySensor, error) {
	return receiver.repo.AddHumidityItem(item)
}
func (receiver Service) UpdateHumidityItem(item *model.HumiditySensor) (*model.HumiditySensor, error) {
	return receiver.repo.UpdateHumidityItem(item)
}
func (receiver Service) GetHumidityItemByID(id string) (*model.HumiditySensor, error) {
	return receiver.repo.GetHumidityItemByID(id)
}
func (receiver Service) GetAllHumidity() ([]model.HumiditySensor, error) {
	return receiver.repo.GetAllHumidityItems()
}
