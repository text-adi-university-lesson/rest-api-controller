package db

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User     string
	Name     string
	Password string
	Host     string
	Port     string
}

func (c *Config) GetDSN() string {
	return "host=" + c.Host + " user=" + c.User + " password=" + c.Password + " database=" + c.Name + " port=" + c.Port + " sslmode=disable TimeZone=UTC"
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	return &Config{
		User:     os.Getenv("DATABASE_USER"),
		Name:     os.Getenv("DATABASE_NAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
	}, nil
}
