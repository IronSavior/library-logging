package logging

// MergeMeta implements Metadata by merging Metadata of many values
type MergeMeta []Metadata

// LogMeta reduces the slice of Metadata to one map
// NOTE: Later values will overwrite earlier ones in cases of key collision
func (metas MergeMeta) LogMeta() map[string]interface{} {
	var merged map[string]interface{}

	for _, meta := range metas {
		for k, v := range meta.LogMeta() {
			if merged == nil {
				merged = map[string]interface{}{}
			}

			merged[k] = v
		}
	}

	return merged
}
