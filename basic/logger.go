package logging

// Logger writes printf-style messages at standard log levels
type Logger interface {
	// Debugf writes a log message. Debugf level is for low-level information or non-significant application events.
	Debugf(format string, args ...interface{})

	// Infof writes a log message. Info level is for recording significant application events that are not errors.
	Infof(format string, args ...interface{})

	// Warnf writes a log message. Warn level is for advisory messages about conditions that are not necessarily errors.
	Warnf(format string, args ...interface{})

	// Errorf writes a log message. Error level is for events where the application could not do the right thing.
	Errorf(format string, args ...interface{})
}
