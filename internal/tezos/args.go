package tezos

import (
	"blockwatch.cc/tzgo/codec"
	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

type TokenMint struct {
	To     tezos.Address `json:"to_"`
	CoinID int           `json:"coinId"`
	Value  tezos.Z       `json:"value"`
}

type TokenMintArgs struct {
	contract.TxArgs
	Mint TokenMint
}

func (a TokenMintArgs) Parameters() *micheline.Parameters {
	return &micheline.Parameters{
		Entrypoint: "mint",
		Value: micheline.NewPair(
			micheline.NewString(a.Mint.To.String()),
			micheline.NewPair(
				micheline.NewInt64(int64(a.Mint.CoinID)),
				micheline.NewNat(a.Mint.Value.Big()),
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
