package db

import (
	"log"

	"rest-api-controller/internal/module/sensors/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg *Config) (*gorm.DB, error) {
	dsn := cfg.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// create tables if they don't exist
	if err := db.AutoMigrate(
		&model.TemperatureSensor{},
		&model.HumiditySensor{},
	); err != nil {
		log.Printf("db migrate error: %v", err)
		return nil, err
	}

	return db, nil
}
