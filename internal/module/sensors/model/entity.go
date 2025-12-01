package model

import "time"

type TemperatureSensor struct {
	ID         string        `json:"id"`
	Timestamp  time.Duration `json:"timestamp"`
	SensorName string        `json:"sensor_name"`
	Value      float64       `json:"value"`
	Unit       string        `json:"unit"`
}

type HumiditySensor struct {
	ID         string        `json:"id"`
	Timestamp  time.Duration `json:"timestamp"`
	SensorName string        `json:"sensor_name"`
	Value      float64       `json:"value"`
	Unit       string        `json:"unit"`
}
