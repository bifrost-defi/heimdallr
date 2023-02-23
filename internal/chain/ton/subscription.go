package ton

import (
	"context"
	"fmt"
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
	onLocked chan LockEvent
	errs     chan error
}

const (
	checkInterval = 3 * time.Second
	txsBatchSize  = 15

	LockEventID = 101
)

type LockEvent struct {
	user        *address.Address
	amount      *big.Int
	coinId      int
	destination string
}

func (e LockEvent) User() string {
	return e.user.String()
}

func (e LockEvent) Amount() *big.Int {
	return e.amount
}

func (e LockEvent) CoinID() int {
	return e.coinId
}

func (e LockEvent) Destination() string {
	return e.destination
}

func newSubscription(client *ton.APIClient, contract *address.Address) *Subscription {
	return &Subscription{
		client:   client,
		contract: contract,
		onLocked: make(chan LockEvent),
		errs:     make(chan error),
	}
}

func (s *Subscription) OnLocked() <-chan LockEvent {
	return s.onLocked
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
	destCoinID := cs.MustLoadInt(32)
	fromAddressHash := cs.MustLoadBigUInt(256)
	fromAddress := address.NewAddress(0, byte(msg.SrcAddr.Workchain()), fromAddressHash.Bytes())
	amount := cs.MustLoadBigCoins()

	s.onLocked <- LockEvent{
		user:        fromAddress,
		amount:      amount,
		coinId:      int(destCoinID),
		destination: destAddress,
	}
}
