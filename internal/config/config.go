package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Configuration struct {
	PositionServerRPC string `env:"POSITION_SERVER_RPC_ADDR"`
	PriceServerRPC    string `env:"PRICE_SERVER_RPC_ADDR"`
	HttpPort          string `env:"HTTP_PORT"`
}

func GetConfig() (*Configuration, error) {
	conf := Configuration{}
	err := env.Parse(&conf)
	if err != nil {
		return nil, fmt.Errorf("config / GetConfig / error parse : %v", err)
	}
	return &conf, nil
}
