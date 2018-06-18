package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

var _ = json.Unmarshal
var _ = fmt.Fprintf
var _ = log.Printf
var _ = strconv.ParseInt
var _ = errors.Cause

type schemaMarshalProxy struct {
	Reference           string                `json:"$ref,omitempty"`
	Type                PrimitiveType         `json:"type,omitempty"`
	Format              string                `json:"format,omitempty"`
	Title               string                `json:"title,omitempty"`
	MultipleOf          *float64              `json:"multipleOf,omitempty"`
	Maximum             *float64              `json:"maximum,omitempty"`
	ExclusiveMaximum    *float64              `json:"exclusiveMaximum,omitempty"`
	Minimum             *float64              `json:"minimum,omitempty"`
	ExclusiveMinimum    *float64              `json:"exclusiveMinimum,omitempty"`
	MaxLength           *int                  `json:"maxLength,omitempty"`
	MinLength           *int                  `json:"minLength,omitempty"`
	Pattern             string                `json:"pattern,omitempty"`
	MaxItems            *int                  `json:"maxItems,omitempty"`
	MinItems            *int                  `json:"minItems,omitempty"`
	UniqueItems         bool                  `json:"uniqueItems,omitempty"`
	MaxProperties       *int                  `json:"maxProperties,omitempty"`
	MinProperties       *int                  `json:"minProperties,omitempty"`
	Required            StringList            `json:"required,omitempty"`
	Enum                InterfaceList         `json:"enum,omitempty"`
	AllOf               SchemaList            `json:"allOf,omitempty"`
	Items               Schema                `json:"items,omitempty"`
	Properties          SchemaMap             `json:"properties,omitempty"`
	AdditionaProperties SchemaMap             `json:"additionalProperties,omitempty"`
	Default             interface{}           `json:"default,omitempty"`
	Discriminator       string                `json:"discriminator,omitempty"`
	ReadOnly            bool                  `json:"readOnly,omitempty"`
	ExternalDocs        ExternalDocumentation `json:"externalDocs,omitempty"`
	Example             interface{}           `json:"example,omitempty"`
	Deprecated          bool                  `json:"deprecated,omitempty"`
	XML                 XML                   `json:"xml,omitempty"`
}

