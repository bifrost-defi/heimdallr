package bridge

import (
	"context"
	"fmt"

	"bridge-oracle/internal/avalanche"
	"bridge-oracle/internal/tezos"
)

type Bridge struct {
	avalanche *avalanche.Avalanche
	tezos     *tezos.Tezos
}

func New(avalanche *avalanche.Avalanche, tezos *tezos.Tezos) *Bridge {
	return &Bridge{
		avalanche: avalanche,
		tezos:     tezos,
	}
}

func (b *Bridge) Run(ctx context.Context) error {
	avaSub, err := b.avalanche.Subscribe(ctx)
	if err != nil {
		return fmt.Errorf("subscribe avalanche: %w", err)
	}

	tzsSub, err := b.tezos.Subscribe(ctx)
	if err != nil {
		return fmt.Errorf("subscribe tezos: %w", err)
	}

	b.loop(ctx, avaSub, tzsSub)

	return nil
}

func (b *Bridge) loop(ctx context.Context, avaSub *avalanche.Subscription, tzsSub *tezos.Subscription) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-avaSub.OnAVAXLocked():
		case <-avaSub.OnUSDCLocked():
		case <-tzsSub.OnWAVAXBurned():
		case <-tzsSub.OnWUSDCBurned():
		}
	}
}
