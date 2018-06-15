package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// ExternalDocumentationBuilder is used to build an instance of ExternalDocumentation. The user must
// call `Do()` after providing all the necessary information to
// build an instance of ExternalDocumentation
type ExternalDocumentationBuilder struct {
	target *externalDocumentation
}

// Do finalizes the building process for ExternalDocumentation and returns the result
func (b *ExternalDocumentationBuilder) Do(options ...Option) (ExternalDocumentation, error) {
	validate := true
	for _, option := range options {
		switch option.Name() {
		case optkeyValidate:
			validate = option.Value().(bool)
		}
	}
	if validate {
		if err := b.target.Validate(false); err != nil {
			return nil, errors.Wrap(err, `validation failed`)
		}
	}
	return b.target, nil
}

// NewExternalDocumentation creates a new builder object for ExternalDocumentation
func NewExternalDocumentation(url string) *ExternalDocumentationBuilder {
	return &ExternalDocumentationBuilder{
		target: &externalDocumentation{
			url: url,
		},
	}
}

// Description sets the description field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Description(v string) *ExternalDocumentationBuilder {
	b.target.description = v
	return b
}

// Reference sets the $ref (reference) field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Reference(v string) *ExternalDocumentationBuilder {
	b.target.reference = v
	return b
}

// Extension sets an arbitrary element (an extension) to the
// object ExternalDocumentation. The extension name should start with a "x-"
func (b *ExternalDocumentationBuilder) Extension(name string, value interface{}) *ExternalDocumentationBuilder {
	b.target.extensions[name] = value
	return b
}
