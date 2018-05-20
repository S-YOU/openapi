package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// ExternalDocumentationBuilder is used to build an instance of ExternalDocumentation. The user must
// call `Build()` after providing all the necessary information to
// build an instance of ExternalDocumentation
type ExternalDocumentationBuilder struct {
	target *externalDocumentation
}

// Build finalizes the building process for ExternalDocumentation and returns the result
func (b *ExternalDocumentationBuilder) Build() ExternalDocumentation {
	return b.target
}

// NewExternalDocumentation creates a new builder object for ExternalDocumentation
func NewExternalDocumentation() *ExternalDocumentationBuilder {
	return &ExternalDocumentationBuilder{
		target: &externalDocumentation{},
	}
}

// Description sets the Description field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Description(v string) *ExternalDocumentationBuilder {
	b.target.description = v
	return b
}

// URL sets the URL field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) URL(v string) *ExternalDocumentationBuilder {
	b.target.uRL = v
	return b
}