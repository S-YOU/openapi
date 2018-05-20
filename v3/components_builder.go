package openapi

// Schema sets the schema identified by `name` to `s`
func (b *ComponentsBuilder) Schema(name string, s Schema) *ComponentsBuilder {
	if b.target.schemas == nil {
		b.target.schemas = make(map[string]Schema)
	}
	b.target.schemas[name] = s.Clone()
	return b
}

