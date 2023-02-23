package ton

import (
	"context"
	"fmt"
	"heimdallr/internal/chain"
	"math/big"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

type TON struct {
	client   *ton.APIClient
	wallet   *wallet.Wallet
	contract *address.Address
}

var _ chain.Chain = (*TON)(nil)

func New(client *ton.APIClient, wallet *wallet.Wallet, contractAddr string) *TON {
	return &TON{
		client:   client,
		wallet:   wallet,
		contract: address.MustParseAddr(contractAddr),
	}
}

func (t *TON) Subscribe(ctx context.Context) (*Subscription, error) {
	s := newSubscription(t.client, t.contract)
	go s.loop(ctx)

	return s, nil
}

func (t *TON) UnlockCoins(ctx context.Context, account string, amount *big.Int) (string, *big.Int, error) {
	destination, err := address.ParseAddr(account)
	if err != nil {
		return "", nil, fmt.Errorf("parse destination: %w", err)
	}

	body := cell.BeginCell().
		MustStoreUInt(UnlockOpCode, 32). // op code
		MustStoreUInt(0, 64).            // query id
		MustStoreAddr(destination).
		MustStoreCoins(amount.Uint64()).
		EndCell()

	err = t.wallet.Send(ctx, &wallet.Message{
		Mode: 1, // pay fees separately (from balance, not from amount)
		InternalMessage: &tlb.InternalMessage{
			DstAddr: t.contract,
			Amount:  tlb.MustFromTON("0.03"),
			Body:    body,
		},
	}, true)
	if err != nil {
		return "", nil, fmt.Errorf("send message: %w", err)
	}

	// TODO: calculate hash and fee
	return "", big.NewInt(0), nil
}

func (t *TON) MintToken(ctx context.Context, destination string, coinId int, amount *big.Int) (string, *big.Int, error) {
	//TODO implement me
	panic("implement me")
}
