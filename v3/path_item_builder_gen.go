package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T19:11:33+09:00
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

// Get sets the Get field for object PathItem.
func (b *PathItemBuilder) Get(v Operation) *PathItemBuilder {
	b.target.get = v
	return b
}

// Put sets the Put field for object PathItem.
func (b *PathItemBuilder) Put(v Operation) *PathItemBuilder {
	b.target.put = v
	return b
}

// Post sets the Post field for object PathItem.
func (b *PathItemBuilder) Post(v Operation) *PathItemBuilder {
	b.target.post = v
	return b
}

// Delete sets the Delete field for object PathItem.
func (b *PathItemBuilder) Delete(v Operation) *PathItemBuilder {
	b.target.delete = v
	return b
}

// Options sets the Options field for object PathItem.
func (b *PathItemBuilder) Options(v Operation) *PathItemBuilder {
	b.target.options = v
	return b
}

// Head sets the Head field for object PathItem.
func (b *PathItemBuilder) Head(v Operation) *PathItemBuilder {
	b.target.head = v
	return b
}

// Patch sets the Patch field for object PathItem.
func (b *PathItemBuilder) Patch(v Operation) *PathItemBuilder {
	b.target.patch = v
	return b
}

// Trace sets the Trace field for object PathItem.
func (b *PathItemBuilder) Trace(v Operation) *PathItemBuilder {
	b.target.trace = v
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
