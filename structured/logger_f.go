package logging

// LoggerF accepts messages at standard log levels using Printf-style interpolation
type LoggerF interface {
	// Debug is for low-level information or non-significant application events
	Debugf(format string, args ...interface{})

	// Info for recording significant application events that are not errors
	Infof(format string, args ...interface{})

	// Warn is for advisory messages about conditions that are not necessarily errors
	Warnf(format string, args ...interface{})

	// Error is for events where the application could not do the right thing
	Errorf(format string, args ...interface{})
}
