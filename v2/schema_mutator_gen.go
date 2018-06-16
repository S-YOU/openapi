package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"log"
)

var _ = log.Printf

// SchemaMutator is used to build an instance of Schema. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Schema with new values
type SchemaMutator struct {
	proxy  *schema
	target *schema
}

// Do finalizes the matuation process for Schema and returns the result
func (m *SchemaMutator) Do() error {
	*m.target = *m.proxy
	return nil
}

// MutateSchema creates a new mutator object for Schema
func MutateSchema(v Schema) *SchemaMutator {
	return &SchemaMutator{
		target: v.(*schema),
		proxy:  v.Clone().(*schema),
	}
}

// Name sets the Name field for object Schema.
func (m *SchemaMutator) Name(v string) *SchemaMutator {
	m.proxy.name = v
	return m
}

// Type sets the Type field for object Schema.
func (m *SchemaMutator) Type(v PrimitiveType) *SchemaMutator {
	m.proxy.typ = v
	return m
}

// Format sets the Format field for object Schema.
func (m *SchemaMutator) Format(v string) *SchemaMutator {
	m.proxy.format = v
	return m
}

// Title sets the Title field for object Schema.
func (m *SchemaMutator) Title(v string) *SchemaMutator {
	m.proxy.title = v
	return m
}

// ClearMultipleOf clears the multipleOf field
func (m *SchemaMutator) ClearMultipleOf() *SchemaMutator {
	m.proxy.multipleOf = nil
	return m
}

// MultipleOf sets the multipleOf field.%!(EXTRA string=Schema)
func (m *SchemaMutator) MultipleOf(v float64) *SchemaMutator {
	m.proxy.multipleOf = &v
	return m
}

// ClearMaximum clears the maximum field
func (m *SchemaMutator) ClearMaximum() *SchemaMutator {
	m.proxy.maximum = nil
	return m
}

// Maximum sets the maximum field.%!(EXTRA string=Schema)
func (m *SchemaMutator) Maximum(v float64) *SchemaMutator {
	m.proxy.maximum = &v
	return m
}

// ClearExclusiveMaximum clears the exclusiveMaximum field
func (m *SchemaMutator) ClearExclusiveMaximum() *SchemaMutator {
	m.proxy.exclusiveMaximum = nil
	return m
}

// ExclusiveMaximum sets the exclusiveMaximum field.%!(EXTRA string=Schema)
func (m *SchemaMutator) ExclusiveMaximum(v float64) *SchemaMutator {
	m.proxy.exclusiveMaximum = &v
	return m
}

// ClearMinimum clears the minimum field
func (m *SchemaMutator) ClearMinimum() *SchemaMutator {
	m.proxy.minimum = nil
	return m
}

// Minimum sets the minimum field.%!(EXTRA string=Schema)
func (m *SchemaMutator) Minimum(v float64) *SchemaMutator {
	m.proxy.minimum = &v
	return m
}

// ClearExclusiveMinimum clears the exclusiveMinimum field
func (m *SchemaMutator) ClearExclusiveMinimum() *SchemaMutator {
	m.proxy.exclusiveMinimum = nil
	return m
}

// ExclusiveMinimum sets the exclusiveMinimum field.%!(EXTRA string=Schema)
func (m *SchemaMutator) ExclusiveMinimum(v float64) *SchemaMutator {
	m.proxy.exclusiveMinimum = &v
	return m
}

// ClearMaxLength clears the maxLength field
func (m *SchemaMutator) ClearMaxLength() *SchemaMutator {
	m.proxy.maxLength = nil
	return m
}

// MaxLength sets the maxLength field.%!(EXTRA string=Schema)
func (m *SchemaMutator) MaxLength(v int) *SchemaMutator {
	m.proxy.maxLength = &v
	return m
}

// ClearMinLength clears the minLength field
func (m *SchemaMutator) ClearMinLength() *SchemaMutator {
	m.proxy.minLength = nil
	return m
}

// MinLength sets the minLength field.%!(EXTRA string=Schema)
func (m *SchemaMutator) MinLength(v int) *SchemaMutator {
	m.proxy.minLength = &v
	return m
}

// Pattern sets the Pattern field for object Schema.
func (m *SchemaMutator) Pattern(v string) *SchemaMutator {
	m.proxy.pattern = v
	return m
}

// ClearMaxItems clears the maxItems field
func (m *SchemaMutator) ClearMaxItems() *SchemaMutator {
	m.proxy.maxItems = nil
	return m
}

// MaxItems sets the maxItems field.%!(EXTRA string=Schema)
func (m *SchemaMutator) MaxItems(v int) *SchemaMutator {
	m.proxy.maxItems = &v
	return m
}

// ClearMinItems clears the minItems field
func (m *SchemaMutator) ClearMinItems() *SchemaMutator {
	m.proxy.minItems = nil
	return m
}

