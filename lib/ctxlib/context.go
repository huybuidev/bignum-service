package ctxlib

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Context contains all context information in a request
type Context struct {
	context.Context
	Logger *zerolog.Logger
}

// Background returns a non-nil, empty Context. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context {
	defaultLogger := log.Logger.With().Logger()
	return Context{
		Context: context.Background(),
		Logger:  &defaultLogger,
	}
}

// WithCancel returns a copy of parent with a new Done channel. The returned
// context's Done channel is closed when the returned cancel function is called
// or when the parent context's Done channel is closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func WithCancel(parent Context) (ctx Context, cancel context.CancelFunc) {
	newCtx, cFunc := context.WithCancel(parent)
	return Context{
		Context: newCtx,
		Logger:  parent.Logger,
	}, cFunc
}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete:
//
//	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
//		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
//		defer cancel()  // releases resources if slowOperation completes before timeout elapses
//		return slowOperation(ctx)
//	}
func WithTimeout(parent Context, timeout time.Duration) (ctx Context, cancel context.CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

// WithDeadline returns a copy of the parent context with the deadline adjusted
// to be no later than d. If the parent's deadline is already earlier than d,
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
// context's Done channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is
// closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func WithDeadline(parent Context, d time.Time) (Context, context.CancelFunc) {
	newCtx, cFunc := context.WithDeadline(parent, d)
	return Context{
		Context: newCtx,
		Logger:  parent.Logger,
	}, cFunc
}

// WithLogger returns background context with providing logger
func WithLogger(logger *zerolog.Logger) Context {
	return Context{
		Context: context.Background(),
		Logger:  logger,
	}
}
