package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:56:43+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// SchemaMutator is used to build an instance of Schema. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Schema with new values
type SchemaMutator struct {
	proxy  *schema
	target *schema
}

// Do finalizes the matuation process for Schema and returns the result
func (b *SchemaMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateSchema creates a new mutator object for Schema
func MutateSchema(v Schema) *SchemaMutator {
	return &SchemaMutator{
		target: v.(*schema),
		proxy:  v.Clone().(*schema),
	}
}

// Reference sets the Reference field for object Schema.
func (b *SchemaMutator) Reference(v string) *SchemaMutator {
	b.proxy.reference = v
	return b
}

// Title sets the Title field for object Schema.
func (b *SchemaMutator) Title(v string) *SchemaMutator {
	b.proxy.title = v
	return b
}

// MultipleOf sets the MultipleOf field for object Schema.
func (b *SchemaMutator) MultipleOf(v float64) *SchemaMutator {
	b.proxy.multipleOf = v
	return b
}

// Maximum sets the Maximum field for object Schema.
func (b *SchemaMutator) Maximum(v float64) *SchemaMutator {
	b.proxy.maximum = v
	return b
}

// ExclusiveMaximum sets the ExclusiveMaximum field for object Schema.
func (b *SchemaMutator) ExclusiveMaximum(v float64) *SchemaMutator {
	b.proxy.exclusiveMaximum = v
	return b
}

// Minimum sets the Minimum field for object Schema.
func (b *SchemaMutator) Minimum(v float64) *SchemaMutator {
	b.proxy.minimum = v
	return b
}

// ExclusiveMinimum sets the ExclusiveMinimum field for object Schema.
func (b *SchemaMutator) ExclusiveMinimum(v float64) *SchemaMutator {
	b.proxy.exclusiveMinimum = v
	return b
}

// MaxLength sets the MaxLength field for object Schema.
func (b *SchemaMutator) MaxLength(v int) *SchemaMutator {
	b.proxy.maxLength = v
	return b
}

// MinLength sets the MinLength field for object Schema.
func (b *SchemaMutator) MinLength(v int) *SchemaMutator {
	b.proxy.minLength = v
	return b
}

// Pattern sets the Pattern field for object Schema.
func (b *SchemaMutator) Pattern(v string) *SchemaMutator {
	b.proxy.pattern = v
	return b
}

// MaxItems sets the MaxItems field for object Schema.
func (b *SchemaMutator) MaxItems(v int) *SchemaMutator {
	b.proxy.maxItems = v
	return b
}

// MinItems sets the MinItems field for object Schema.
func (b *SchemaMutator) MinItems(v int) *SchemaMutator {
	b.proxy.minItems = v
	return b
}

// UniqueItems sets the UniqueItems field for object Schema.
func (b *SchemaMutator) UniqueItems(v bool) *SchemaMutator {
	b.proxy.uniqueItems = v
	return b
}

// MaxProperties sets the MaxProperties field for object Schema.
func (b *SchemaMutator) MaxProperties(v int) *SchemaMutator {
	b.proxy.maxProperties = v
	return b
}

// MinProperties sets the MinProperties field for object Schema.
func (b *SchemaMutator) MinProperties(v int) *SchemaMutator {
	b.proxy.minProperties = v
	return b
}

// Required sets the Required field for object Schema.
func (b *SchemaMutator) Required(v []string) *SchemaMutator {
	b.proxy.required = v
	return b
}

// Enum sets the Enum field for object Schema.
func (b *SchemaMutator) Enum(v []interface{}) *SchemaMutator {
	b.proxy.enum = v
	return b
}

// Type sets the Type field for object Schema.
func (b *SchemaMutator) Type(v PrimitiveType) *SchemaMutator {
	b.proxy.typ = v
	return b
}

// AllOf sets the AllOf field for object Schema.
func (b *SchemaMutator) AllOf(v []Schema) *SchemaMutator {
	b.proxy.allOf = v
	return b
}

// OneOf sets the OneOf field for object Schema.
func (b *SchemaMutator) OneOf(v []Schema) *SchemaMutator {
	b.proxy.oneOf = v
	return b
}

// AnyOf sets the AnyOf field for object Schema.
func (b *SchemaMutator) AnyOf(v []Schema) *SchemaMutator {
	b.proxy.anyOf = v
	return b
}

// Not sets the Not field for object Schema.
func (b *SchemaMutator) Not(v Schema) *SchemaMutator {
	b.proxy.not = v
	return b
}

// Items sets the Items field for object Schema.
func (b *SchemaMutator) Items(v Schema) *SchemaMutator {
	b.proxy.items = v
	return b
}

// Properties sets the Properties field for object Schema.
func (b *SchemaMutator) Properties(v map[string]Schema) *SchemaMutator {
	b.proxy.properties = v
	return b
}

// Format sets the Format field for object Schema.
func (b *SchemaMutator) Format(v string) *SchemaMutator {
	b.proxy.format = v
	return b
}

// Default sets the Default field for object Schema.
func (b *SchemaMutator) Default(v interface{}) *SchemaMutator {
	b.proxy.defaultValue = v
	return b
}

// Nullable sets the Nullable field for object Schema.
func (b *SchemaMutator) Nullable(v bool) *SchemaMutator {
	b.proxy.nullable = v
	return b
}

// Discriminator sets the Discriminator field for object Schema.
func (b *SchemaMutator) Discriminator(v Discriminator) *SchemaMutator {
	b.proxy.discriminator = v
	return b
}

// ReadOnly sets the ReadOnly field for object Schema.
func (b *SchemaMutator) ReadOnly(v bool) *SchemaMutator {
	b.proxy.readOnly = v
	return b
}

// WriteOnly sets the WriteOnly field for object Schema.
func (b *SchemaMutator) WriteOnly(v bool) *SchemaMutator {
	b.proxy.writeOnly = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object Schema.
func (b *SchemaMutator) ExternalDocs(v ExternalDocumentation) *SchemaMutator {
	b.proxy.externalDocs = v
	return b
}

// Example sets the Example field for object Schema.
func (b *SchemaMutator) Example(v interface{}) *SchemaMutator {
	b.proxy.example = v
	return b
}

// Deprecated sets the Deprecated field for object Schema.
func (b *SchemaMutator) Deprecated(v bool) *SchemaMutator {
	b.proxy.deprecated = v
	return b
}
