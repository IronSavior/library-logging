package logging

// Error is a wrapper type for getting Metadata from standard errors
type Error struct{ error }

// LogMeta implements Metadata for Error
func (err Error) LogMeta() map[string]interface{} {
	return Meta{"error": err.Error()}
}
