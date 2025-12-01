package repository

import "rest-api-controller/internal/module/sensors/model"

type Repository interface {
	AddTemperatureItem(*model.TemperatureSensor) (*model.TemperatureSensor, error)
	UpdateTemperatureItem(*model.TemperatureSensor) (*model.TemperatureSensor, error)
	GetAllTemperatureItems() ([]model.TemperatureSensor, error)
	DeleteTemperatureByID(string) error
	GetTemperatureItemByID(string) (*model.TemperatureSensor, error)

	AddHumidityItem(*model.HumiditySensor) (*model.HumiditySensor, error)
	UpdateHumidityItem(*model.HumiditySensor) (*model.HumiditySensor, error)
	GetAllHumidityItems() ([]model.HumiditySensor, error)
	DeleteHumidityByID(string) error
	GetHumidityItemByID(string) (*model.HumiditySensor, error)
}
