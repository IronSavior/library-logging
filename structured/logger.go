package logging

import (
	basic "github.com/IronSavior/library-logging/basic"
)

// Logger writes printf-style messages at standard log levels with structured fields
type Logger interface {
	basic.Logger

	// WithFields derives a new Logger with the current and given fields. The original Logger MUST NOT be altered.
	WithFields(fields Fieldser) Logger
}
