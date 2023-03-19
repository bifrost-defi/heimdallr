package evm

import (
	"context"

	"github.com/ethereum/go-ethereum/event"
	"heimdallr/internal/chain"
	"heimdallr/internal/chain/evm/wrapping-bridge"
)

type Subscription struct {
	onTokenBurned chan chain.Event
	onCoinsLocked chan chain.Event
	errs          chan error
}

func newSubscription() *Subscription {
	return &Subscription{
		onTokenBurned: make(chan chain.Event),
		onCoinsLocked: make(chan chain.Event),
		errs:          make(chan error),
	}
}

func (s *Subscription) OnTokenBurned() <-chan chain.Event {
	return s.onTokenBurned
}

func (s *Subscription) OnCoinsLocked() <-chan chain.Event {
	return s.onCoinsLocked
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
			s.onCoinsLocked <- chain.NewEvent(
				ev.From.Hex(),
				ev.Value,
				int(ev.DestChain.Int64()),
				ev.DestAddress,
			)
		case err := <-sub.Err():
			s.errs <- err
		}
	}
}
