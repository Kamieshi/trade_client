package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// Configuration config storage from Application
type Configuration struct {
	PositionServerRPCHost string `env:"POSITION_SERVER_RPC_HOST"`
	PositionServerRPCPort string `env:"POSITION_SERVER_RPC_PORT"`
	PriceServerRPCHost    string `env:"PRICE_SERVER_RPC_HOST"`
	PriceServerRPCPort    string `env:"PRICE_SERVER_RPC_PORT"`
	HttpEchoPort          string `env:"HTTP_ECHO_PORT"`
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

func (c *Configuration) PositionServerRPCAddr() string {
	return fmt.Sprintf("%s:%s", c.PositionServerRPCHost, c.PositionServerRPCPort)
}

func (c *Configuration) PriceServerRPCAddr() string {
	return fmt.Sprintf("%s:%s", c.PriceServerRPCHost, c.PriceServerRPCPort)
}
