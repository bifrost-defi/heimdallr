package tezos

import "blockwatch.cc/tzgo/rpc"

type Tezos struct {
	client *rpc.Client
}

func New(client *rpc.Client) *Tezos {
	return &Tezos{
		client: client,
	}
}

func (t *Tezos) MintUSDC() error {
	panic("implement me!")
}

func (t *Tezos) MintWAVAX() error {
	panic("implement me!")
}
