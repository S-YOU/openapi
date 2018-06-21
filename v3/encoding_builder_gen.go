package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// EncodingBuilder is used to build an instance of Encoding. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Encoding
type EncodingBuilder struct {
	target *encoding
}

// Do finalizes the building process for Encoding and returns the result
func (b *EncodingBuilder) Do() Encoding {
	return b.target
}

// NewEncoding creates a new builder object for Encoding
func NewEncoding() *EncodingBuilder {
	return &EncodingBuilder{
		target: &encoding{},
	}
}

// ContentType sets the ContentType field for object Encoding.
func (b *EncodingBuilder) ContentType(v string) *EncodingBuilder {
	b.target.contentType = v
	return b
}

// Headers sets the Headers field for object Encoding.
func (b *EncodingBuilder) Headers(v map[string]Header) *EncodingBuilder {
	b.target.headers = v
	return b
}

// Explode sets the Explode field for object Encoding.
func (b *EncodingBuilder) Explode(v bool) *EncodingBuilder {
	b.target.explode = v
	return b
}

// AllowReserved sets the AllowReserved field for object Encoding.
func (b *EncodingBuilder) AllowReserved(v bool) *EncodingBuilder {
	b.target.allowReserved = v
	return b
}

// Reference sets the $ref (reference) field for object Encoding.
func (b *EncodingBuilder) Reference(v string) *EncodingBuilder {
	b.target.reference = v
	return b
}
