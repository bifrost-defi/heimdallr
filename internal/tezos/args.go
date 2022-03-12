package tezos

import (
	"blockwatch.cc/tzgo/codec"
	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

type TokenMint struct {
	Value tezos.Z `json:"value"`
}

type TokenMintArgs struct {
	contract.TxArgs
	Mint TokenMint
}

func (a TokenMintArgs) Parameters() *micheline.Parameters {
	return &micheline.Parameters{
		Entrypoint: "mint",
		Value:      micheline.NewNat(a.Mint.Value.Big()),
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
