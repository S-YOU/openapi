package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// ContactBuilder is used to build an instance of Contact. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Contact
type ContactBuilder struct {
	target *contact
}

// Do finalizes the building process for Contact and returns the result
func (b *ContactBuilder) Do(options ...Option) (Contact, error) {
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

// NewContact creates a new builder object for Contact
func NewContact() *ContactBuilder {
	return &ContactBuilder{
		target: &contact{},
	}
}

// Name sets the Name field for object Contact.
func (b *ContactBuilder) Name(v string) *ContactBuilder {
	b.target.name = v
	return b
}

// URL sets the URL field for object Contact.
func (b *ContactBuilder) URL(v string) *ContactBuilder {
	b.target.url = v
	return b
}

// Email sets the Email field for object Contact.
func (b *ContactBuilder) Email(v string) *ContactBuilder {
	b.target.email = v
	return b
}

// Reference sets the $ref (reference) field for object Contact.
func (b *ContactBuilder) Reference(v string) *ContactBuilder {
	b.target.reference = v
	return b
}
