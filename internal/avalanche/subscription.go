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
	onAVAXLocked chan LockedEvent
	onUSDCLocked chan LockedEvent

	abi  abi.ABI
	errs chan error
}

type LockedEvent struct {
	user        common.Address
	amount      *big.Int
	destination common.Address
}

var (
	avaxLockedSig = []byte("AVAXLocked(address,uint256,string)")
	usdcLockedSig = []byte("USDCLocked(address,uint256,string)")

	avaxLockedSigHash = crypto.Keccak256Hash(avaxLockedSig)
	usdcLockedSigHash = crypto.Keccak256Hash(usdcLockedSig)
)

func newSubscription(abi abi.ABI) *Subscription {
	return &Subscription{
		onAVAXLocked: make(chan LockedEvent),
		onUSDCLocked: make(chan LockedEvent),
		abi:          abi,
		errs:         make(chan error),
	}
}

func (s *Subscription) OnAVAXLocked() <-chan LockedEvent {
	return s.onAVAXLocked
}

func (s *Subscription) OnUSDCLocked() <-chan LockedEvent {
	return s.onUSDCLocked
}

func (s *Subscription) Err() <-chan error {
	return s.errs
}

func (s *Subscription) loop(logs <-chan types.Log) {
	for v := range logs {
		switch v.Topics[0].Hex() {
		case avaxLockedSigHash.Hex():
			var event LockedEvent
			if err := s.abi.UnpackIntoInterface(&event, "AVAXLocked", v.Data); err != nil {
				s.errs <- fmt.Errorf("unpack AVAXLocked event: %w", err)
			}
			s.onAVAXLocked <- event
		case usdcLockedSigHash.Hex():
			var event LockedEvent
			if err := s.abi.UnpackIntoInterface(&event, "USDCLocked", v.Data); err != nil {
				s.errs <- fmt.Errorf("unpack USDCLocked event: %w", err)
			}
			s.onUSDCLocked <- event
		}
	}
}
