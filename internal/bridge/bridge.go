package bridge

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"heimdallr/internal/chain"
	"heimdallr/internal/chain/evm"
	"heimdallr/internal/chain/tezos"
	"heimdallr/internal/chain/ton"
)

type Bridge struct {
	chains map[ChainID]chain.Chain

	logger *zap.SugaredLogger
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

func (b *Bridge) loop(ctx context.Context, ethSub chain.Subscription, tzsSub chain.Subscription, tonSub chain.Subscription) {
	atomicWrap := NewAtomic(
		WithChecker(b.checkOperation),
		OnPerform(b.wrap),
		OnRollback(b.unwrap),
	)

	atomicUnwrap := NewAtomic(
		WithChecker(b.checkOperation),
		OnPerform(b.wrap),
		OnRollback(b.unwrap),
	)

	for {
		select {
		// Break loop on interruption
		case <-ctx.Done():
			return

		// Handle events from chains and call another chain
		case event := <-tonSub.OnCoinsLocked():
			swap := atomicWrap.NewOperation(
				WithName("TON coins locked"),
			)
			go swap.Run(ctx, event)
		case event := <-tonSub.OnTokenBurned():
			swap := atomicUnwrap.NewOperation(
				WithName("TON jetton burned"),
			)
			go swap.Run(ctx, event)
		case event := <-ethSub.OnCoinsLocked():
			swap := atomicWrap.NewOperation(
				WithName("ETH coins locked"),
			)
			go swap.Run(ctx, event)
		case event := <-ethSub.OnTokenBurned():
			swap := atomicUnwrap.NewOperation(
				WithName("Ethereum token burned"),
			)
			go swap.Run(ctx, event)
		case event := <-tzsSub.OnCoinsLocked():
			swap := atomicWrap.NewOperation(
				WithName("Tezos coins locked"),
			)
			go swap.Run(ctx, event)
		case event := <-tzsSub.OnTokenBurned():
			swap := atomicUnwrap.NewOperation(
				WithName("Tezos token burned"),
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

func (b *Bridge) wrap(ctx context.Context, event chain.Event) (success bool) {
	chainID := ChainID(event.CoinID())
	destChain, ok := b.chains[chainID]
	if !ok {
		b.logger.Errorf("unknown chain id: %d", chainID)

		return false
	}

	hash, fee, err := destChain.MintToken(ctx, event.Destination(), event.CoinID(), event.Amount())
	if err != nil {
		b.logger.Errorf("mint token: %e", err)

		return false
	}

	b.logger.With(
		zap.String("user", event.User()),
		zap.Int64("amount", event.Amount().Int64()),
		zap.Int("chain_id", event.CoinID()),
		zap.String("destination", event.Destination()),
		zap.String("tx_hash", hash),
		zap.Int64("fee", fee.Int64()),
	).Info("token minted")

	return true
}

func (b *Bridge) unwrap(ctx context.Context, event chain.Event) (success bool) {
	chainID := ChainID(event.CoinID())
	destChain, ok := b.chains[chainID]
	if !ok {
		b.logger.Errorf("unknown chain id: %d", chainID)

		return false
	}

	hash, fee, err := destChain.UnlockCoins(ctx, event.Destination(), event.Amount())
	if err != nil {
		b.logger.Errorf("unlock coins: %e", err)

		return false
	}

	b.logger.With(
		zap.String("user", event.User()),
		zap.Int64("amount", event.Amount().Int64()),
		zap.Int("chain_id", event.CoinID()),
		zap.String("destination", event.Destination()),
		zap.String("tx_hash", hash),
		zap.Int64("fee", fee.Int64()),
	).Info("coins unlocked")

	return true
}

func (b *Bridge) checkOperation(op Checker, event chain.Event) {
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
