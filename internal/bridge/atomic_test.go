package bridge

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
	op.Run(nil, nil)

	require.Equal(s.T(), <-op.Fail(), ErrRollbackFailed)
}

func TestAtomic(t *testing.T) {
	suite.Run(t, new(AtomicSuite))
}
