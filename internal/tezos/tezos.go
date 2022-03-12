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

const confirmations = 5

func New(client *rpc.Client, privateKey string) *Tezos {
	return &Tezos{
		privateKey: privateKey,
		client:     client,
	}
}

func (t *Tezos) LoadContracts(ctx context.Context, wavaxContractAddr string, wusdcContractAddr string) error {
	if err := t.client.Init(ctx); err != nil {
		return fmt.Errorf("init client id: %w", err)
	}
	t.client.Listen()

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
	wavaxStorage := newStorage(t.getBigMapLoader(t.wavaxContract))
	wusdcStorage := newStorage(t.getBigMapLoader(t.wusdcContract))

	s := newSubscription(wavaxStorage, wusdcStorage)
	go s.loop(ctx)

	return s, nil
}

func (t *Tezos) MintWUSDC(ctx context.Context, amount *big.Int) (string, *big.Int, error) {
	pk, err := tezos.ParsePrivateKey(t.privateKey)
	if err != nil {
		return "", nil, fmt.Errorf("parse private key: %w", err)
	}

	mint := TokenMint{
		Value: tezos.Z(*amount),
	}
	opts := &contract.CallOptions{
		Confirmations: confirmations,
		TTL:           120,
		Signer:        newSigner(pk),
	}
	args := &TokenMintArgs{Mint: mint}

	tx, err := t.wusdcContract.Call(ctx, args, opts)
	if err != nil {
		return "", nil, fmt.Errorf("call contract: %w", err)
	}

	return tx.Op.Hash.String(), big.NewInt(tx.Costs()[0].Fee), nil
}

func (t *Tezos) TransferWUSDC(ctx context.Context, user string, amount *big.Int) (string, *big.Int, error) {
	userAddr, err := tezos.ParseAddress(user)
	if err != nil {
		return "", nil, fmt.Errorf("parse user address: %w", err)
	}

	pk, err := tezos.ParsePrivateKey(t.privateKey)
	if err != nil {
		return "", nil, fmt.Errorf("parse private key: %w", err)
	}

	transfer := contract.FA1Transfer{
		From:   pk.Address(),
		To:     userAddr,
		Amount: tezos.Z(*amount),
	}
	opts := &contract.CallOptions{
		Confirmations: confirmations,
		TTL:           120,
		Signer:        newSigner(pk),
	}

	tx, err := t.wusdcContract.Call(ctx, &contract.FA1TransferArgs{Transfer: transfer}, opts)
	if err != nil {
		return "", nil, fmt.Errorf("call contract: %w", err)
	}

	return tx.Op.Hash.String(), big.NewInt(tx.Costs()[0].Fee), nil
}

func (t *Tezos) MintWAVAX(ctx context.Context, amount *big.Int) (string, *big.Int, error) {
	pk, err := tezos.ParsePrivateKey(t.privateKey)
	if err != nil {
		return "", nil, fmt.Errorf("parse private key: %w", err)
	}

	mint := TokenMint{
		Value: tezos.Z(*amount),
	}
	opts := &contract.CallOptions{
		Confirmations: confirmations,
		TTL:           120,
		Signer:        newSigner(pk),
	}

	tx, err := t.wavaxContract.Call(ctx, &TokenMintArgs{Mint: mint}, opts)
	if err != nil {
		return "", nil, fmt.Errorf("call contract: %w", err)
	}

	return tx.Op.Hash.String(), big.NewInt(tx.Costs()[0].Fee), nil
}

func (t *Tezos) TransferWAVAX(ctx context.Context, user string, amount *big.Int) (string, *big.Int, error) {
	userAddr, err := tezos.ParseAddress(user)
	if err != nil {
		return "", nil, fmt.Errorf("parse user address: %w", err)
	}

	pk, err := tezos.ParsePrivateKey(t.privateKey)
	if err != nil {
		return "", nil, fmt.Errorf("parse private key: %w", err)
	}

	transfer := contract.FA1Transfer{
		From:   pk.Address(),
		To:     userAddr,
		Amount: tezos.Z(*amount),
	}
	opts := &contract.CallOptions{
		Confirmations: confirmations,
		TTL:           120,
		Signer:        newSigner(pk),
	}

	tx, err := t.wavaxContract.Call(ctx, &contract.FA1TransferArgs{Transfer: transfer}, opts)
	if err != nil {
		return "", nil, fmt.Errorf("call contract: %w", err)
	}

	return tx.Op.Hash.String(), big.NewInt(tx.Costs()[0].Fee), nil
}

func (t *Tezos) getBigMapLoader(contract *contract.Contract) bigmapLoader {
	return func(ctx context.Context, name string) (map[string]interface{}, error) {
		script := contract.Script()
		storage, err := t.client.GetContractStorage(ctx, contract.Address(), rpc.Head)
		if err != nil {
			return nil, fmt.Errorf("get contract storage: %w", err)
		}

		val := micheline.NewValue(script.StorageType(), storage)

		m, err := val.Map()
		if err != nil {
			return nil, fmt.Errorf("get map: %w", err)
		}

		s, ok := m.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid map")
		}

		burningsID, ok := s[name]
		if !ok {
			return nil, fmt.Errorf("burnings bigmap not found")
		}

		info, err := t.client.GetBigmapInfo(ctx, burningsID.(int64), rpc.Head)
		if err != nil {
			return nil, fmt.Errorf("get bigmap info: %w", err)
		}

		keys, err := t.client.ListBigmapKeys(ctx, burningsID.(int64), rpc.Head)
		if err != nil {
			return nil, fmt.Errorf("list bigmap values: %w", err)
		}

		bigmap := make(map[string]interface{}, len(keys))
		for _, k := range keys {
			v, err := t.client.GetBigmapValue(ctx, burningsID.(int64), k, rpc.Head)
			if err != nil {
				return nil, fmt.Errorf("get bigmap value: %w", err)
			}

			val := micheline.NewValue(micheline.NewType(info.ValueType), v)
			m, err := val.Map()
			if err != nil {
				return nil, fmt.Errorf("map: %w", err)
			}

			bigmap[k.String()] = m
		}

		return bigmap, nil
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
