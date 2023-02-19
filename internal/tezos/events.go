package tezos

import (
	"blockwatch.cc/tzgo/tezos"
	"math/big"
	"time"
)

type MichelsonEvent struct {
	User        tezos.Address `json:"user"`
	Amount      tezos.Z       `json:"amount"`
	Destination string        `json:"destAddress"`
	CoinID      int           `json:"destCoinId"`
	Timestamp   time.Time     `json:"ts"`
}

type Event struct {
	user        tezos.Address
	amount      *big.Int
	coinId      int
	destination string
}

func (e Event) User() string {
	return e.user.String()
}

func (e Event) Amount() *big.Int {
	return e.amount
}

func (e Event) CoinID() int {
	return e.coinId
}

func (e Event) Destination() string {
	return e.destination
}
