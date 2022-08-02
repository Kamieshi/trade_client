package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// Configuration config storage from Application
type Configuration struct {
	PositionServerRPC string `env:"POSITION_SERVER_RPC_ADDR"`
	PriceServerRPC    string `env:"PRICE_SERVER_RPC_ADDR"`
	HttpEchoPort      string `env:"HTTP_ECHO_PORT"`
}

// GetConfig init configuration from OS ENV
func GetConfig() (*Configuration, error) {
	conf := Configuration{}
	err := env.Parse(&conf)
	if err != nil {
		return nil, fmt.Errorf("config / GetConfig / error parse : %v", err)
	}
	return &conf, nil
}
