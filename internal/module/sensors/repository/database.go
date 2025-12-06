package repository

import (
	"fmt"
	"rest-api-controller/internal/module/sensors/model"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type databaseRepository struct {
	db *gorm.DB
}

func (d *databaseRepository) AddTemperatureItem(sensor *model.TemperatureSensor) (*model.TemperatureSensor, error) {
	if err := d.db.Create(sensor).Error; err != nil {
		return nil, err
	}
	return sensor, nil
}

func (d *databaseRepository) UpdateTemperatureItem(sensor *model.TemperatureSensor) (*model.TemperatureSensor, error) {
	if err := d.db.Save(sensor).Error; err != nil {
		return nil, err
	}
	return sensor, nil
}

func (d *databaseRepository) GetAllTemperatureItems() ([]model.TemperatureSensor, error) {
	var items []model.TemperatureSensor
	if err := d.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (d *databaseRepository) DeleteTemperatureByID(s string) error {
	res := d.db.Delete(&model.TemperatureSensor{}, "id = ?", s)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("item not found")
	}
	return nil
}

func (d *databaseRepository) GetTemperatureItemByID(s string) (*model.TemperatureSensor, error) {
	var item model.TemperatureSensor
	if err := d.db.First(&item, "id = ?", s).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("item not found")
		}
		return nil, err
	}
	return &item, nil
}

func (d *databaseRepository) AddHumidityItem(sensor *model.HumiditySensor) (*model.HumiditySensor, error) {
	if err := d.db.Create(sensor).Error; err != nil {
		return nil, err
	}
	return sensor, nil
}

func (d *databaseRepository) UpdateHumidityItem(sensor *model.HumiditySensor) (*model.HumiditySensor, error) {
	if err := d.db.Save(sensor).Error; err != nil {
		return nil, err
	}
	return sensor, nil
}

func (d *databaseRepository) GetAllHumidityItems() ([]model.HumiditySensor, error) {
	var items []model.HumiditySensor
	if err := d.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (d *databaseRepository) DeleteHumidityByID(s string) error {
	res := d.db.Delete(&model.HumiditySensor{}, "id = ?", s)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("item not found")
	}
	return nil
}

func (d *databaseRepository) GetHumidityItemByID(s string) (*model.HumiditySensor, error) {
	var item model.HumiditySensor
	if err := d.db.First(&item, "id = ?", s).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("item not found")
		}
		return nil, err
	}
	return &item, nil
}

type dbParams struct {
	fx.In

	DB *gorm.DB
}

func NewDatabaseRepository(p dbParams) Repository {
	// use db for queries
	return &databaseRepository{db: p.DB}
}
