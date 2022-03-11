package avalanche

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"heimdallr/internal/avalanche/locker"
)

type Subscription struct {
	onAVAXLocked chan LockEvent
	onUSDCLocked chan LockEvent

	errs chan error
}

type LockEvent struct {
	user        common.Address
	amount      *big.Int
	destination string
}

func (e LockEvent) User() string {
	return e.user.Hex()
}

func (e LockEvent) Amount() *big.Int {
	return e.amount
}

func (e LockEvent) Destination() string {
	return e.destination
}

func newSubscription() *Subscription {
	return &Subscription{
		onAVAXLocked: make(chan LockEvent),
		onUSDCLocked: make(chan LockEvent),
		errs:         make(chan error),
	}
}

func (s *Subscription) OnAVAXLocked() <-chan LockEvent {
	return s.onAVAXLocked
}

func (s *Subscription) OnUSDCLocked() <-chan LockEvent {
	return s.onUSDCLocked
}

func (s *Subscription) Err() <-chan error {
	return s.errs
}

func (s *Subscription) loopAVAX(
	ctx context.Context,
	avaxSub event.Subscription,
	avaxEvents <-chan *locker.LockerAVAXLocked,
) {
	for {
		select {
		case <-ctx.Done():
			avaxSub.Unsubscribe()
			return
		case ev := <-avaxEvents:
			s.onAVAXLocked <- LockEvent{
				user:        ev.User,
				amount:      ev.Amount,
				destination: ev.Destination,
			}
		case err := <-avaxSub.Err():
			s.errs <- err
		}
	}
}

func (s *Subscription) loopUSDC(
	ctx context.Context,
	usdcSub event.Subscription,
	usdcEvents <-chan *locker.LockerUSDCLocked,
) {
	for {
		select {
		case <-ctx.Done():
			usdcSub.Unsubscribe()
			return
		case ev := <-usdcEvents:
			s.onUSDCLocked <- LockEvent{
				user:        ev.User,
				amount:      ev.Amount,
				destination: ev.Destination,
			}
		case err := <-usdcSub.Err():
			s.errs <- err
		}
	}
}
