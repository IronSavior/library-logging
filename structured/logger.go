package logging

// Logger is an alias to the preferred Logger variation
type Logger = LoggerM

// LoggerWM is a Logger that can also derive a new Logger with implicit metadata
type LoggerWM interface {
	// WithMeta creates a new Logger with additional implicit metadata
	WithMeta(metas ...Metadata) LoggerWM

	Logger
}
