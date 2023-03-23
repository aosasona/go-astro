package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type DatabaseConfig struct {
	DB       string
	Host     string
	User     string
	Password string
	Port     int
}

func Connect(config DatabaseConfig) (*xorm.Engine, error) {
	dbn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DB,
	)
	engine, err := xorm.NewEngine("postgres", dbn)
	if err != nil {
		return nil, err
	}

	return engine, nil
}
