package tezos

import (
	"context"
	"fmt"
	"math/big"

	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/rpc"
	"blockwatch.cc/tzgo/tezos"
)

type Tezos struct {
	// WAVAX Token contract.
	wavaxContract *contract.Contract
	// WUSDC Token contract.
	wusdcContract *contract.Contract

	privateKey string

	client *rpc.Client
}

func New(client *rpc.Client, privateKey string) *Tezos {
	return &Tezos{
		privateKey: privateKey,
		client:     client,
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

// Subscribe starts listening to events and returns Subscription.
func (t *Tezos) Subscribe(ctx context.Context) (*Subscription, error) {
	wavaxStorage := newStorage(t.getStorageLoader(t.wavaxContract))
	wusdcStorage := newStorage(t.getStorageLoader(t.wusdcContract))

	s := newSubscription(wavaxStorage, wusdcStorage)
	go s.loop(ctx)

	return s, nil
}

func (t *Tezos) MintWUSDC(ctx context.Context, user string, amount *big.Int) (string, *big.Int, error) {
	userAddr, err := tezos.ParseAddress(user)
	if err != nil {
		return "", nil, fmt.Errorf("parse user address")
	}

	pk, err := tezos.ParsePrivateKey(t.privateKey)
	if err != nil {
		return "", nil, fmt.Errorf("parse private key: %w", err)
	}

	wusdc := TokenMint{
		User:   userAddr,
		Amount: tezos.Z(*amount),
	}
	opts := &contract.CallOptions{
		Signer: newSigner(pk),
	}

	tx, err := t.wusdcContract.Call(ctx, &TokenMintArgs{Mint: wusdc}, opts)
	if err != nil {
		return "", nil, fmt.Errorf("call contract: %w", err)
	}

	return tx.Op.Hash.String(), big.NewInt(tx.Costs()[0].Fee), nil
}

func (t *Tezos) MintWAVAX(ctx context.Context, user string, amount *big.Int) (string, *big.Int, error) {
	userAddr, err := tezos.ParseAddress(user)
	if err != nil {
		return "", nil, fmt.Errorf("parse user address")
	}

	pk, err := tezos.ParsePrivateKey(t.privateKey)
	if err != nil {
		return "", nil, fmt.Errorf("parse private key: %w", err)
	}

	wavax := TokenMint{
		User:   userAddr,
		Amount: tezos.Z(*amount),
	}
	opts := &contract.CallOptions{
		Signer: newSigner(pk),
	}

	tx, err := t.wavaxContract.Call(ctx, &TokenMintArgs{Mint: wavax}, opts)
	if err != nil {
		return "", nil, fmt.Errorf("call contract: %w", err)
	}

	return tx.Op.Hash.String(), big.NewInt(tx.Costs()[0].Fee), nil
}

func (t *Tezos) getStorageLoader(contract *contract.Contract) storageLoader {
	return func(ctx context.Context) (map[string]interface{}, error) {
		script, err := t.client.GetContractScript(ctx, contract.Address())
		if err != nil {
			return nil, fmt.Errorf("get contract script: %w", err)
		}

		val := micheline.NewValue(script.StorageType(), script.Storage)

		m, err := val.Map()
		if err != nil {
			return nil, fmt.Errorf("get map: %w", err)
		}

		s, ok := m.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid map")
		}

		return s, nil
	}
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
