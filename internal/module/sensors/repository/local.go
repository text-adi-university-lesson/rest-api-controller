package repository

import (
	"fmt"
	"rest-api-controller/internal/module/sensors/model"

	"go.uber.org/fx"
)

type localParams struct {
	fx.In
}

type localRepository struct {
	storage []model.TemperatureSensor
}

func NewLocalRepository(p localParams) Repository {
	return &localRepository{
		storage: make([]model.TemperatureSensor, 0),
	}
}

func (r *localRepository) AddItem(item *model.TemperatureSensor) (*model.TemperatureSensor, error) {
	r.storage = append(r.storage, *item)
	return item, nil
}
func (r *localRepository) UpdateItem(data *model.TemperatureSensor) (*model.TemperatureSensor, error) {
	for i, item := range r.storage {
		if item.ID == data.ID {
			r.storage[i] = *data
		}
	}
	return data, nil
}
func (r *localRepository) GetAllItems() ([]model.TemperatureSensor, error) {
	return r.storage, nil
}
func (r *localRepository) DeleteByID(id string) error {
	for i, item := range r.storage {
		if item.ID == id {
			r.storage = append(r.storage[:i], r.storage[i+1:]...)
		}
	}
	return nil
}
func (r *localRepository) GetItemByID(id string) (*model.TemperatureSensor, error) {
	for _, item := range r.storage {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("item not found")
}
