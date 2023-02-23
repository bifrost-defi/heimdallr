package bridge

import (
	"context"
	"fmt"
	"heimdallr/internal/chain"
	"heimdallr/internal/chain/evm"
	"heimdallr/internal/chain/tezos"
	"heimdallr/internal/chain/ton"
	"math/big"

	"go.uber.org/zap"
)

type Bridge struct {
	chains map[ChainID]chain.Chain

	logger *zap.SugaredLogger
}

type Event interface {
	User() string
	Amount() *big.Int
	CoinID() int
	Destination() string
}

func New(ethereum *evm.EVM, tezos *tezos.Tezos, ton *ton.TON, logger *zap.SugaredLogger) *Bridge {
	chains := map[ChainID]chain.Chain{
		EthereumID: ethereum,
		TezosID:    tezos,
		TonID:      ton,
	}

	return &Bridge{
		chains: chains,
		logger: logger,
	}
}

func (b *Bridge) Run(ctx context.Context) error {
	ethSub, err := b.chains[EthereumID].(*evm.EVM).Subscribe(ctx)
	if err != nil {
		return fmt.Errorf("subscribe evm: %w", err)
	}

	tzsSub, err := b.chains[TezosID].(*tezos.Tezos).Subscribe(ctx)
	if err != nil {
		return fmt.Errorf("subscribe tezos: %w", err)
	}

	tonSub, err := b.chains[TonID].(*ton.TON).Subscribe(ctx)
	if err != nil {
		return fmt.Errorf("subscribe ton: %w", err)
	}

	b.logger.Info("Heimdallr is watching")
	b.loop(ctx, ethSub, tzsSub, tonSub)

	return nil
}

func (b *Bridge) loop(ctx context.Context, ethSub *evm.Subscription, tzsSub *tezos.Subscription, sub *ton.Subscription) {
	atomic := NewAtomic(
		WithChecker(b.checkOperation),
	)

	for {
		select {
		// Break loop on interruption
		case <-ctx.Done():
			return

		// Handle events from chains and call another chain
		case event := <-ethSub.OnETHLocked():
			swap := atomic.NewOperation(
				WithName("TODO"),
			)
			go swap.Run(ctx, event)
		case event := <-tzsSub.OnTokenBurned():
			swap := atomic.NewOperation(
				WithName("TODO"),
			)
			go swap.Run(ctx, event)

		// Handle errors occurred during chains subscriptions
		case err := <-ethSub.Err():
			b.logger.Errorf("evm subscribtion error: %s", err)
		case err := <-tzsSub.Err():
			b.logger.Errorf("tezos subscribtion error: %s", err)
		case err := <-tzsSub.Err():
			b.logger.Errorf("ton subscribtion error: %s", err)
		}
	}
}

func (b *Bridge) checkOperation(op Checker, event Event) {
	select {
	case <-op.Complete():
		b.logger.With(
			zap.String("from", event.User()),
			zap.String("to", event.Destination()),
			zap.Int64("amount", event.Amount().Int64()),
		).Info("swap complete")
	case <-op.Rollback():
		b.logger.With(
			zap.String("from", event.User()),
			zap.String("to", event.Destination()),
			zap.Int64("amount", event.Amount().Int64()),
		).Info("swap rolled back")
	// Should not happen ever, because operation failing leads to coins lost.
	// Only contract owner will be able to unlock or mint lost coins.
	case err := <-op.Fail():
		b.logger.With(
			zap.String("from", event.User()),
			zap.String("to", event.Destination()),
			zap.Int64("amount", event.Amount().Int64()),
			zap.Error(err),
		).Debug("swap failed")
	}
}
