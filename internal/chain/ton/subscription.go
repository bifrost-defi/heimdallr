package ton

import (
	"context"
	"fmt"
	"heimdallr/internal/chain"
	"math/big"
	"time"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"heimdallr/internal/utils"
)

type Subscription struct {
	client   *ton.APIClient
	contract *address.Address

	onTokenBurned chan chain.Event
	onCoinsLocked chan chain.Event
	errs          chan error
}

const (
	checkInterval = 3 * time.Second
	txsBatchSize  = 15

	LockEventID = 101
)

func newSubscription(client *ton.APIClient, contract *address.Address) *Subscription {
	return &Subscription{
		client:        client,
		contract:      contract,
		onTokenBurned: make(chan chain.Event),
		onCoinsLocked: make(chan chain.Event),
		errs:          make(chan error),
	}
}

func (s *Subscription) OnTokenBurned() <-chan chain.Event {
	return s.onTokenBurned
}

func (s *Subscription) OnCoinsLocked() <-chan chain.Event {
	return s.onCoinsLocked
}

func (s *Subscription) Err() <-chan error {
	return s.errs
}

func (s *Subscription) loop(ctx context.Context) {
	tick := time.NewTicker(checkInterval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			s.checkTransactions(ctx)
		}
	}
}

func (s *Subscription) checkTransactions(ctx context.Context) {
	b, err := s.client.CurrentMasterchainInfo(ctx)
	if err != nil {
		s.errs <- fmt.Errorf("get current masterchain info: %w", err)

		return
	}
	acc, err := s.client.GetAccount(ctx, b, s.contract)
	if err != nil {
		s.errs <- fmt.Errorf("get account: %w", err)

		return
	}

	txs, err := s.client.ListTransactions(ctx, s.contract, txsBatchSize, acc.LastTxLT, acc.LastTxHash)
	if err != nil {
		s.errs <- fmt.Errorf("list transactions: %w", err)

		return
	}

	for _, tx := range txs {
		if tx.OutMsgCount == 0 {
			continue
		}

		for _, msg := range tx.IO.Out {
			s.processMessage(msg.AsExternalOut())
		}
	}
}

func (s *Subscription) processMessage(msg *tlb.ExternalMessageOut) {
	eventIdBytes := msg.DestAddr().Data()
	eventId := new(big.Int).SetBytes(eventIdBytes).Uint64()

	if eventId != LockEventID {
		return
	}

	cs := msg.Payload().BeginParse()
	destAddress := utils.BigIntToHex(cs.MustLoadBigUInt(160))
	coinID := cs.MustLoadInt(32)
	fromAddressHash := cs.MustLoadBigUInt(256)
	fromAddress := address.NewAddress(0, byte(msg.SrcAddr.Workchain()), fromAddressHash.Bytes())
	amount := cs.MustLoadBigCoins()

	switch eventId {
	case LockTonEventID:
		s.onCoinsLocked <- chain.NewEvent(
			fromAddress.String(),
			amount,
			int(coinID),
			destAddress,
		)
	case BurnJettonEventID:
		s.onCoinsLocked <- chain.NewEvent(
			fromAddress.String(),
			amount,
			int(coinID),
			destAddress,
		)
	}
}