func (v *schema) MarshalJSON() ([]byte, error) {
	var proxy schemaMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Type = v.typ
	proxy.Format = v.format
	proxy.Title = v.title
	proxy.MultipleOf = v.multipleOf
	proxy.Maximum = v.maximum
	proxy.ExclusiveMaximum = v.exclusiveMaximum
	proxy.Minimum = v.minimum
	proxy.ExclusiveMinimum = v.exclusiveMinimum
	proxy.MaxLength = v.maxLength
	proxy.MinLength = v.minLength
	proxy.Pattern = v.pattern
	proxy.MaxItems = v.maxItems
	proxy.MinItems = v.minItems
	proxy.UniqueItems = v.uniqueItems
	proxy.MaxProperties = v.maxProperties
	proxy.MinProperties = v.minProperties
	proxy.Required = v.required
	proxy.Enum = v.enum
	proxy.AllOf = v.allOf
	proxy.Items = v.items
	proxy.Properties = v.properties
	proxy.AdditionaProperties = v.additionaProperties
	proxy.Default = v.defaultValue
	proxy.Discriminator = v.discriminator
	proxy.ReadOnly = v.readOnly
	proxy.ExternalDocs = v.externalDocs
	proxy.Example = v.example
	proxy.Deprecated = v.deprecated
	proxy.XML = v.xml
	buf, err := json.Marshal(proxy)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal struct`)
	}
	if len(v.extensions) > 0 {
		extBuf, err := json.Marshal(v.extensions)
		if err != nil || len(extBuf) <= 2 {
			return nil, errors.Wrap(err, `failed to marshal struct (extensions)`)
		}
		buf = append(append(buf[:len(buf)-1], ','), extBuf[1:]...)
	}
	return buf, nil
}

// UnmarshalJSON defines how schema is deserialized from JSON
func (v *schema) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if raw, ok := proxy["$ref"]; ok {
		if err := json.Unmarshal(raw, &v.reference); err != nil {
			return errors.Wrap(err, `failed to unmarshal $ref`)
		}
		return nil
	}

	mutator := MutateSchema(v)

	const typMapKey = "type"
	if raw, ok := proxy[typMapKey]; ok {
		var decoded PrimitiveType
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field type`)
		}
		mutator.Type(decoded)
		delete(proxy, typMapKey)
	}

	const formatMapKey = "format"
	if raw, ok := proxy[formatMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field format`)
		}
		mutator.Format(decoded)
		delete(proxy, formatMapKey)
	}

	const titleMapKey = "title"
	if raw, ok := proxy[titleMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field title`)
		}
		mutator.Title(decoded)
		delete(proxy, titleMapKey)
	}

	const multipleOfMapKey = "multipleOf"
	if raw, ok := proxy[multipleOfMapKey]; ok {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field multipleOf`)
		}
		mutator.MultipleOf(decoded)
		delete(proxy, multipleOfMapKey)
	}

	const maximumMapKey = "maximum"
	if raw, ok := proxy[maximumMapKey]; ok {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field maximum`)
		}
		mutator.Maximum(decoded)
		delete(proxy, maximumMapKey)
	}

	const exclusiveMaximumMapKey = "exclusiveMaximum"
	if raw, ok := proxy[exclusiveMaximumMapKey]; ok {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field exclusiveMaximum`)
		}
		mutator.ExclusiveMaximum(decoded)
		delete(proxy, exclusiveMaximumMapKey)
	}

	const minimumMapKey = "minimum"
	if raw, ok := proxy[minimumMapKey]; ok {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field minimum`)
		}
		mutator.Minimum(decoded)
		delete(proxy, minimumMapKey)
	}

	const exclusiveMinimumMapKey = "exclusiveMinimum"
	if raw, ok := proxy[exclusiveMinimumMapKey]; ok {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field exclusiveMinimum`)
		}
		mutator.ExclusiveMinimum(decoded)
		delete(proxy, exclusiveMinimumMapKey)
	}

	const maxLengthMapKey = "maxLength"
	if raw, ok := proxy[maxLengthMapKey]; ok {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field maxLength`)
		}
		mutator.MaxLength(decoded)
		delete(proxy, maxLengthMapKey)
	}

	const minLengthMapKey = "minLength"
	if raw, ok := proxy[minLengthMapKey]; ok {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field minLength`)
		}
		mutator.MinLength(decoded)
		delete(proxy, minLengthMapKey)
	}

	const patternMapKey = "pattern"
	if raw, ok := proxy[patternMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field pattern`)
		}
		mutator.Pattern(decoded)
		delete(proxy, patternMapKey)
	}

	const maxItemsMapKey = "maxItems"
	if raw, ok := proxy[maxItemsMapKey]; ok {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field maxItems`)
		}
		mutator.MaxItems(decoded)
		delete(proxy, maxItemsMapKey)
	}

	const minItemsMapKey = "minItems"
	if raw, ok := proxy[minItemsMapKey]; ok {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field minItems`)
		}
		mutator.MinItems(decoded)
		delete(proxy, minItemsMapKey)
	}

	const uniqueItemsMapKey = "uniqueItems"
	if raw, ok := proxy[uniqueItemsMapKey]; ok {
		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field uniqueItems`)
		}
		mutator.UniqueItems(decoded)
		delete(proxy, uniqueItemsMapKey)
	}

	const maxPropertiesMapKey = "maxProperties"
	if raw, ok := proxy[maxPropertiesMapKey]; ok {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field maxProperties`)
		}
		mutator.MaxProperties(decoded)
		delete(proxy, maxPropertiesMapKey)
	}

	const minPropertiesMapKey = "minProperties"
	if raw, ok := proxy[minPropertiesMapKey]; ok {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field minProperties`)
		}
		mutator.MinProperties(decoded)
		delete(proxy, minPropertiesMapKey)
	}

	const requiredMapKey = "required"
	if raw, ok := proxy[requiredMapKey]; ok {
		var decoded StringList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Required`)
		}
		for _, elem := range decoded {
			mutator.Required(elem)
		}
		delete(proxy, requiredMapKey)
	}

	const enumMapKey = "enum"
	if raw, ok := proxy[enumMapKey]; ok {
		var decoded InterfaceList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Enum`)
		}
		for _, elem := range decoded {
			mutator.Enum(elem)
		}
		delete(proxy, enumMapKey)
	}

	const allOfMapKey = "allOf"
	if raw, ok := proxy[allOfMapKey]; ok {
		var decoded SchemaList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field AllOf`)
		}
		for _, elem := range decoded {
			mutator.AllOf(elem)
		}
		delete(proxy, allOfMapKey)
	}

	const itemsMapKey = "items"
	if raw, ok := proxy[itemsMapKey]; ok {
		var decoded schema
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Items`)
		}

		mutator.Items(&decoded)
		delete(proxy, itemsMapKey)
	}

	const propertiesMapKey = "properties"
	if raw, ok := proxy[propertiesMapKey]; ok {
		var decoded SchemaMap
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Properties`)
		}
		for key, elem := range decoded {
			mutator.Property(key, elem)
		}
		delete(proxy, propertiesMapKey)
	}

	const additionaPropertiesMapKey = "additionalProperties"
	if raw, ok := proxy[additionaPropertiesMapKey]; ok {
		var decoded SchemaMap
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field AdditionaProperties`)
		}
		for key, elem := range decoded {
			mutator.AdditionaProperty(key, elem)
		}
		delete(proxy, additionaPropertiesMapKey)
	}

	const defaultValueMapKey = "default"
	if raw, ok := proxy[defaultValueMapKey]; ok {
		var decoded interface{}
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field default`)
		}
		mutator.Default(decoded)
		delete(proxy, defaultValueMapKey)
	}

	const discriminatorMapKey = "discriminator"
	if raw, ok := proxy[discriminatorMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field discriminator`)
		}
		mutator.Discriminator(decoded)
		delete(proxy, discriminatorMapKey)
	}

	const readOnlyMapKey = "readOnly"
	if raw, ok := proxy[readOnlyMapKey]; ok {
		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field readOnly`)
		}
		mutator.ReadOnly(decoded)
		delete(proxy, readOnlyMapKey)
	}

	const externalDocsMapKey = "externalDocs"
	if raw, ok := proxy[externalDocsMapKey]; ok {
		var decoded externalDocumentation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		mutator.ExternalDocs(&decoded)
		delete(proxy, externalDocsMapKey)
	}

	const exampleMapKey = "example"
	if raw, ok := proxy[exampleMapKey]; ok {
		var decoded interface{}
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field example`)
		}
		mutator.Example(decoded)
		delete(proxy, exampleMapKey)
	}

	const deprecatedMapKey = "deprecated"
	if raw, ok := proxy[deprecatedMapKey]; ok {
		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field deprecated`)
		}
		mutator.Deprecated(decoded)
		delete(proxy, deprecatedMapKey)
	}

	const xmlMapKey = "xml"
	if raw, ok := proxy[xmlMapKey]; ok {
		var decoded xml
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field XML`)
		}

		mutator.XML(&decoded)
		delete(proxy, xmlMapKey)
	}

	for name, raw := range proxy {
		if strings.HasPrefix(name, `x-`) {
			var ext interface{}
			if err := json.Unmarshal(raw, &ext); err != nil {
				return errors.Wrapf(err, `failed to unmarshal field %s`, name)
			}
			mutator.Extension(name, ext)
		}
	}

	if err := mutator.Do(); err != nil {
		return errors.Wrap(err, `failed to  unmarshal JSON`)
	}
	return nil
}

func (v *schema) QueryJSON(path string) (ret interface{}, ok bool) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	if path == "" {
		return v, true
	}

	var frag string
	if i := strings.Index(path, "/"); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ""
	}

	var target interface{}

	switch frag {
	case "type":
		target = v.typ
	case "format":
		target = v.format
	case "title":
		target = v.title
	case "multipleOf":
		target = v.multipleOf
	case "maximum":
		target = v.maximum
	case "exclusiveMaximum":
		target = v.exclusiveMaximum
	case "minimum":
		target = v.minimum
	case "exclusiveMinimum":
		target = v.exclusiveMinimum
	case "maxLength":
		target = v.maxLength
	case "minLength":
		target = v.minLength
	case "pattern":
		target = v.pattern
	case "maxItems":
		target = v.maxItems
	case "minItems":
		target = v.minItems
	case "uniqueItems":
		target = v.uniqueItems
	case "maxProperties":
		target = v.maxProperties
	case "minProperties":
		target = v.minProperties
	case "required":
		target = v.required
	case "enum":
		target = v.enum
	case "allOf":
		target = v.allOf
	case "items":
		target = v.items
	case "properties":
		target = v.properties
	case "additionalProperties":
		target = v.additionaProperties
	case "default":
		target = v.defaultValue
	case "discriminator":
		target = v.discriminator
	case "readOnly":
		target = v.readOnly
	case "externalDocs":
		target = v.externalDocs
	case "example":
		target = v.example
	case "deprecated":
		target = v.deprecated
	case "xml":
		target = v.xml
	default:
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}
	if path == "" {
		return target, true
	}
	return nil, false
}

// SchemaFromJSON constructs a Schema from JSON buffer. `dst` must
// be a pointer to `Schema`
func SchemaFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Schema)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Schema, but got %T`, dst)
	}
	var tmp schema
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}
	*v = &tmp
	return nil
}
