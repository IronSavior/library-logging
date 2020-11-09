package logging

// Metadata allows arbitrary types to contribute log metadata
type Metadata interface {
	LogMeta() map[string]interface{}
}
