package logging

import (
	"context"
)

// Contextual handles operations on context.Context
type Contextual interface {
	// FromContext returns a Logger based on the given Context
	FromContext(ctx context.Context) Logger
}
