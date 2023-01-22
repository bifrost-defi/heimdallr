package tezos

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Subscription struct {
	onTokenBurned chan BurnEvent

	bridgeStorage *Storage

	errs chan error
}

const checkInterval = 3 * time.Second

type BurnEvent struct {
	user        tezos.Address
	amount      *big.Int
	coinId      int
	destination string
}

func (e BurnEvent) User() string {
	return e.user.String()
}

func (e BurnEvent) Amount() *big.Int {
	return e.amount
}

func (e BurnEvent) CoinID() int {
	return e.coinId
}

func (e BurnEvent) Destination() string {
	return e.destination
}

func newSubscription(bridgeStorage *Storage) *Subscription {
	return &Subscription{
		onTokenBurned: make(chan BurnEvent),
		bridgeStorage: bridgeStorage,
		errs:          make(chan error),
	}
}

func (s *Subscription) OnTokenBurned() <-chan BurnEvent {
	return s.onTokenBurned
}

func (s *Subscription) Err() <-chan error {
	return s.errs
}

func (s *Subscription) loop(ctx context.Context) {
	tick := time.NewTicker(checkInterval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			s.checkBurnings(ctx)
		}
	}
}

func (s *Subscription) checkBurnings(ctx context.Context) {
	burnings, err := s.bridgeStorage.UpdateBurnings(ctx)
	if err != nil {
		s.errs <- fmt.Errorf("update burnings: %w", err)
	}

	for _, v := range burnings {
		s.onTokenBurned <- BurnEvent{
			user:        v.User,
			amount:      v.Amount.Big(),
			coinId:      v.CoinID,
			destination: v.Destination,
		}
	}
}
