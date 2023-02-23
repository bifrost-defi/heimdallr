package chain

import (
	"context"
	"math/big"
)

type Chain interface {
	MintToken(ctx context.Context, destination string, coinId int, amount *big.Int) (string, *big.Int, error)
	UnlockCoins(ctx context.Context, account string, amount *big.Int) (string, *big.Int, error)
}
