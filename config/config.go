package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Avalanche struct {
		URL        string `env:"AVALANCHE_URL"`
		Contract   string `env:"AVALANCHE_CONTRACT"`
		PrivateKey string `env:"AVALANCHE_PRIVATE_KEY"`
	}

	Tezos struct {
		URL           string `env:"TEZOS_URL"`
		WAVAXContract string `env:"TEZOS_WAVAX_CONTRACT"`
		WUSDCContract string `env:"TEZOS_WUSDC_CONTRACT"`
		PrivateKey    string `env:"TEZOS_PRIVATE_KEY"`
	}
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return c, nil
}
