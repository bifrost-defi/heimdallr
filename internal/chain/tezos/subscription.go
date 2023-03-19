package tezos

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/rpc"
	"blockwatch.cc/tzgo/tezos"
	"heimdallr/internal/chain"
)

type Subscription struct {
	onTokenBurned chan chain.Event
	onCoinsLocked chan chain.Event

	contract   *contract.Contract
	client     *rpc.Client
	blockLevel int64

	errs chan error
}

const checkInterval = 3 * time.Second

func newSubscription(contract *contract.Contract) *Subscription {
	return &Subscription{
		onTokenBurned: make(chan chain.Event),
		onCoinsLocked: make(chan chain.Event),
		contract:      contract,
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

func (s *Subscription) loop(ctx context.Context) {
	tick := time.NewTicker(checkInterval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			s.checkOperations(ctx)
		}
	}
}

func (s *Subscription) checkOperations(ctx context.Context) {
	b, err := s.client.GetBlock(ctx, rpc.Head)
	if err != nil {
		s.errs <- fmt.Errorf("get last block: %w", err)

		return
	}

	for s.blockLevel <= b.GetLevel() {
		ops, err := s.client.GetBlockOperationList(ctx, rpc.BlockLevel(s.blockLevel), 0)
		if err != nil {
			s.errs <- fmt.Errorf("get operations for block %d: %w", s.blockLevel, err)

			return
		}

		for _, op := range ops {
			for _, c := range op.Contents {
				if c.Kind() == tezos.OpTypeTransaction {
					go s.collectEvents(c.Meta().InternalResults)
				}
			}
		}
	}
}

func (s *Subscription) collectEvents(results []*rpc.InternalResult) {
	collector := func(result *rpc.InternalResult) {
		eventData := result.Payload.Args.Last().Bytes

		var event MichelsonEvent
		if err := json.Unmarshal(eventData, &event); err != nil {
			s.errs <- fmt.Errorf("unmarshal event: %w", err)

			return
		}

		switch result.Tag {
		case "lock":
			s.onCoinsLocked <- chain.NewEvent(
				event.User.String(),
				event.Amount.Big(),
				event.CoinID,
				event.Destination,
			)
		case "burn":
			s.onTokenBurned <- chain.NewEvent(
				event.User.String(),
				event.Amount.Big(),
				event.CoinID,
				event.Destination,
			)
		}
	}

	for _, r := range results {
		if r.Destination.Equal(s.contract.Address()) && r.Kind == tezos.OpTypeEvent {
			go collector(r)
		}
	}
}
