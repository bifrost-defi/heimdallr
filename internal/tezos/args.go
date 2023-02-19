package tezos

import (
	"blockwatch.cc/tzgo/codec"
	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

type TokenMint struct {
	CoinID int           `json:"coinId"`
	To     tezos.Address `json:"to_"`
	Value  tezos.Z       `json:"value"`
}

type TokenMintArgs struct {
	contract.TxArgs
	TokenMint TokenMint
}

func (a TokenMintArgs) Parameters() *micheline.Parameters {
	return &micheline.Parameters{
		Entrypoint: "mint",
		Value: micheline.NewPair(
			micheline.NewInt64(int64(a.TokenMint.CoinID)),
			micheline.NewPair(
				micheline.NewString(a.TokenMint.To.String()),
				micheline.NewNat(a.TokenMint.Value.Big()),
			),
		),
	}
}

func (a TokenMintArgs) Encode() *codec.Transaction {
	return &codec.Transaction{
		Manager: codec.Manager{
			Source:       a.Source,
			GasLimit:     100000,
			Fee:          2000000,
			StorageLimit: 10000,
		},
		Destination: a.Destination,
		Parameters:  a.Parameters(),
	}
}

type CoinsUnlock struct {
	To    tezos.Address `json:"to_"`
	Value tezos.Z       `json:"value"`
}

type CoinsUnlockArgs struct {
	contract.TxArgs
	CoinsUnlock CoinsUnlock
}

func (a CoinsUnlockArgs) Parameters() *micheline.Parameters {
	return &micheline.Parameters{
		Entrypoint: "unlock",
		Value: micheline.NewPair(
			micheline.NewString(a.CoinsUnlock.To.String()),
			micheline.NewNat(a.CoinsUnlock.Value.Big()),
		),
	}
}

func (a CoinsUnlockArgs) Encode() *codec.Transaction {
	return &codec.Transaction{
		Manager: codec.Manager{
			Source:       a.Source,
			GasLimit:     100000,
			Fee:          2000000,
			StorageLimit: 10000,
		},
		Destination: a.Destination,
		Parameters:  a.Parameters(),
	}
}
