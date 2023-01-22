package tezos

import (
	"context"
	"fmt"
	"math/big"

	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/rpc"
	"blockwatch.cc/tzgo/signer"
	"blockwatch.cc/tzgo/tezos"
)

type Tezos struct {
	bridgeContract *contract.Contract

	privateKey string
	client     *rpc.Client
}

const confirmations = 5

func New(client *rpc.Client, privateKey string) *Tezos {
	return &Tezos{
		privateKey: privateKey,
		client:     client,
	}
}

func (t *Tezos) LoadContracts(ctx context.Context, bridgeContractAddr string) error {
	if err := t.client.Init(ctx); err != nil {
		return fmt.Errorf("init client id: %w", err)
	}
	t.client.Listen()

	bridge, err := t.loadContract(ctx, bridgeContractAddr, true)
	if err != nil {
		return fmt.Errorf("load contract: %w", err)
	}
	t.bridgeContract = bridge

	return nil
}

// Subscribe starts listening to events and returns Subscription.
func (t *Tezos) Subscribe(ctx context.Context) (*Subscription, error) {
	bridgeStorage := newStorage(t.getBigMapLoader(t.bridgeContract))

	s := newSubscription(bridgeStorage)
	go s.loop(ctx)

	return s, nil
}

func (t *Tezos) MintToken(ctx context.Context, destination string, coinId int, amount *big.Int) (string, *big.Int, error) {
	pk, err := tezos.ParsePrivateKey(t.privateKey)
	if err != nil {
		return "", nil, fmt.Errorf("parse private key: %w", err)
	}

	address, err := tezos.ParseAddress(destination)
	if err != nil {
		return "", nil, fmt.Errorf("parse destination address: %w", err)
	}

	mint := TokenMint{
		To:     address,
		CoinID: coinId,
		Value:  tezos.Z(*amount),
	}
	opts := &rpc.CallOptions{
		Confirmations: confirmations,
		TTL:           120,
		Signer:        signer.NewFromKey(pk),
	}
	args := &TokenMintArgs{Mint: mint}

	tx, err := t.bridgeContract.Call(ctx, args, opts)
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
