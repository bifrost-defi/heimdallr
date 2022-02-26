package tezos

import (
	"math/big"

	"blockwatch.cc/tzgo/tezos"
)

type Subscription struct {
	onWAVAXBurned chan BurnEvent
	onWUSDCBurned chan BurnEvent

	errs chan error
}

type BurnEvent struct {
	user        tezos.Address
	amount      *big.Int
	destination tezos.Address
}

func (e BurnEvent) User() string {
	return e.user.String()
}

func (e BurnEvent) Amount() *big.Int {
	return e.amount
}

func (e BurnEvent) Destination() string {
	return e.destination.String()
}

func (s *Subscription) OnWAVAXBurned() <-chan BurnEvent {
	return s.onWAVAXBurned
}

func (s *Subscription) OnWUSDCBurned() <-chan BurnEvent {
	return s.onWUSDCBurned
}

func (s *Subscription) Err() <-chan error {
	return s.errs
}
