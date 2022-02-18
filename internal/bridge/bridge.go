package bridge

import (
	"context"
	"fmt"

	"bridge-oracle/internal/avalanche"
	"bridge-oracle/internal/tezos"
)

type Bridge struct {
	avaContract string
	tzsContract string

	avalanche *avalanche.Avalanche
	tezos     *tezos.Tezos
}

func New(avalanche *avalanche.Avalanche, tezos *tezos.Tezos) *Bridge {
	return &Bridge{
		avalanche: avalanche,
		tezos:     tezos,
	}
}

func (b *Bridge) SetAvalancheContract(address string) {
	b.avaContract = address
}

func (b *Bridge) SetTezosContract(address string) {
	b.tzsContract = address
}

func (b *Bridge) Run(ctx context.Context) error {
	sub, err := b.avalanche.Subscribe(ctx, b.avaContract)
	if err != nil {
		return fmt.Errorf("subscribe avalanche: %w", err)
	}

	b.loop(ctx, sub)

	return nil
}

func (b *Bridge) loop(ctx context.Context, sub *avalanche.Subscription) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-sub.OnAVAXLocked():
		case <-sub.OnUSDCLocked():
		}
	}
}
