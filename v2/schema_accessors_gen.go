package openapi

// This file was automatically generated by gentyeps.go on 2018-05-29T10:54:12+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *schema) Name() string {
	return v.name
}

func (v *schema) Type() PrimitiveType {
	return v.typ
}

func (v *schema) Format() string {
	return v.format
}

func (v *schema) Title() string {
	return v.title
}

func (v *schema) MultipleOf() float64 {
	return v.multipleOf
}

func (v *schema) Maximum() float64 {
	return v.maximum
}

func (v *schema) ExclusiveMaximum() float64 {
	return v.exclusiveMaximum
}

func (v *schema) Minimum() float64 {
	return v.minimum
}

func (v *schema) ExclusiveMinimum() float64 {
	return v.exclusiveMinimum
}

func (v *schema) MaxLength() int {
	return v.maxLength
}

func (v *schema) MinLength() int {
	return v.minLength
}

func (v *schema) Pattern() string {
	return v.pattern
}

func (v *schema) MaxItems() int {
	return v.maxItems
}

func (v *schema) MinItems() int {
	return v.minItems
}

func (v *schema) UniqueItems() bool {
	return v.uniqueItems
}

func (v *schema) MaxProperties() int {
	return v.maxProperties
}

func (v *schema) MinProperties() int {
	return v.minProperties
}

func (v *schema) Required() *StringListIterator {
	var items []interface{}
	for _, item := range v.required {
		items = append(items, item)
	}
	var iter StringListIterator
	iter.items = items
	return &iter
}

func (v *schema) Enum() *InterfaceListIterator {
	var items []interface{}
	for _, item := range v.enum {
		items = append(items, item)
	}
	var iter InterfaceListIterator
	iter.items = items
	return &iter
}

func (v *schema) AllOf() *SchemaListIterator {
	var items []interface{}
	for _, item := range v.allOf {
		items = append(items, item)
	}
	var iter SchemaListIterator
	iter.items = items
	return &iter
}

func (v *schema) Items() Schema {
	return v.items
}

func (v *schema) Properties() *SchemaMapIterator {
	var items []interface{}
	for key, item := range v.properties {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SchemaMapIterator
	iter.list.items = items
	return &iter
}

func (v *schema) AdditionaProperties() *SchemaMapIterator {
	var items []interface{}
	for key, item := range v.additionaProperties {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SchemaMapIterator
	iter.list.items = items
	return &iter
}

func (v *schema) DefaultValue() interface{} {
	return v.defaultValue
}

func (v *schema) Discriminator() string {
	return v.discriminator
}

func (v *schema) ReadOnly() bool {
	return v.readOnly
}

func (v *schema) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

func (v *schema) Example() interface{} {
	return v.example
}

func (v *schema) Deprecated() bool {
	return v.deprecated
}

func (v *schema) XML() XML {
	return v.xml
}

func (v *schema) Reference() string {
	return v.reference
}

func (v *schema) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *schema) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *schema) Validate() error {
	return nil
}
