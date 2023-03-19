package tezos

import (
	"context"
	"fmt"
	"math/big"

	"blockwatch.cc/tzgo/contract"
	"blockwatch.cc/tzgo/rpc"
	"blockwatch.cc/tzgo/signer"
	"blockwatch.cc/tzgo/tezos"
	"heimdallr/internal/chain"
)

type Tezos struct {
	bridgeContract *contract.Contract

	privateKey string
	client     *rpc.Client
}

var _ chain.Chain = (*Tezos)(nil)

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
func (t *Tezos) Subscribe(ctx context.Context) (chain.Subscription, error) {
	s := newSubscription(t.bridgeContract)
	go s.loop(ctx)

	return s, nil
}

func (t *Tezos) MintToken(ctx context.Context, to string, coinId int, amount *big.Int) (string, *big.Int, error) {
	pk, err := tezos.ParsePrivateKey(t.privateKey)
	if err != nil {
		return "", nil, fmt.Errorf("parse private key: %w", err)
	}

	address, err := tezos.ParseAddress(to)
	if err != nil {
		return "", nil, fmt.Errorf("parse address: %w", err)
	}

	tm := TokenMint{
		CoinID: coinId,
		To:     address,
		Value:  tezos.Z(*amount),
	}
	opts := &rpc.CallOptions{
		Confirmations: confirmations,
		TTL:           120,
		Signer:        signer.NewFromKey(pk),
	}
	args := &TokenMintArgs{TokenMint: tm}

	tx, err := t.bridgeContract.Call(ctx, args, opts)
	if err != nil {
		return "", nil, fmt.Errorf("call contract: %w", err)
	}

	return tx.Op.Hash.String(), big.NewInt(tx.Costs()[0].Fee), nil
}

func (t *Tezos) UnlockCoins(ctx context.Context, user string, amount *big.Int) (string, *big.Int, error) {
	pk, err := tezos.ParsePrivateKey(t.privateKey)
	if err != nil {
		return "", nil, fmt.Errorf("parse private key: %w", err)
	}

	address, err := tezos.ParseAddress(user)
	if err != nil {
		return "", nil, fmt.Errorf("parse destination address: %w", err)
	}

	cu := CoinsUnlock{
		To:    address,
		Value: tezos.Z(*amount),
	}
	opts := &rpc.CallOptions{
		Confirmations: confirmations,
		TTL:           120,
		Signer:        signer.NewFromKey(pk),
	}
	args := &CoinsUnlockArgs{CoinsUnlock: cu}

	tx, err := t.bridgeContract.Call(ctx, args, opts)
	if err != nil {
		return "", nil, fmt.Errorf("call contract: %w", err)
	}

	return tx.Op.Hash.String(), big.NewInt(tx.Costs()[0].Fee), nil
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
