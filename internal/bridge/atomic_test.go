package bridge

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"heimdallr/internal/chain"
)

type AtomicSuite struct {
	suite.Suite
	atomic *Atomic
}

func (s *AtomicSuite) SetupTest() {
	s.atomic = NewAtomic()
}

func (s *AtomicSuite) TestFailedRollback() {
	f := func(_ context.Context, _ chain.Event) bool {
		return false
	}
	rf := func(_ context.Context, _ chain.Event) bool {
		return false
	}

	op := s.atomic.NewOperation(
		OnPerform(f),
		OnRollback(rf),
	)
	go op.Run(nil, chain.Event{})

	require.Equal(s.T(), true, errors.Is(<-op.Fail(), ErrRollbackFailed))
}

func TestAtomic(t *testing.T) {
	suite.Run(t, new(AtomicSuite))
}
