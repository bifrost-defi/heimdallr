package bridge

import (
	"context"
	"errors"
	"fmt"
)

// Atomic represents wrapper for functions call
// with rolling back if it is required after call.
type Atomic struct {
	options []Option
	errs    chan error
}

var (
	ErrNothingToPerform = errors.New("nothing to perform")
	ErrRollbackFailed   = errors.New("rollback failed")
)

// NewAtomic return Atomic instance with options that will be applied
// for every operation created from this instance.
func NewAtomic(options ...Option) *Atomic {
	return &Atomic{options: options, errs: make(chan error)}
}

func (a *Atomic) Errs() <-chan error {
	return a.errs
}

type Operation struct {
	name string

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

// WithName sets the name of an operation.
func WithName(s string) Option {
	return optionFn(func(o *Operation) {
		o.name = s
	})
}

// OnPerform sets function than should be run by operation.
func OnPerform(fn Fn) Option {
	return optionFn(func(o *Operation) {
		o.perform = fn
	})
}

// OnRollback sets rollback function for operation.
func OnRollback(fn Fn) Option {
	return optionFn(func(o *Operation) {
		o.rollback = fn
	})
}

// NewOperation creates operation from options.
// Global options will be overwritten on collision.
func (a *Atomic) NewOperation(options ...Option) *Operation {
	o := new(Operation)
	o.errs = a.errs

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
func (o *Operation) Run(ctx context.Context, event Event) {
	if o.perform == nil {
		o.errs <- fmt.Errorf("%s: %w", o.name, ErrNothingToPerform)

		return
	}

	if ok := o.perform(ctx, event); ok {
		// Everything fine
		return
	}

	if o.rollback != nil {
		if ok := o.rollback(ctx, event); !ok {
			// Nothing fine
			o.errs <- fmt.Errorf("%s: %w", o.name, ErrRollbackFailed)
		}
	}
}
