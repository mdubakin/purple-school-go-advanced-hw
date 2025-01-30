package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server   `yaml:"server"`
	Database `yaml:"database"`
}

type Server struct {
	Port string `yaml:"port" env:"SERVER_PORT" env-required:"true"`
}

type DBConfig interface {
	*Postgres
	isConfig()
}

type Database struct {
	Postgres `yaml:"postgres"`
}

type Postgres struct {
	DBName   string `yaml:"dbname" env:"DB_PG_DBNAME" env-required:"true"`
	Host     string `yaml:"host" env:"DB_PG_HOST" env-required:"true"`
	Port     string `yaml:"port" env:"DB_PG_PORT" env-required:"true"`
	User     string `yaml:"user" env:"DB_PG_USER" env-required:"true"`
	Password string `yaml:"password" env:"DB_PG_PASSWORD" env-required:"true"`
}

func (p *Postgres) isConfig() {}

func MustLoadConfig() *Config {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		log.Panicf("error loading config.yaml file: %v", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Panicf("error loading .env file: %v", err)
	}

	return cfg
}
