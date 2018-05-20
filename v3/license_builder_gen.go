package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// LicenseBuilder is used to build an instance of License. The user must
// call `Build()` after providing all the necessary information to
// build an instance of License
type LicenseBuilder struct {
	target *license
}

// Build finalizes the building process for License and returns the result
func (b *LicenseBuilder) Build() License {
	return b.target
}

// NewLicense creates a new builder object for License
func NewLicense(name string) *LicenseBuilder {
	return &LicenseBuilder{
		target: &license{
			name: name,
		},
	}
}

// URL sets the URL field for object License.
func (b *LicenseBuilder) URL(v string) *LicenseBuilder {
	b.target.uRL = v
	return b
}