package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
	"sync"
)

var _ = errors.Cause

// SecurityRequirementBuilder is used to build an instance of SecurityRequirement. The user must
// call `Build()` after providing all the necessary information to
// build an instance of SecurityRequirement.
// Builders may NOT be reused. It must be created for every instance
// of SecurityRequirement that you want to create
type SecurityRequirementBuilder struct {
	mu     sync.Mutex
	target *securityRequirement
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *SecurityRequirementBuilder) MustBuild(options ...Option) SecurityRequirement {
	v, err := b.Build(options...)
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for SecurityRequirement and returns the result
// By default, Build() will validate if the given structure is valid
func (b *SecurityRequirementBuilder) Build(options ...Option) (SecurityRequirement, error) {
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

// NewSecurityRequirement creates a new builder object for SecurityRequirement
func NewSecurityRequirement() *SecurityRequirementBuilder {
	var b SecurityRequirementBuilder
	b.target = &securityRequirement{}
	return &b
}

// Name sets the name field for object SecurityRequirement.

func (b *SecurityRequirementBuilder) Name(v string) *SecurityRequirementBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.target == nil {
		return b
	}
	b.target.name = v
	return b
}

// Reference sets the $ref (reference) field for object SecurityRequirement.
func (b *SecurityRequirementBuilder) Reference(v string) *SecurityRequirementBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.target == nil {
		return b
	}
	b.target.reference = v
	return b
}

// Extension sets an arbitrary element (an extension) to the
// object SecurityRequirement. The extension name should start with a "x-"
func (b *SecurityRequirementBuilder) Extension(name string, value interface{}) *SecurityRequirementBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.target == nil {
		return b
	}
	b.target.extensions[name] = value
	return b
}
