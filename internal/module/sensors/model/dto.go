package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateTemperatureSensorsDTO struct {
	Timestamp  *time.Time `json:"timestamp"`
	SensorName string     `json:"sensor_name" validate:"required"`
	Value      float64    `json:"value" validate:"required"`
	Unit       string     `json:"unit" validate:"required"`
}

func (dto *CreateTemperatureSensorsDTO) ToModel() TemperatureSensor {
	var ts = time.Now()
	if dto.Timestamp != nil {
		ts = *dto.Timestamp
	}

	return TemperatureSensor{
		ID:         uuid.NewString(),
		Timestamp:  ts,
		SensorName: dto.SensorName,
		Value:      dto.Value,
		Unit:       dto.Unit,
	}
}

type UpdateTemperatureSensorsDTO struct {
	Timestamp  time.Time `json:"timestamp" validate:"required"`
	SensorName string    `json:"sensor_name" validate:"required"`
	Value      float64   `json:"value" validate:"required"`
	Unit       string    `json:"unit" validate:"required"`
}

func (dto *UpdateTemperatureSensorsDTO) ToModel(id string) TemperatureSensor {
	return TemperatureSensor{
		ID:         id,
		Timestamp:  dto.Timestamp,
		SensorName: dto.SensorName,
		Value:      dto.Value,
		Unit:       dto.Unit,
	}
}

type CreateHumiditySensorsDTO struct {
	Timestamp  *time.Time `json:"timestamp"`
	SensorName string     `json:"sensor_name" validate:"required"`
	Value      float64    `json:"value" validate:"required"`
	Unit       string     `json:"unit" validate:"required"`
}

func (dto *CreateHumiditySensorsDTO) ToModel() HumiditySensor {
	var ts = time.Now()
	if dto.Timestamp != nil {
		ts = *dto.Timestamp
	}

	return HumiditySensor{
		ID:         uuid.NewString(),
		Timestamp:  ts,
		SensorName: dto.SensorName,
		Value:      dto.Value,
		Unit:       dto.Unit,
	}
}

type UpdateHumiditySensorsDTO struct {
	Timestamp  time.Time `json:"timestamp" validate:"required"`
	SensorName string    `json:"sensor_name" validate:"required"`
	Value      float64   `json:"value" validate:"required"`
	Unit       string    `json:"unit" validate:"required"`
}

func (dto *UpdateHumiditySensorsDTO) ToModel(id string) HumiditySensor {
	return HumiditySensor{
		ID:         id,
		Timestamp:  dto.Timestamp,
		SensorName: dto.SensorName,
		Value:      dto.Value,
		Unit:       dto.Unit,
	}
}
