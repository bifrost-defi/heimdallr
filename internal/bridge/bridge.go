package bridge

import "bridge-oracle/internal/avalanche"

type Bridge struct {
	avalanche *avalanche.Avalanche
}

func New(avalanche *avalanche.Avalanche) *Bridge {
	return &Bridge{
		avalanche: avalanche,
	}
}
