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

func (receiver Service) DeleteItem(id string) error {
	return receiver.repo.DeleteByID(id)
}
func (receiver Service) CreateItem(item *model.TemperatureSensor) (*model.TemperatureSensor, error) {
	return receiver.repo.AddItem(item)
}
func (receiver Service) UpdateItem(item *model.TemperatureSensor) (*model.TemperatureSensor, error) {
	return receiver.repo.UpdateItem(item)
}
func (receiver Service) GetItemByID(id string) (*model.TemperatureSensor, error) {
	return receiver.repo.GetItemByID(id)
}
func (receiver Service) GetAll() ([]model.TemperatureSensor, error) {
	return receiver.repo.GetAllItems()
}
