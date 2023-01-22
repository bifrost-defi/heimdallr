package evm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"heimdallr/internal/evm/locker"
)

type Subscription struct {
	onETHLocked chan LockEvent
	errs        chan error
}

type LockEvent struct {
	user        common.Address
	amount      *big.Int
	coinId      int
	destination string
}

func (e LockEvent) User() string {
	return e.user.Hex()
}

func (e LockEvent) Amount() *big.Int {
	return e.amount
}

func (e LockEvent) CoinID() int {
	return e.coinId
}

func (e LockEvent) Destination() string {
	return e.destination
}

func newSubscription() *Subscription {
	return &Subscription{
		onETHLocked: make(chan LockEvent),
		errs:        make(chan error),
	}
}

func (s *Subscription) OnETHLocked() <-chan LockEvent {
	return s.onETHLocked
}

func (s *Subscription) Err() <-chan error {
	return s.errs
}

func (s *Subscription) loop(
	ctx context.Context,
	sub event.Subscription,
	events <-chan *locker.LockerAVAXLocked,
) {
	for {
		select {
		case <-ctx.Done():
			sub.Unsubscribe()
			return
		case ev := <-events:
			s.onETHLocked <- LockEvent{
				user:        ev.User,
				amount:      ev.Amount,
				destination: ev.Destination,
			}
		case err := <-sub.Err():
			s.errs <- err
		}
	}
}
