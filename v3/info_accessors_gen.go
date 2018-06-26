package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *info) Title() string {
	return v.title
}

func (v *info) Description() string {
	return v.description
}

func (v *info) TermsOfService() string {
	return v.termsOfService
}

func (v *info) Contact() Contact {
	return v.contact
}

func (v *info) License() License {
	return v.license
}

func (v *info) Version() string {
	return v.version
}

func (v *info) Reference() string {
	return v.reference
}

func (v *info) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *info) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *info) recurseValidate() error {
	if elem := v.contact; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "contact"`)
		}
	}
	if elem := v.license; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "license"`)
		}
	}
	return nil
}
