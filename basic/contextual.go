package logging

import (
	"context"
)

// Contextual returns a Logger based on the given Context. Implementations MUST return a non-nil Logger.
type Contextual interface {
	FromContext(ctx context.Context) Logger
}

// FromContextFunc allows a function to implement Contextual
type FromContextFunc func(context.Context) Logger

// FromContext invokes a FromContextFunc
func (f FromContextFunc) FromContext(ctx context.Context) Logger {
	return f(ctx)
}
