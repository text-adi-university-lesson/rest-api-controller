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
	storageTemperature []model.TemperatureSensor
	storageHumidity    []model.HumiditySensor
}

func NewLocalRepository(p localParams) Repository {
	return &localRepository{
		storageTemperature: make([]model.TemperatureSensor, 0),
		storageHumidity:    make([]model.HumiditySensor, 0),
	}
}

func (r *localRepository) AddHumidityItem(sensor *model.HumiditySensor) (*model.HumiditySensor, error) {
	r.storageHumidity = append(r.storageHumidity, *sensor)
	return sensor, nil
}

func (r *localRepository) UpdateHumidityItem(sensor *model.HumiditySensor) (*model.HumiditySensor, error) {
	for i, item := range r.storageTemperature {
		if item.ID == sensor.ID {
			r.storageHumidity[i] = *sensor
		}
	}
	return sensor, nil
}

func (r *localRepository) GetAllHumidityItems() ([]model.HumiditySensor, error) {
	return r.storageHumidity, nil
}

func (r *localRepository) DeleteHumidityByID(s string) error {
	for i, item := range r.storageHumidity {
		if item.ID == s {
			r.storageHumidity = append(r.storageHumidity[:i], r.storageHumidity[i+1:]...)
		}
	}
	return nil
}

func (r *localRepository) GetHumidityItemByID(s string) (*model.HumiditySensor, error) {
	for _, item := range r.storageHumidity {
		if item.ID == s {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("item not found")
}

func (r *localRepository) AddTemperatureItem(item *model.TemperatureSensor) (*model.TemperatureSensor, error) {
	r.storageTemperature = append(r.storageTemperature, *item)
	return item, nil
}
func (r *localRepository) UpdateTemperatureItem(data *model.TemperatureSensor) (*model.TemperatureSensor, error) {
	for i, item := range r.storageTemperature {
		if item.ID == data.ID {
			r.storageTemperature[i] = *data
		}
	}
	return data, nil
}
func (r *localRepository) GetAllTemperatureItems() ([]model.TemperatureSensor, error) {
	return r.storageTemperature, nil
}
func (r *localRepository) DeleteTemperatureByID(id string) error {
	for i, item := range r.storageTemperature {
		if item.ID == id {
			r.storageTemperature = append(r.storageTemperature[:i], r.storageTemperature[i+1:]...)
		}
	}
	return nil
}
func (r *localRepository) GetTemperatureItemByID(id string) (*model.TemperatureSensor, error) {
	for _, item := range r.storageTemperature {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("item not found")
}
