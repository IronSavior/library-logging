package logging

import "context"

// Contextual handles context operations for retrieving Loggers from context and deriving contexts with log fields set
type Contextual interface {
	FromContexter
	ContextFieldser
}

// FromContexter returns a Logger based on the given Context. Implementations MUST return a non-nil Logger.
type FromContexter interface {
	FromContext(ctx context.Context) Logger
}

// ContextFieldser derives a new Context containing a new Logger with the given structured fields. Implementations SHOULD
// inherit structured fields from an existing Logger (if a Logger is found in the context). The structured fields of
// existing Loggers already in the Context may be superseded by the new Logger in the new Context, but any existing
// Loggers MUST NOT be altered.
type ContextFieldser interface {
	ContextFields(ctx context.Context, fieldser Fieldser) context.Context
}

// ContextFuncs combines two functions into one object implementing Contextual
type ContextFuncs struct {
	FromContextFunc
	ContextFieldsFunc
}

// FromContextFunc allows a function to implement FromContexter
type FromContextFunc func(context.Context) Logger

// FromContext invokes a FromContextFn
func (f FromContextFunc) FromContext(ctx context.Context) Logger {
	return f(ctx)
}

// ContextFieldsFunc allows a function to implement ContextFieldser
type ContextFieldsFunc func(context.Context, Fieldser) context.Context

// ContextFields invokes a ContextFieldsFunc
func (f ContextFieldsFunc) ContextFields(ctx context.Context, fields Fieldser) context.Context {
	return f(ctx, fields)
}
