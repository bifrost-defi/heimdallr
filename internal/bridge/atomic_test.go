package bridge

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"heimdallr/internal/evm"
)

type AtomicSuite struct {
	suite.Suite
	atomic *Atomic
}

func (s *AtomicSuite) SetupTest() {
	s.atomic = NewAtomic()
}

func (s *AtomicSuite) TestFailedRollback() {
	f := func(_ context.Context, _ Event) bool {
		return false
	}
	rf := func(_ context.Context, _ Event) bool {
		return false
	}

	op := s.atomic.NewOperation(
		OnPerform(f),
		OnRollback(rf),
	)
	go op.Run(nil, evm.LockEvent{})

	require.Equal(s.T(), true, errors.Is(<-op.Fail(), ErrRollbackFailed))
}

func TestAtomic(t *testing.T) {
	suite.Run(t, new(AtomicSuite))
}
