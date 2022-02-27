package tezos

import (
	"context"

	"blockwatch.cc/tzgo/codec"
	"blockwatch.cc/tzgo/tezos"
)

type signer struct {
	privateKey tezos.PrivateKey
}

func newSigner(pk tezos.PrivateKey) *signer {
	return &signer{
		privateKey: pk,
	}
}

func (s *signer) Key(_ context.Context) (tezos.Key, error) {
	return s.privateKey.Public(), nil
}
func (s *signer) SignOperation(_ context.Context, op *codec.Op) (tezos.Signature, error) {
	if err := op.Sign(s.privateKey); err != nil {
		return tezos.Signature{}, err
	}

	return op.Signature, nil
}

// Address has no usages.
func (s *signer) Address(_ context.Context) (tezos.Address, error) {
	panic("not implemented")
}

// SignMessage has no usages.
func (s *signer) SignMessage(context.Context, string) (tezos.Signature, error) {
	panic("not implemented")
}

// SignBlock has no usages.
func (s *signer) SignBlock(context.Context, *codec.BlockHeader) (tezos.Signature, error) {
	panic("not implemented")
}
