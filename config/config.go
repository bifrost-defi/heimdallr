package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Ethereum struct {
		RPC            string `env:"ETHEREUM_RPC_URL"`
		WS             string `env:"ETHEREUM_WS_URL"`
		BridgeContract string `env:"ETHEREUM__BRIDGE_CONTRACT"`
		PrivateKey     string `env:"ETHEREUM_PRIVATE_KEY"`
	}

	Tezos struct {
		URL            string `env:"TEZOS_URL"`
		BridgeContract string `env:"TEZOS_BRIDGE_CONTRACT"`
		PrivateKey     string `env:"TEZOS_PRIVATE_KEY"`
	}

	TON struct {
		URL            string `env:"TON_URL"`
		BridgeContract string `env:"TON_BRIDGE_CONTRACT"`
		PrivateKey     string `env:"TON_PRIVATE_KEY"`
	}
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return c, nil
}
