package bridge

import (
	"context"
	"fmt"
	"math/big"

	"go.uber.org/zap"
	"heimdallr/internal/evm"
	"heimdallr/internal/tezos"
	"heimdallr/internal/ton"
)

type Bridge struct {
	ethereum *evm.EVM
	tezos    *tezos.Tezos
	ton      *ton.TON

	logger *zap.SugaredLogger
}

type Event interface {
	User() string
	Amount() *big.Int
	CoinID() int
	Destination() string
}

func New(ethereum *evm.EVM, tezos *tezos.Tezos, ton *ton.TON, logger *zap.SugaredLogger) *Bridge {
	return &Bridge{
		ethereum: ethereum,
		tezos:    tezos,
		ton:      ton,
		logger:   logger,
	}
}

func (b *Bridge) Run(ctx context.Context) error {
	ethSub, err := b.ethereum.Subscribe(ctx)
	if err != nil {
		return fmt.Errorf("subscribe evm: %w", err)
	}

	tzsSub, err := b.tezos.Subscribe(ctx)
	if err != nil {
		return fmt.Errorf("subscribe tezos: %w", err)
	}

	b.logger.Info("Heimdallr is watching")
	b.loop(ctx, ethSub, tzsSub)

	return nil
}

func (b *Bridge) loop(ctx context.Context, ethSub *evm.Subscription, tzsSub *tezos.Subscription) {
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
				OnPerform(b.mintToken),
				OnRollback(b.unlockETH),
			)
			go swap.Run(ctx, event)
		case event := <-tzsSub.OnTokenBurned():
			swap := atomic.NewOperation(
				WithName("TODO"),
				OnPerform(b.unlockETH),
				OnRollback(b.mintToken),
			)
			go swap.Run(ctx, event)

		// Handle errors occurred during chains subscriptions
		case err := <-ethSub.Err():
			b.logger.Errorf("evm subscribtion error: %s", err)
		case err := <-tzsSub.Err():
			b.logger.Errorf("tezos subscribtion error: %s", err)
		}
	}
}

func (b *Bridge) mintToken(ctx context.Context, event Event) bool {
	hash, fee, err := b.tezos.MintToken(ctx, event.Destination(), event.CoinID(), event.Amount())
	if err != nil {
		b.logger.Errorf("mint token: %s", err)

		return false
	}

	b.logger.With(
		zap.String("user", event.User()),
		zap.Int64("amount", event.Amount().Int64()),
		zap.String("destination", event.Destination()),
		zap.String("tx_hash", hash),
		zap.Int64("fee", fee.Int64()),
	).Info("token minted")

	return true
}

func (b *Bridge) unlockETH(ctx context.Context, event Event) bool {
	hash, fee, err := b.ethereum.UnlockETH(ctx, event.Destination(), event.Amount())
	if err != nil {
		b.logger.Errorf("unlock eth: %s", err)

		return false
	}

	b.logger.With(
		zap.String("user", event.User()),
		zap.Int64("amount", event.Amount().Int64()),
		zap.String("destination", event.Destination()),
		zap.String("tx_hash", hash),
		zap.Int64("fee", fee.Int64()),
	).Info("eth unlocked")

	return true
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
