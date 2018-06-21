package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// PathItemBuilder is used to build an instance of PathItem. The user must
// call `Do()` after providing all the necessary information to
// build an instance of PathItem
type PathItemBuilder struct {
	target *pathItem
}

// Do finalizes the building process for PathItem and returns the result
func (b *PathItemBuilder) Do() PathItem {
	return b.target
}

// NewPathItem creates a new builder object for PathItem
func NewPathItem() *PathItemBuilder {
	return &PathItemBuilder{
		target: &pathItem{},
	}
}

// Path sets the Path field for object PathItem.
func (b *PathItemBuilder) Path(v string) *PathItemBuilder {
	b.target.path = v
	return b
}

// Summary sets the Summary field for object PathItem.
func (b *PathItemBuilder) Summary(v string) *PathItemBuilder {
	b.target.summary = v
	return b
}

// Description sets the Description field for object PathItem.
func (b *PathItemBuilder) Description(v string) *PathItemBuilder {
	b.target.description = v
	return b
}

// Servers sets the Servers field for object PathItem.
func (b *PathItemBuilder) Servers(v []Server) *PathItemBuilder {
	b.target.servers = v
	return b
}

// Parameters sets the Parameters field for object PathItem.
func (b *PathItemBuilder) Parameters(v []Parameter) *PathItemBuilder {
	b.target.parameters = v
	return b
}

// Reference sets the $ref (reference) field for object PathItem.
func (b *PathItemBuilder) Reference(v string) *PathItemBuilder {
	b.target.reference = v
	return b
}
