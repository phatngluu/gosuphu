package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr      string `env:"SERVICE_ADDR,required"`
	Version         string `env:"SERVICE_VERSION,required"`
	DBConnectString string `env:"SERVICE_DB_CONNECTION_STRING,required"`
}

func NewConfig() *Config {
	godotenv.Load()
	var config = &Config{}
	if err := env.Parse(config); err != nil {
		panic(fmt.Errorf("%+v", err))
	}
	return config
}
