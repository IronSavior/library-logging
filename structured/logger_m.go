package logging

// LoggerM accepts messages at standard log levels with optional metadata
type LoggerM interface {
	// Debug is for low-level information or non-significant application events
	Debug(msg string, metas ...Metadata)

	// Info for recording significant application events that are not errors
	Info(msg string, metas ...Metadata)

	// Warn is for advisory messages about conditions that are not necessarily errors
	Warn(msg string, metas ...Metadata)

	// Error is for events where the application could not do the right thing
	Error(msg string, metas ...Metadata)
}
