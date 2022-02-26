package tezos

import (
	"context"
	"fmt"
	"math/big"

	"blockwatch.cc/tzgo/rpc"
	"blockwatch.cc/tzgo/tezos"
)

type Tezos struct {
	// WAVAX Token contract address
	wavaxContract tezos.Address
	// WUSDC Token contract address
	wusdcContract tezos.Address

	client *rpc.Client
}

func New(client *rpc.Client) *Tezos {
	return &Tezos{
		client: client,
	}
}

func (t *Tezos) SetContracts(wavaxContractAddr string, wusdcContractAddr string) error {
	wavax, err := tezos.ParseAddress(wavaxContractAddr)
	if err != nil {
		return fmt.Errorf("parse wavax address: %w", err)
	}
	t.wavaxContract = wavax

	wusdc, err := tezos.ParseAddress(wusdcContractAddr)
	if err != nil {
		return fmt.Errorf("parse wusdc address: %w", err)
	}
	t.wusdcContract = wusdc

	return nil
}

func (t *Tezos) Subscribe(ctx context.Context) (*Subscription, error) {
	return &Subscription{}, nil
}

func (t *Tezos) MintWUSDC(user string, amount *big.Int) (string, *big.Int, error) {
	panic("implement me!")
}

func (t *Tezos) MintWAVAX(user string, amount *big.Int) (string, *big.Int, error) {
	panic("implement me!")
}
