package chain

import (
	"context"
	"math/big"
)

type Subscription interface {
	OnTokenBurned() <-chan Event
	OnCoinsLocked() <-chan Event
	Err() <-chan error
}

type Chain interface {
	Subscribe(ctx context.Context) (Subscription, error)
	MintToken(ctx context.Context, to string, coinId int, amount *big.Int) (string, *big.Int, error)
	UnlockCoins(ctx context.Context, account string, amount *big.Int) (string, *big.Int, error)
}