// MinItems sets the minItems field.%!(EXTRA string=Schema)
func (m *SchemaMutator) MinItems(v int) *SchemaMutator {
	m.proxy.minItems = &v
	return m
}

// UniqueItems sets the UniqueItems field for object Schema.
func (m *SchemaMutator) UniqueItems(v bool) *SchemaMutator {
	m.proxy.uniqueItems = v
	return m
}

// ClearMaxProperties clears the maxProperties field
func (m *SchemaMutator) ClearMaxProperties() *SchemaMutator {
	m.proxy.maxProperties = nil
	return m
}

// MaxProperties sets the maxProperties field.%!(EXTRA string=Schema)
func (m *SchemaMutator) MaxProperties(v int) *SchemaMutator {
	m.proxy.maxProperties = &v
	return m
}

// ClearMinProperties clears the minProperties field
func (m *SchemaMutator) ClearMinProperties() *SchemaMutator {
	m.proxy.minProperties = nil
	return m
}

// MinProperties sets the minProperties field.%!(EXTRA string=Schema)
func (m *SchemaMutator) MinProperties(v int) *SchemaMutator {
	m.proxy.minProperties = &v
	return m
}

// ClearRequired clears all elements in required
func (m *SchemaMutator) ClearRequired() *SchemaMutator {
	_ = m.proxy.required.Clear()
	return m
}

// Required appends a value to required
func (m *SchemaMutator) Required(value string) *SchemaMutator {
	m.proxy.required = append(m.proxy.required, value)
	return m
}

// ClearEnum clears all elements in enum
func (m *SchemaMutator) ClearEnum() *SchemaMutator {
	_ = m.proxy.enum.Clear()
	return m
}

// Enum appends a value to enum
func (m *SchemaMutator) Enum(value interface{}) *SchemaMutator {
	m.proxy.enum = append(m.proxy.enum, value)
	return m
}

// ClearAllOf clears all elements in allOf
func (m *SchemaMutator) ClearAllOf() *SchemaMutator {
	_ = m.proxy.allOf.Clear()
	return m
}

// AllOf appends a value to allOf
func (m *SchemaMutator) AllOf(value Schema) *SchemaMutator {
	m.proxy.allOf = append(m.proxy.allOf, value)
	return m
}

// Items sets the Items field for object Schema.
func (m *SchemaMutator) Items(v Schema) *SchemaMutator {
	m.proxy.items = v
	return m
}

// ClearProperties removes all values in properties field
func (m *SchemaMutator) ClearProperties() *SchemaMutator {
	_ = m.proxy.properties.Clear()
	return m
}

// Property sets the value of properties
func (m *SchemaMutator) Property(key SchemaMapKey, value Schema) *SchemaMutator {
	if m.proxy.properties == nil {
		m.proxy.properties = SchemaMap{}
	}

	m.proxy.properties[key] = value
	return m
}

// ClearAdditionaProperties removes all values in additionaProperties field
func (m *SchemaMutator) ClearAdditionaProperties() *SchemaMutator {
	_ = m.proxy.additionaProperties.Clear()
	return m
}

// AdditionaProperty sets the value of additionaProperties
func (m *SchemaMutator) AdditionaProperty(key SchemaMapKey, value Schema) *SchemaMutator {
	if m.proxy.additionaProperties == nil {
		m.proxy.additionaProperties = SchemaMap{}
	}

	m.proxy.additionaProperties[key] = value
	return m
}

// DefaultValue sets the DefaultValue field for object Schema.
func (m *SchemaMutator) DefaultValue(v interface{}) *SchemaMutator {
	m.proxy.defaultValue = v
	return m
}

// Discriminator sets the Discriminator field for object Schema.
func (m *SchemaMutator) Discriminator(v string) *SchemaMutator {
	m.proxy.discriminator = v
	return m
}

// ReadOnly sets the ReadOnly field for object Schema.
func (m *SchemaMutator) ReadOnly(v bool) *SchemaMutator {
	m.proxy.readOnly = v
	return m
}

// ExternalDocs sets the ExternalDocs field for object Schema.
func (m *SchemaMutator) ExternalDocs(v ExternalDocumentation) *SchemaMutator {
	m.proxy.externalDocs = v
	return m
}

// Example sets the Example field for object Schema.
func (m *SchemaMutator) Example(v interface{}) *SchemaMutator {
	m.proxy.example = v
	return m
}

// Deprecated sets the Deprecated field for object Schema.
func (m *SchemaMutator) Deprecated(v bool) *SchemaMutator {
	m.proxy.deprecated = v
	return m
}

// XML sets the XML field for object Schema.
func (m *SchemaMutator) XML(v XML) *SchemaMutator {
	m.proxy.xml = v
	return m
}

// Extension sets an arbitrary extension field in Schema
func (m *SchemaMutator) Extension(name string, value interface{}) *SchemaMutator {
	if m.proxy.extensions == nil {
		m.proxy.extensions = Extensions{}
	}
	m.proxy.extensions[name] = value
	return m
}
