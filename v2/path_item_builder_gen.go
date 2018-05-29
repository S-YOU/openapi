package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T10:54:12+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// PathItemBuilder is used to build an instance of PathItem. The user must
// call `Do()` after providing all the necessary information to
// build an instance of PathItem
type PathItemBuilder struct {
	target *pathItem
}

// Do finalizes the building process for PathItem and returns the result
func (b *PathItemBuilder) Do() (PathItem, error) {
	if err := b.target.Validate(); err != nil {
		return nil, errors.Wrap(err, `validation failed`)
	}
	return b.target, nil
}

// NewPathItem creates a new builder object for PathItem
func NewPathItem() *PathItemBuilder {
	return &PathItemBuilder{
		target: &pathItem{},
	}
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

// Parameters sets the Parameters field for object PathItem.
func (b *PathItemBuilder) Parameters(v ...Parameter) *PathItemBuilder {
	b.target.parameters = v
	return b
}

// Reference sets the $ref (reference) field for object PathItem.
func (b *PathItemBuilder) Reference(v string) *PathItemBuilder {
	b.target.reference = v
	return b
}
