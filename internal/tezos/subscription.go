package tezos

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Subscription struct {
	onWAVAXBurned chan BurnEvent
	onWUSDCBurned chan BurnEvent

	wavaxStorage *Storage
	wusdcStorage *Storage

	errs chan error
}

const (
	wavaxCheckInterval = 2 * time.Second
	wusdcCheckInterval = 2 * time.Second
)

type BurnEvent struct {
	user        tezos.Address
	amount      *big.Int
	destination string
}

func (e BurnEvent) User() string {
	return e.user.String()
}

func (e BurnEvent) Amount() *big.Int {
	return e.amount
}

func (e BurnEvent) Destination() string {
	return e.destination
}

func newSubscription(wavaxStorage *Storage, wusdcStorage *Storage) *Subscription {
	return &Subscription{
		onWAVAXBurned: make(chan BurnEvent),
		onWUSDCBurned: make(chan BurnEvent),
		wavaxStorage:  wavaxStorage,
		wusdcStorage:  wusdcStorage,
		errs:          make(chan error),
	}
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

func (s *Subscription) loop(ctx context.Context) {
	wavaxTick := time.NewTicker(wavaxCheckInterval)
	wusdcTick := time.NewTicker(wusdcCheckInterval)
	defer wavaxTick.Stop()
	defer wusdcTick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-wavaxTick.C:
			s.checkWAVAXBurnings(ctx)
		case <-wusdcTick.C:
			s.checkWUSDCBurnings(ctx)
		}
	}
}

func (s *Subscription) checkWAVAXBurnings(ctx context.Context) {
	burnings, err := s.wavaxStorage.UpdateBurnings(ctx)
	if err != nil {
		s.errs <- fmt.Errorf("update wavax burnings: %w", err)
	}

	for _, v := range burnings {
		s.onWAVAXBurned <- BurnEvent{
			user:        v.User,
			amount:      v.Amount.Big(),
			destination: v.Destination,
		}
	}
}

func (s *Subscription) checkWUSDCBurnings(ctx context.Context) {
	burnings, err := s.wusdcStorage.UpdateBurnings(ctx)
	if err != nil {
		s.errs <- fmt.Errorf("update wavax burnings: %w", err)
	}

	for _, v := range burnings {
		s.onWUSDCBurned <- BurnEvent{
			user:        v.User,
			amount:      v.Amount.Big(),
			destination: v.Destination,
		}
	}
}
