package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Ethereum Ethereum
	Tezos    Tezos
	TON      TON
}

type Ethereum struct {
	RPC            string `env:"ETHEREUM_RPC_URL"`
	WS             string `env:"ETHEREUM_WS_URL"`
	BridgeContract string `env:"ETHEREUM__BRIDGE_CONTRACT"`
	PrivateKey     string `env:"ETHEREUM_PRIVATE_KEY"`
}

type Tezos struct {
	URL            string `env:"TEZOS_URL"`
	BridgeContract string `env:"TEZOS_BRIDGE_CONTRACT"`
	PrivateKey     string `env:"TEZOS_PRIVATE_KEY"`
}

type TON struct {
	ConfigURL      string `env:"TON_CONFIG_URL"`
	BridgeContract string `env:"TON_BRIDGE_CONTRACT"`
	WalletSeed     string `env:"TON_WALLET_SEED"`
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return c, nil
}
