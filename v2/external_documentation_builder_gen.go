package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T20:40:48+09:00
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
func (b *ExternalDocumentationBuilder) Do() (ExternalDocumentation, error) {
	if err := b.target.Validate(); err != nil {
		return nil, errors.Wrap(err, `validation failed`)
	}
	return b.target, nil
}

// NewExternalDocumentation creates a new builder object for ExternalDocumentation
func NewExternalDocumentation(uRL string) *ExternalDocumentationBuilder {
	return &ExternalDocumentationBuilder{
		target: &externalDocumentation{
			uRL: uRL,
		},
	}
}

// Description sets the Description field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Description(v string) *ExternalDocumentationBuilder {
	b.target.description = v
	return b
}

// Reference sets the $ref (reference) field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Reference(v string) *ExternalDocumentationBuilder {
	b.target.reference = v
	return b
}
