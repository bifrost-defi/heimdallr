package tezos

import (
	"blockwatch.cc/tzgo/tezos"
	"time"
)

type MichelsonEvent struct {
	User        tezos.Address `json:"user"`
	Amount      tezos.Z       `json:"amount"`
	Destination string        `json:"destAddress"`
	CoinID      int           `json:"destCoinId"`
	Timestamp   time.Time     `json:"ts"`
}
