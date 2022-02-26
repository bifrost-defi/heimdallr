package tezos

import (
	"context"

	"blockwatch.cc/tzgo/rpc"
)

type Tezos struct {
	client *rpc.Client
}

func New(client *rpc.Client) *Tezos {
	return &Tezos{
		client: client,
	}
}

func (t *Tezos) Subscribe(ctx context.Context, contract string) (*Subscription, error) {
	return &Subscription{}, nil
}

func (t *Tezos) MintWUSDC() error {
	panic("implement me!")
}

func (t *Tezos) MintWAVAX() error {
	panic("implement me!")
}
