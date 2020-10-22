package logging

// Fields are key/value pairs of metadata that can be added to a structured logger
type Fields map[string]interface{}

// Fieldser allows any type to provide log fields
type Fieldser interface {
	Fields() Fields
}

// Fields implements Fieldser
func (f Fields) Fields() Fields {
	return f
}

// MergeFields implements Fields() by merging a slice of objects that provide log fields in the order given
type MergeFields []Fieldser

// Fields implements Fieldser
func (mf MergeFields) Fields() Fields {
	var merged Fields

	for _, f := range mf {
		for k, v := range f.Fields() {
			if merged == nil {
				merged = Fields{}
			}

			merged[k] = v
		}
	}

	return merged
}
