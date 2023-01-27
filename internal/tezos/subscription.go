package tezos

import (
	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/rpc"
	"context"
	"fmt"
	"math/big"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Subscription struct {
	onTokenBurned chan BurnEvent

	contract   *contract.Contract
	client     *rpc.Client
	blockLevel int64

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

func newSubscription(contract *contract.Contract) *Subscription {
	return &Subscription{
		onTokenBurned: make(chan BurnEvent),
		contract:      contract,
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
		// TODO: parse result
	}

	for _, r := range results {
		if r.Destination.Equal(s.contract.Address()) && r.Kind == tezos.OpTypeEvent {
			go collector(r)
		}
	}
}
