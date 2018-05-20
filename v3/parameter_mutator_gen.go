package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:56:43+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// ParameterMutator is used to build an instance of Parameter. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Parameter with new values
type ParameterMutator struct {
	proxy  *parameter
	target *parameter
}

// Do finalizes the matuation process for Parameter and returns the result
func (b *ParameterMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateParameter creates a new mutator object for Parameter
func MutateParameter(v Parameter) *ParameterMutator {
	return &ParameterMutator{
		target: v.(*parameter),
		proxy:  v.Clone().(*parameter),
	}
}

// Name sets the Name field for object Parameter.
func (b *ParameterMutator) Name(v string) *ParameterMutator {
	b.proxy.name = v
	return b
}

// In sets the In field for object Parameter.
func (b *ParameterMutator) In(v Location) *ParameterMutator {
	b.proxy.in = v
	return b
}

// Required sets the Required field for object Parameter.
func (b *ParameterMutator) Required(v bool) *ParameterMutator {
	b.proxy.required = v
	return b
}

// Description sets the Description field for object Parameter.
func (b *ParameterMutator) Description(v string) *ParameterMutator {
	b.proxy.description = v
	return b
}

// Deprecated sets the Deprecated field for object Parameter.
func (b *ParameterMutator) Deprecated(v bool) *ParameterMutator {
	b.proxy.deprecated = v
	return b
}

// AllowEmptyValue sets the AllowEmptyValue field for object Parameter.
func (b *ParameterMutator) AllowEmptyValue(v bool) *ParameterMutator {
	b.proxy.allowEmptyValue = v
	return b
}

// Explode sets the Explode field for object Parameter.
func (b *ParameterMutator) Explode(v bool) *ParameterMutator {
	b.proxy.explode = v
	return b
}

// AllowReserved sets the AllowReserved field for object Parameter.
func (b *ParameterMutator) AllowReserved(v bool) *ParameterMutator {
	b.proxy.allowReserved = v
	return b
}

// Schema sets the Schema field for object Parameter.
func (b *ParameterMutator) Schema(v Schema) *ParameterMutator {
	b.proxy.schema = v
	return b
}

// Example sets the Example field for object Parameter.
func (b *ParameterMutator) Example(v interface{}) *ParameterMutator {
	b.proxy.example = v
	return b
}

// Examples sets the Examples field for object Parameter.
func (b *ParameterMutator) Examples(v map[string]Example) *ParameterMutator {
	b.proxy.examples = v
	return b
}

// Content sets the Content field for object Parameter.
func (b *ParameterMutator) Content(v map[string]MediaType) *ParameterMutator {
	b.proxy.content = v
	return b
}
