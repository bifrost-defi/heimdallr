package avalanche

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Subscription struct {
	onAVAXLocked chan LockEvent
	onUSDCLocked chan LockEvent

	abi  abi.ABI
	errs chan error
}

type LockEvent struct {
	user        common.Address
	amount      *big.Int
	destination common.Address
}

func (e LockEvent) User() string {
	return e.user.Hex()
}

func (e LockEvent) Amount() *big.Int {
	return e.amount
}

func (e LockEvent) Destination() string {
	return e.destination.Hex()
}

var (
	avaxLockedSig = []byte("AVAXLocked(address,uint256,string)")
	usdcLockedSig = []byte("USDCLocked(address,uint256,string)")

	avaxLockedSigHash = crypto.Keccak256Hash(avaxLockedSig)
	usdcLockedSigHash = crypto.Keccak256Hash(usdcLockedSig)
)

func newSubscription(abi abi.ABI) *Subscription {
	return &Subscription{
		onAVAXLocked: make(chan LockEvent),
		onUSDCLocked: make(chan LockEvent),
		abi:          abi,
		errs:         make(chan error),
	}
}

func (s *Subscription) OnAVAXLocked() <-chan LockEvent {
	return s.onAVAXLocked
}

func (s *Subscription) OnUSDCLocked() <-chan LockEvent {
	return s.onUSDCLocked
}

func (s *Subscription) Err() <-chan error {
	return s.errs
}

func (s *Subscription) loop(logs <-chan types.Log) {
	for v := range logs {
		switch v.Topics[0].Hex() {
		case avaxLockedSigHash.Hex():
			var event LockEvent
			if err := s.abi.UnpackIntoInterface(&event, "AVAXLocked", v.Data); err != nil {
				s.errs <- fmt.Errorf("unpack AVAXLocked event: %w", err)
			}
			s.onAVAXLocked <- event
		case usdcLockedSigHash.Hex():
			var event LockEvent
			if err := s.abi.UnpackIntoInterface(&event, "USDCLocked", v.Data); err != nil {
				s.errs <- fmt.Errorf("unpack USDCLocked event: %w", err)
			}
			s.onUSDCLocked <- event
		}
	}
}
