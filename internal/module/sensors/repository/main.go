package repository

import "rest-api-controller/internal/module/sensors/model"

type Repository interface {
	AddItem(*model.TemperatureSensor) (*model.TemperatureSensor, error)
	UpdateItem(*model.TemperatureSensor) (*model.TemperatureSensor, error)
	GetAllItems() ([]model.TemperatureSensor, error)
	DeleteByID(string) error
	GetItemByID(string) (*model.TemperatureSensor, error)
}
