package database

import (
	"errors"
	"fmt"
	"orderapi/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConn[T config.DBConfig](dbConfig T) (*gorm.DB, error) {
	switch cfg := any(dbConfig).(type) {
	case *config.Postgres:
		return GetPGConn(cfg.Host, cfg.Port, cfg.DBName, cfg.User, cfg.Password)
	default:
		return nil, errors.New("unknown database config")
	}
}

func GetPGConn(host, port, dbname, user, password string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable", host, port, dbname, user, password)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
