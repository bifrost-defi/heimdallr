package tezos

import (
	"blockwatch.cc/tzgo/codec"
	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

type TokenMint struct {
	User   tezos.Address `json:"user"`
	Amount tezos.Z       `json:"amount"`
}

type TokenMintArgs struct {
	contract.TxArgs
	Mint TokenMint
}

func (a TokenMintArgs) Parameters() *micheline.Parameters {
	return &micheline.Parameters{
		Entrypoint: "mintAndTransfer",
		Value: micheline.NewPair(
			micheline.NewBytes(a.Mint.User.Bytes22()),
			micheline.NewNat(a.Mint.Amount.Big()),
		),
	}
}

func (a TokenMintArgs) Encode() *codec.Transaction {
	return &codec.Transaction{
		Manager: codec.Manager{
			Source: a.Source,
		},
		Destination: a.Destination,
		Parameters:  a.Parameters(),
	}
}
