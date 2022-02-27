package tezos

import (
	"context"
	"fmt"
	"math/big"

	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/rpc"
	"blockwatch.cc/tzgo/tezos"
)

type Tezos struct {
	// WAVAX Token contract address
	wavaxContract *contract.Contract
	// WUSDC Token contract address
	wusdcContract *contract.Contract

	client *rpc.Client
}

func New(client *rpc.Client) *Tezos {
	return &Tezos{
		client: client,
	}
}

func (t *Tezos) LoadContracts(ctx context.Context, wavaxContractAddr string, wusdcContractAddr string) error {
	wavax, err := t.loadContract(ctx, wavaxContractAddr, true)
	if err != nil {
		return fmt.Errorf("load wavax contract: %w", err)
	}
	t.wavaxContract = wavax

	wusdc, err := t.loadContract(ctx, wusdcContractAddr, true)
	if err != nil {
		return fmt.Errorf("load wusdc contract: %w", err)
	}
	t.wusdcContract = wusdc

	return nil
}

func (t *Tezos) Subscribe(ctx context.Context) (*Subscription, error) {
	return &Subscription{}, nil
}

func (t *Tezos) MintWUSDC(ctx context.Context, user string, amount *big.Int) (string, *big.Int, error) {
	panic("implement me!")
}

func (t *Tezos) MintWAVAX(ctx context.Context, user string, amount *big.Int) (string, *big.Int, error) {
	panic("implement me!")
}

func (t *Tezos) loadContract(ctx context.Context, addr string, resolve bool) (*contract.Contract, error) {
	a, err := tezos.ParseAddress(addr)
	if err != nil {
		return nil, err
	}
	if a.Type != tezos.AddressTypeContract {
		return nil, fmt.Errorf("invalid contract address")
	}
	c := contract.NewContract(a, t.client)

	if resolve {
		if err := c.Resolve(ctx); err != nil {
			return nil, err
		}
	}

	return c, nil
}
