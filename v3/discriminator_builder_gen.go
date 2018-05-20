package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:06:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// DiscriminatorBuilder is used to build an instance of Discriminator. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Discriminator
type DiscriminatorBuilder struct {
	target *discriminator
}

// Build finalizes the building process for Discriminator and returns the result
func (b *DiscriminatorBuilder) Build() Discriminator {
	return b.target
}

// NewDiscriminator creates a new builder object for Discriminator
func NewDiscriminator() *DiscriminatorBuilder {
	return &DiscriminatorBuilder{
		target: &discriminator{},
	}
}

// PropertyName sets the PropertyName field for object Discriminator.
func (b *DiscriminatorBuilder) PropertyName(v string) *DiscriminatorBuilder {
	b.target.propertyName = v
	return b
}

// Mapping sets the Mapping field for object Discriminator.
func (b *DiscriminatorBuilder) Mapping(v map[string]string) *DiscriminatorBuilder {
	b.target.mapping = v
	return b
}