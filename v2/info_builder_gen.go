package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T10:54:12+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// InfoBuilder is used to build an instance of Info. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Info
type InfoBuilder struct {
	target *info
}

// Do finalizes the building process for Info and returns the result
func (b *InfoBuilder) Do() (Info, error) {
	if err := b.target.Validate(); err != nil {
		return nil, errors.Wrap(err, `validation failed`)
	}
	return b.target, nil
}

// NewInfo creates a new builder object for Info
func NewInfo(title string, version string) *InfoBuilder {
	return &InfoBuilder{
		target: &info{
			title:   title,
			version: version,
		},
	}
}

// Description sets the Description field for object Info.
func (b *InfoBuilder) Description(v string) *InfoBuilder {
	b.target.description = v
	return b
}

// TermsOfService sets the TermsOfService field for object Info.
func (b *InfoBuilder) TermsOfService(v string) *InfoBuilder {
	b.target.termsOfService = v
	return b
}

// Contact sets the Contact field for object Info.
func (b *InfoBuilder) Contact(v Contact) *InfoBuilder {
	b.target.contact = v
	return b
}

// License sets the License field for object Info.
func (b *InfoBuilder) License(v License) *InfoBuilder {
	b.target.license = v
	return b
}

// Reference sets the $ref (reference) field for object Info.
func (b *InfoBuilder) Reference(v string) *InfoBuilder {
	b.target.reference = v
	return b
}
