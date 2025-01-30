package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server      `yaml:"server"`
	Database    `yaml:"database"`
	EmailConfig `yaml:"email"`
}

type Server struct {
	Port string `yaml:"port" env:"SERVER_PORT" env-required:"true"`
}

type Database struct {
	LocalJSONConfig `yaml:"local-json"`
}

type LocalJSONConfig struct {
	Path string `yaml:"path" env:"LOCAL_JSON_PATH" env-required:"true"`
}

type EmailConfig struct {
	Login    string `yaml:"login" env:"SMTP_EMAIL" env-required:"true"`
	Password string `yaml:"password" env:"SMTP_PASSWORD" env-required:"true"`
	SMTPHost string `yaml:"smtp_host" env:"SMTP_HOST" env-required:"true"`
	SMTPPort string `yaml:"smtp_port" env:"SMTP_PORT" env-required:"true"`
}

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
