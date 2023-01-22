package evm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	wrappingBridge "heimdallr/internal/evm/wrapping-bridge"
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
	events <-chan *wrappingBridge.WrappingBridgeLock,
) {
	for {
		select {
		case <-ctx.Done():
			sub.Unsubscribe()
			return
		case ev := <-events:
			s.onETHLocked <- LockEvent{
				user:        ev.From,
				amount:      ev.Value,
				coinId:      int(ev.DestChain.Int64()),
				destination: ev.DestAddress,
			}
		case err := <-sub.Err():
			s.errs <- err
		}
	}
}
