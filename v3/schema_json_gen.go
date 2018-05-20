package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T21:56:43+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type schemaMarshalProxy struct {
	Reference        string                `json:"$ref,omitempty"`
	Title            string                `json:"title,omitempty"`
	MultipleOf       float64               `json:"multipleOf,omitempty"`
	Maximum          float64               `json:"maximum,omitempty"`
	ExclusiveMaximum float64               `json:"exclusiveMaximum,omitempty"`
	Minimum          float64               `json:"minimum,omitempty"`
	ExclusiveMinimum float64               `json:"exclusiveMinimum,omitempty"`
	MaxLength        int                   `json:"maxLength,omitempty"`
	MinLength        int                   `json:"minLength,omitempty"`
	Pattern          string                `json:"pattern,omitempty"`
	MaxItems         int                   `json:"maxItems,omitempty"`
	MinItems         int                   `json:"minItems,omitempty"`
	UniqueItems      bool                  `json:"uniqueItems,omitempty"`
	MaxProperties    int                   `json:"maxProperties,omitempty"`
	MinProperties    int                   `json:"minProperties,omitempty"`
	Required         []string              `json:"required,omitempty"`
	Enum             []interface{}         `json:"enum,omitempty"`
	Type             PrimitiveType         `json:"type,omitempty"`
	AllOf            []Schema              `json:"allOf,omitempty"`
	OneOf            []Schema              `json:"oneOf,omitempty"`
	AnyOf            []Schema              `json:"anyOf,omitempty"`
	Not              Schema                `json:"not,omitempty"`
	Items            Schema                `json:"items,omitempty"`
	Properties       map[string]Schema     `json:"properties,omitempty"`
	Format           string                `json:"format,omitempty"`
	Default          interface{}           `json:"default,omitempty"`
	Nullable         bool                  `json:"nullable,omitempty"`
	Discriminator    Discriminator         `json:"discriminator,omitempty"`
	ReadOnly         bool                  `json:"read_only,omitempty"`
	WriteOnly        bool                  `json:"write_only,omitempty"`
	ExternalDocs     ExternalDocumentation `json:"externalDocs,omitempty"`
	Example          interface{}           `json:"example,omitempty"`
	Deprecated       bool                  `json:"deprecated,omitempty"`
}

type schemaUnmarshalProxy struct {
	Reference        string                     `json:"$ref,omitempty"`
	Title            string                     `json:"title,omitempty"`
	MultipleOf       float64                    `json:"multipleOf,omitempty"`
	Maximum          float64                    `json:"maximum,omitempty"`
	ExclusiveMaximum float64                    `json:"exclusiveMaximum,omitempty"`
	Minimum          float64                    `json:"minimum,omitempty"`
	ExclusiveMinimum float64                    `json:"exclusiveMinimum,omitempty"`
	MaxLength        int                        `json:"maxLength,omitempty"`
	MinLength        int                        `json:"minLength,omitempty"`
	Pattern          string                     `json:"pattern,omitempty"`
	MaxItems         int                        `json:"maxItems,omitempty"`
	MinItems         int                        `json:"minItems,omitempty"`
	UniqueItems      bool                       `json:"uniqueItems,omitempty"`
	MaxProperties    int                        `json:"maxProperties,omitempty"`
	MinProperties    int                        `json:"minProperties,omitempty"`
	Type             PrimitiveType              `json:"type,omitempty"`
	AllOf            []json.RawMessage          `json:"allOf,omitempty"`
	OneOf            []json.RawMessage          `json:"oneOf,omitempty"`
	AnyOf            []json.RawMessage          `json:"anyOf,omitempty"`
	Not              json.RawMessage            `json:"not,omitempty"`
	Items            json.RawMessage            `json:"items,omitempty"`
	Properties       map[string]json.RawMessage `json:"properties,omitempty"`
	Format           string                     `json:"format,omitempty"`
	Default          interface{}                `json:"default,omitempty"`
	Nullable         bool                       `json:"nullable,omitempty"`
	Discriminator    json.RawMessage            `json:"discriminator,omitempty"`
	ReadOnly         bool                       `json:"read_only,omitempty"`
	WriteOnly        bool                       `json:"write_only,omitempty"`
	ExternalDocs     json.RawMessage            `json:"externalDocs,omitempty"`
	Example          interface{}                `json:"example,omitempty"`
	Deprecated       bool                       `json:"deprecated,omitempty"`
}

