package bridge

import (
	"context"
	"errors"
	"fmt"

	"heimdallr/internal/chain"
)

// Atomic represents wrapper for functions call
// with rolling back if it is required after call.
type Atomic struct {
	options []Option
}

var (
	ErrNothingToPerform = errors.New("nothing to perform")
	ErrRollbackFailed   = errors.New("rollback failed")
)

// NewAtomic returns Atomic instance with options that will be applied
// for every operation created from this instance.
func NewAtomic(options ...Option) *Atomic {
	return &Atomic{options: options}
}

type Operation struct {
	name string

	performFn  Fn
	rollbackFn Fn
	checkerFn  CheckerFn

	completeCh chan struct{}
	rollbackCh chan struct{}
	failCh     chan error
}

type (
	Fn func(ctx context.Context, event chain.Event) (ok bool)

	// Checker exposes methods for operation state checking.
	Checker interface {
		Complete() <-chan struct{}
		Rollback() <-chan struct{}
		Fail() <-chan error
	}
	CheckerFn func(op Checker, event chain.Event)
)

type Option interface {
	apply(*Operation)
}

type optionFn func(*Operation)

func (f optionFn) apply(o *Operation) {
	f(o)
}

// WithName sets the name of an operation.
func WithName(s string) Option {
	return optionFn(func(o *Operation) {
		o.name = s
	})
}

// OnPerform sets function than should be run by operation.
func OnPerform(fn Fn) Option {
	return optionFn(func(o *Operation) {
		o.performFn = fn
	})
}

// OnRollback sets rollback function for operation.
func OnRollback(fn Fn) Option {
	return optionFn(func(o *Operation) {
		o.rollbackFn = fn
	})
}

// WithChecker sets function which takes Checker and Event to
// access operation state.
func WithChecker(fn CheckerFn) Option {
	return optionFn(func(o *Operation) {
		o.checkerFn = fn
	})
}

// NewOperation creates operation from options.
// Global options will be overwritten on collision.
func (a *Atomic) NewOperation(options ...Option) *Operation {
	o := new(Operation)
	o.completeCh = make(chan struct{})
	o.rollbackCh = make(chan struct{})
	o.failCh = make(chan error)

	// Apply global options first.
	for _, op := range a.options {
		op.apply(o)
	}
	// Apply operation options. Global options will be overwritten.
	for _, op := range options {
		op.apply(o)
	}

	return o
}

// Run runs operation and controls its depending on options.
func (o *Operation) Run(ctx context.Context, event chain.Event) {
	if o.checkerFn != nil {
		go o.checkerFn(o, event)
	}

	if o.performFn == nil {
		o.failCh <- fmt.Errorf("%s: %w", o.name, ErrNothingToPerform)

		return
	}

	if ok := o.performFn(ctx, event); ok {
		// Everything fine
		close(o.completeCh)

		return
	}

	if o.rollbackFn != nil {
		if ok := o.rollbackFn(ctx, rollbackEvent(event)); !ok {
			// Nothing fine
			// TODO?: try to repeat operation on failure
			o.failCh <- fmt.Errorf("%s: %w", o.name, ErrRollbackFailed)

			return
		}
		close(o.rollbackCh)
	}
}

// Complete returns chan which is closed if operation completed successfully.
func (o *Operation) Complete() <-chan struct{} {
	return o.completeCh
}

// Rollback returns chan which is closed if operation rolled back.
func (o *Operation) Rollback() <-chan struct{} {
	return o.rollbackCh
}

// Fail returns error if operation failed.
func (o *Operation) Fail() <-chan error {
	return o.failCh
}
