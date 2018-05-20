package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// ParameterBuilder is used to build an instance of Parameter. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Parameter
type ParameterBuilder struct {
	target *parameter
}

// Build finalizes the building process for Parameter and returns the result
func (b *ParameterBuilder) Build() Parameter {
	return b.target
}

// NewParameter creates a new builder object for Parameter
func NewParameter(name string, in Location) *ParameterBuilder {
	return &ParameterBuilder{
		target: &parameter{
			required: defaultParameterRequiredFromLocation(in),
			name:     name,
			in:       in,
		},
	}
}

// Required sets the Required field for object Parameter. If this is not called,
// a default value (defaultParameterRequiredFromLocation(in)) is assigned to this field
func (b *ParameterBuilder) Required(v bool) *ParameterBuilder {
	b.target.required = v
	return b
}

// Description sets the Description field for object Parameter.
func (b *ParameterBuilder) Description(v string) *ParameterBuilder {
	b.target.description = v
	return b
}

// Deprecated sets the Deprecated field for object Parameter.
func (b *ParameterBuilder) Deprecated(v bool) *ParameterBuilder {
	b.target.deprecated = v
	return b
}

// AllowEmptyValue sets the AllowEmptyValue field for object Parameter.
func (b *ParameterBuilder) AllowEmptyValue(v bool) *ParameterBuilder {
	b.target.allowEmptyValue = v
	return b
}

// Explode sets the Explode field for object Parameter.
func (b *ParameterBuilder) Explode(v bool) *ParameterBuilder {
	b.target.explode = v
	return b
}

// AllowReserved sets the AllowReserved field for object Parameter.
func (b *ParameterBuilder) AllowReserved(v bool) *ParameterBuilder {
	b.target.allowReserved = v
	return b
}

// Schema sets the Schema field for object Parameter.
func (b *ParameterBuilder) Schema(v Schema) *ParameterBuilder {
	b.target.schema = v
	return b
}

// Example sets the Example field for object Parameter.
func (b *ParameterBuilder) Example(v interface{}) *ParameterBuilder {
	b.target.example = v
	return b
}

// Examples sets the Examples field for object Parameter.
func (b *ParameterBuilder) Examples(v map[string]Example) *ParameterBuilder {
	b.target.examples = v
	return b
}

// Content sets the Content field for object Parameter.
func (b *ParameterBuilder) Content(v map[string]MediaType) *ParameterBuilder {
	b.target.content = v
	return b
}