func (v *schema) MarshalJSON() ([]byte, error) {
	var proxy schemaMarshalProxy
	proxy.Reference = v.reference
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
	proxy.Type = v.typ
	proxy.AllOf = v.allOf
	proxy.OneOf = v.oneOf
	proxy.AnyOf = v.anyOf
	proxy.Not = v.not
	proxy.Items = v.items
	proxy.Properties = v.properties
	proxy.Format = v.format
	proxy.Default = v.defaultValue
	proxy.Nullable = v.nullable
	proxy.Discriminator = v.discriminator
	proxy.ReadOnly = v.readOnly
	proxy.WriteOnly = v.writeOnly
	proxy.ExternalDocs = v.externalDocs
	proxy.Example = v.example
	proxy.Deprecated = v.deprecated
	return json.Marshal(proxy)
}

func (v *schema) UnmarshalJSON(data []byte) error {
	var proxy schemaUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.reference = proxy.Reference
	v.title = proxy.Title
	v.multipleOf = proxy.MultipleOf
	v.maximum = proxy.Maximum
	v.exclusiveMaximum = proxy.ExclusiveMaximum
	v.minimum = proxy.Minimum
	v.exclusiveMinimum = proxy.ExclusiveMinimum
	v.maxLength = proxy.MaxLength
	v.minLength = proxy.MinLength
	v.pattern = proxy.Pattern
	v.maxItems = proxy.MaxItems
	v.minItems = proxy.MinItems
	v.uniqueItems = proxy.UniqueItems
	v.maxProperties = proxy.MaxProperties
	v.minProperties = proxy.MinProperties
	v.typ = proxy.Type

	if len(proxy.AllOf) > 0 {
		var list []Schema
		for i, pv := range proxy.AllOf {
			var decoded schema
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %d of field AllOf`, i)
			}
			list = append(list, &decoded)
		}
		v.allOf = list
	}

	if len(proxy.OneOf) > 0 {
		var list []Schema
		for i, pv := range proxy.OneOf {
			var decoded schema
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %d of field OneOf`, i)
			}
			list = append(list, &decoded)
		}
		v.oneOf = list
	}

	if len(proxy.AnyOf) > 0 {
		var list []Schema
		for i, pv := range proxy.AnyOf {
			var decoded schema
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %d of field AnyOf`, i)
			}
			list = append(list, &decoded)
		}
		v.anyOf = list
	}

	if len(proxy.Not) > 0 {
		var decoded schema
		if err := json.Unmarshal(proxy.Not, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Not`)
		}

		v.not = &decoded
	}

	if len(proxy.Items) > 0 {
		var decoded schema
		if err := json.Unmarshal(proxy.Items, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Items`)
		}

		v.items = &decoded
	}

	if len(proxy.Properties) > 0 {
		m := make(map[string]Schema)
		for key, pv := range proxy.Properties {
			var decoded schema
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Properties`, key)
			}
			m[key] = &decoded
		}
		v.properties = m
	}
	v.format = proxy.Format
	v.defaultValue = proxy.Default
	v.nullable = proxy.Nullable

	if len(proxy.Discriminator) > 0 {
		var decoded discriminator
		if err := json.Unmarshal(proxy.Discriminator, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Discriminator`)
		}

		v.discriminator = &decoded
	}
	v.readOnly = proxy.ReadOnly
	v.writeOnly = proxy.WriteOnly

	if len(proxy.ExternalDocs) > 0 {
		var decoded externalDocumentation
		if err := json.Unmarshal(proxy.ExternalDocs, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		v.externalDocs = &decoded
	}
	v.example = proxy.Example
	v.deprecated = proxy.Deprecated
	return nil
}
