package logging

// Meta is a map type that implements Metadata
type Meta map[string]interface{}

// LogMeta satisfies Metadata by returning itself
func (m Meta) LogMeta() map[string]interface{} {
	return m
}
