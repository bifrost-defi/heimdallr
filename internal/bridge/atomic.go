package bridge

import (
	"context"
	"errors"
)

// Atomic represents wrapper for functions call
// with rolling back if it is required after call.
type Atomic struct {
	errs chan error
}

var ErrRollbackFailed = errors.New("rollback failed")

func NewAtomic() *Atomic {
	return &Atomic{errs: make(chan error)}
}

func (a *Atomic) Errs() <-chan error {
	return a.errs
}

type Operation struct {
	perform  Fn
	rollback Fn

	errs chan<- error
}

type Fn func(ctx context.Context, event Event) (ok bool)

type Option interface {
	apply(*Operation)
}

type optionFn func(*Operation)

func (f optionFn) apply(o *Operation) {
	f(o)
}

func OnPerform(fn Fn) Option {
	return optionFn(func(a *Operation) {
		a.perform = fn
	})
}

func OnRollback(fn Fn) Option {
	return optionFn(func(a *Operation) {
		a.rollback = fn
	})
}

func (a *Atomic) NewOperation(options ...Option) *Operation {
	o := new(Operation)
	o.errs = a.errs

	for _, op := range options {
		op.apply(o)
	}

	return o
}

func (a *Operation) Run(ctx context.Context, event Event) {
	if a.perform == nil {
		return
	}

	if ok := a.perform(ctx, event); ok {
		// Everything fine
		return
	}

	if a.rollback != nil {
		if ok := a.rollback(ctx, event); !ok {
			// Nothing fine
			a.errs <- ErrRollbackFailed
		}
	}
}
