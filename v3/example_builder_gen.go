package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// ExampleBuilder is used to build an instance of Example. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Example
type ExampleBuilder struct {
	target *example
}

// Do finalizes the building process for Example and returns the result
func (b *ExampleBuilder) Do() Example {
	return b.target
}

// NewExample creates a new builder object for Example
func NewExample() *ExampleBuilder {
	return &ExampleBuilder{
		target: &example{},
	}
}

// Description sets the Description field for object Example.
func (b *ExampleBuilder) Description(v string) *ExampleBuilder {
	b.target.description = v
	return b
}

// Value sets the Value field for object Example.
func (b *ExampleBuilder) Value(v interface{}) *ExampleBuilder {
	b.target.value = v
	return b
}

// ExternalValue sets the ExternalValue field for object Example.
func (b *ExampleBuilder) ExternalValue(v string) *ExampleBuilder {
	b.target.externalValue = v
	return b
}

// Reference sets the $ref (reference) field for object Example.
func (b *ExampleBuilder) Reference(v string) *ExampleBuilder {
	b.target.reference = v
	return b
}
