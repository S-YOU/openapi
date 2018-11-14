package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
	"sync"
)

var _ = errors.Cause

// ExternalDocumentationBuilder is used to build an instance of ExternalDocumentation. The user must
// call `Build()` after providing all the necessary information to
// build an instance of ExternalDocumentation.
// Builders may NOT be reused. It must be created for every instance
// of ExternalDocumentation that you want to create
type ExternalDocumentationBuilder struct {
	mu     sync.Mutex
	target *externalDocumentation
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *ExternalDocumentationBuilder) MustBuild(options ...Option) ExternalDocumentation {
	v, err := b.Build(options...)
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for ExternalDocumentation and returns the result
// By default, Build() will validate if the given structure is valid
func (b *ExternalDocumentationBuilder) Build(options ...Option) (ExternalDocumentation, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.target == nil {
		return nil, errors.New(`builder has already been used`)
	}
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
	defer func() { b.target = nil }()
	return b.target, nil
}

// NewExternalDocumentation creates a new builder object for ExternalDocumentation
func NewExternalDocumentation(url string) *ExternalDocumentationBuilder {
	var b ExternalDocumentationBuilder
	b.target = &externalDocumentation{
		url: url,
	}
	return &b
}

// Description sets the description field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Description(v string) *ExternalDocumentationBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.target == nil {
		return b
	}
	b.target.description = v
	return b
}

// Reference sets the $ref (reference) field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Reference(v string) *ExternalDocumentationBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.target == nil {
		return b
	}
	b.target.reference = v
	return b
}

// Extension sets an arbitrary element (an extension) to the
// object ExternalDocumentation. The extension name should start with a "x-"
func (b *ExternalDocumentationBuilder) Extension(name string, value interface{}) *ExternalDocumentationBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.target == nil {
		return b
	}
	b.target.extensions[name] = value
	return b
}
